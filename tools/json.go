package tools

import (
	"github.com/goccy/go-json"
	"log"
)

func GetJSONRawMessage(path string) json.RawMessage {
	b, err := ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	var rawMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &rawMap); err != nil {
		return b
	}

	if data, exists := rawMap["data"]; exists {
		return data
	}

	return b
}
