package pkg

import (
	"MT-GO/tools"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ResponseBody struct {
	Err    int `json:"err"`
	Errmsg any `json:"errmsg"`
	Data   any `json:"data"`
}

type CRCResponseBody struct {
	Err    any     `json:"err"`
	Errmsg any     `json:"errmsg"`
	Data   any     `json:"data"`
	Crc    *uint32 `json:"crc"`
}

const (
	cookieHeader string = "Cookie"
	prod         string = "https://prod.escapefromtarkov.com"
	imagesPath   string = "assets/images/"
)

// GetSessionID returns current sessionID from the header, if available
func GetSessionID(r *http.Request) (string, error) {
	output := r.Header.Get(cookieHeader)[10:]
	if output == "" {
		return "", errors.New("cookie header is empty")
	}
	return output, nil
}

// ApplyCRCResponseBody appends data to CRCResponseBody and returns it
func ApplyCRCResponseBody(data any, crc *uint32) *CRCResponseBody {
	return &CRCResponseBody{
		Err:    0,
		Errmsg: nil,
		Data:   data,
		Crc:    crc,
	}
}

// ApplyResponseBody appends data to ResponseBody and returns it
func ApplyResponseBody(data any) *ResponseBody {
	return &ResponseBody{
		Err:    0,
		Errmsg: nil,
		Data:   data,
	}
}

var mime = map[string]string{
	".png": "image/png",
	".jpg": "image/jpeg",
}

var downloadLocal bool

func SetDownloadLocal(result bool) {
	downloadLocal = result
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	if err := os.MkdirAll(imagesPath, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	main := chi.URLParam(r, "main")
	typeOf := chi.URLParam(r, "type")
	fileName := chi.URLParam(r, "file")

	dir := filepath.Join(imagesPath, main, typeOf)

	imagePath := filepath.Join(dir, fileName)
	if files, _ := tools.GetFilesFrom(dir); files != nil {
		for ext, mimeType := range mime {
			path := filepath.Join(imagePath, ext)
			if _, ok := files[path]; !ok {
				continue
			}

			log.Println("Image exists in", path, ", serving...")
			ServeFileLocal(w, path, mimeType)
			return
		}
	}

	if !tools.CheckInternet() {
		log.Println("Image does not exist in local directory, and we cannot fetch from servers so... here's a blank image!!!!")
		return
	}

	client := &http.Client{}
	prodURL := prod + r.RequestURI[:len(r.RequestURI)-4]

	for ext, mimeType := range mime {
		path := prodURL + ext

		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			log.Println(err)
			continue
		}

		req.Header.Set("User-Agent", "ballsack")
		response, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			if err := response.Body.Close(); err != nil {
				log.Panicln(err)
			}
			continue
		}

		if !downloadLocal {
			ServeFileURL(w, response.Body, mimeType)
			return
		}

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Println(err)
			continue
		}

		imagePath += ext
		file, err := os.Create(imagePath)
		if err != nil {
			log.Println(err)
			continue
		}
		defer file.Close()

		if _, err = io.Copy(file, response.Body); err != nil {
			log.Println(err)
			continue
		}

		log.Println("Successfully downloaded to", imagePath)
		ServeFileLocal(w, imagePath, mimeType)
		return
	}

}

func ServeFileURL(w http.ResponseWriter, body io.ReadCloser, mime string) {
	w.Header().Set("Content-Type", mime)
	if _, err := io.Copy(w, body); err != nil {
		log.Fatalln(err)
	}
}

func ServeFileLocal(w http.ResponseWriter, imagePath, mime string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	w.Header().Set("Content-Type", mime)
	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatalln(err)
	}
}

type ContextKey struct{}

func GetParsedBody(r *http.Request) any {
	return r.Context().Value(&ContextKey{})
}
