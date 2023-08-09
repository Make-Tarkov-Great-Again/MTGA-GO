package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
)

var handbook = structs.Handbook{}

func GetHandbook() *structs.Handbook {
	return &handbook
}

func setHandbook() {
	raw := tools.GetJSONRawMessage(handbookPath)
	err := json.Unmarshal(raw, &handbook)
	if err != nil {
		panic(err)
	}
}
