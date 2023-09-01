package database

import (
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var items map[string]*DatabaseItem

func GetItems() map[string]*DatabaseItem {
	return items
}

func setItems() {
	raw := tools.GetJSONRawMessage(itemsPath)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		panic(err)
	}
}

type DatabaseItem struct {
	ID         string                 `json:"_id"`
	Name       string                 `json:"_name"`
	Parent     string                 `json:"_parent"`
	Type       string                 `json:"_type"`
	Properties map[string]interface{} `json:"_props"`
	Proto      string                 `json:"_proto"`
}
