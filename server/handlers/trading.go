package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	suites := database.GetProfileByUID(sessionID).Storage.Suites

	storage := map[string]interface{}{
		"_id":    sessionID,
		"suites": suites,
	}

	body := services.ApplyResponseBody(storage)
	services.ZlibJSONReply(w, body)
}

const traderSettingsRoute string = "/client/trading/api/traderSettings"

func TradingTraderSettings(w http.ResponseWriter, r *http.Request) {
	traders := database.GetTraders()
	data := make([]map[string]interface{}, 0, len(traders))

	for _, trader := range traders {
		data = append(data, trader.Base)
	}

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, body)
}

const prod string = "https://prod.escapefromtarkov.com"
const imagesPath string = "assets/images/"

var mime = map[string]string{".jpg": "image/jpeg", ".png": "image/png"}
var extensions = []string{".jpg", ".png"}

func TradingFiles(w http.ResponseWriter, r *http.Request) {
	icon := strings.Split(r.RequestURI, "/")
	imagePath := filepath.Join(imagesPath, icon[2], icon[3], strings.TrimSuffix(icon[4], ".jpg"))

	for _, ext := range extensions {
		path := imagePath + ext
		if !tools.FileExist(path) {
			continue
		}

		fmt.Println("Image exists in ", path, " serving...")
		services.ServeFile(w, path, mime[ext])
		return
	}

	if tools.CheckInternet() {
		client := &http.Client{}
		prodURL := prod + strings.TrimSuffix(r.RequestURI, ".jpg")

		for _, ext := range extensions {
			path := prodURL + ext

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			req.Header.Set("User-Agent", "ballsack")
			response, err := client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			if response.StatusCode != http.StatusOK {
				response.Body.Close()
				continue
			}
			defer response.Body.Close()

			imagePath = imagePath + ext
			dir := filepath.Dir(imagePath)
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			file, err := os.Create(imagePath)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			defer file.Close()

			_, err = io.Copy(file, response.Body)
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			fmt.Println("Successfully downloaded to ", imagePath)
			services.ServeFile(w, imagePath, mime[ext])
			return
		}
	}
}

const (
	customizationPrefix string = "/client/trading/customization/"
	customizationSuffix string = "/offers"
)

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	traderId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, customizationPrefix), customizationSuffix)

	suits := database.GetTraders()[traderId].Suits
	body := services.ApplyResponseBody(suits)
	services.ZlibJSONReply(w, body)
}

const assort string = "/client/trading/api/getTraderAssort/"

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	traderId := strings.TrimPrefix(r.URL.Path, assort)

	assort := database.GetTraders()[traderId].Assort
	body := services.ApplyResponseBody(assort)
	services.ZlibJSONReply(w, body)
	fmt.Println("You need to add proper transactions")
}
