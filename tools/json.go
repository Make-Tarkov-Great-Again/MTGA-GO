package tools

import (
	"log"
	"strings"

	"github.com/goccy/go-json"
)

type Data struct {
	Data json.RawMessage `json:"data"`
}

func GetJSONRawMessage(path string) json.RawMessage {
	b, err := ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	rawJson := json.RawMessage(b)
	var data Data
	if strings.Contains(string(rawJson), "\"data\"") {
		if err := json.Unmarshal(rawJson, &data); err != nil {
			log.Fatalln(err)
		}
		return data.Data
	}
	return rawJson
}
