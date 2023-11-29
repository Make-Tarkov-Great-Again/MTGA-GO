package pkg

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"MT-GO/tools"
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
func GetSessionID(r *http.Request) string {
	return r.Header.Get(cookieHeader)[10:]
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
	".jpg": "image/jpeg",
	".png": "image/png",
}

var downloadLocal bool

func SetDownloadLocal(result bool) {
	downloadLocal = result
}

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	icon := strings.Split(r.RequestURI, "/")
	imagePath := filepath.Join(imagesPath, icon[2], icon[3], strings.TrimSuffix(icon[4], ".jpg"))

	for ext, mimeType := range mime {
		path := imagePath + ext
		if !tools.FileExist(path) {
			continue
		}

		log.Println("Image exists in ", path, " serving...")
		ServeFileLocal(w, path, mimeType)
		return
	}

	if tools.CheckInternet() {
		client := &http.Client{}
		prodURL := prod + strings.TrimSuffix(r.RequestURI, ".jpg")

		for ext, mimeType := range mime {
			path := prodURL + ext

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				log.Fatalln(err)
			}

			req.Header.Set("User-Agent", "ballsack")
			response, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}(response.Body)

			if response.StatusCode != http.StatusOK {
				err := response.Body.Close()
				if err != nil {
					log.Panicln(err)
				}
				continue
			}

			if downloadLocal {
				imagePath += ext
				dir := filepath.Dir(imagePath)
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					log.Fatalln(err)
				}

				file, err := os.Create(imagePath)
				if err != nil {
					log.Fatalln(err)
				}
				defer func(file *os.File) {
					err := file.Close()
					if err != nil {
						log.Fatalln(err)
					}
				}(file)

				_, err = io.Copy(file, response.Body)
				if err != nil {
					log.Fatalln(err)
				}

				log.Println("Successfully downloaded to", imagePath)
				ServeFileLocal(w, imagePath, mimeType)
			} else {
				ServeFileURL(w, response.Body, mimeType)
			}
			return
		}
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

type contextKey string

const ParsedBodyKey contextKey = "ParsedBody"

func GetParsedBody(r *http.Request) any {
	return r.Context().Value(ParsedBodyKey)
}
