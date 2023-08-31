package database

import (
	"MT-GO/structs"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var items map[string]*structs.DatabaseItem

func GetItems() map[string]*structs.DatabaseItem {
	return items
}

func setItems() {
	raw := tools.GetJSONRawMessage(itemsPath)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		panic(err)
	}
}
