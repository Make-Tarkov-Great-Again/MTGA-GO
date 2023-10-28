package LiveFleaPrices

import (
	"MT-GO/database"
	"MT-GO/tools"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	liveFleaPriceDB = "https://raw.githubusercontent.com/Make-Tarkov-Great-Again/TarkovDev/main/prices.json"
	refresh         = time.Duration(3600) * time.Second
)

func Mod() {
	if tools.CheckInternet() {
		go func() {
			timer := time.NewTimer(0) // 0 timer so it immediately updates prices

			for {
				<-timer.C
				response, err := http.Get(liveFleaPriceDB)
				if err != nil {
					log.Println("Error fetching data:", err)
					return
				}
				defer response.Body.Close()

				body, err := io.ReadAll(response.Body)
				if err != nil {
					log.Println("Error reading response body:", err)
					return
				}

				data := make(map[string]int32)
				if err := json.Unmarshal(body, &data); err != nil {
					log.Println("Error unmarshalling JSON:", err)
					return
				}

				prices := database.GetPrices() //get and update prices in cache
				for k, v := range data {
					prices[k] = &v
				}

				timer.Reset(refresh) // set next timer to start in an hour (in seconds)
			}
		}()
	}
	log.Println("LiveFleaPrices won't work without internet connection")
	return
}
