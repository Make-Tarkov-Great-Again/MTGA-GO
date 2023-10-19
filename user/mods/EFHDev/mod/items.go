// Items go brrrrr
package mod

import (
	"MT-GO/database"
	"fmt"
)

func Modify(passed *database.ModInfo) {
	ModifyAmmo()
}

func ModifyAmmo() {
	var count int = 0
	items := database.GetItems()
	for _, item := range items {
		if item.Type != "Item" || item.Parent != "5485a8684bdc2da71d8b4567" {
			continue
		}

		caliber, ok := item.Props["Caliber"].(string)
		if !ok {
			continue
		}

		switch caliber {
		case "Caliber9x19PARA":
		case "Caliber1143x23ACP":
		case "Caliber762x25TT":
		case "Caliber9x18PM":
		case "Caliber9x18PMM":
		case "Caliber9x33R":
		case "Caliber57x28":
		case "Caliber46x30":
		case "Caliber9x21":
		case "Caliber556x45NATO":
		case "Caliber762x54R":
		case "Caliber762x39":
		case "Caliber127x55":
		case "Caliber9x39":
		case "Caliber12g":
		case "Caliber40x46":
		case "Caliber40mmRU":
		case "Caliber366TKM":
		case "Caliber545x39":

			fmt.Println(item.Props["BackgroundColor"])

			//TODO: This isn't sticking for everything which is... weird

			item.Props["BackgroundColor"] = "red"
			item.Props["StackMaxSize"] = 160

			check := database.GetItemByUID(item.ID)
			fmt.Println(check.Props["BackgroundColor"])
			count = count + 1

		}

	}
}
