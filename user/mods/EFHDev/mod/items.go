// Items go brrrrr
package mod

import (
	"MT-GO/database"
	"fmt"
)

var config any

func Modify(passed any) {
	AP63 := database.GetItemByUID("5c925fa22e221601da359b7b")
	AP63.Props["BackgroundColor"] = "red"
	err := database.SetItemByUID("5c925fa22e221601da359b7b", AP63)
	if err != nil {
		fmt.Printf("Error setting item %s: %s", AP63.ID, err)
	} else {
		fmt.Println("Set AP 6.3 to have red background lol")
		test := database.GetItemByUID("5c925fa22e221601da359b7b")
		fmt.Printf("Just to be sure... its background is now %s in the database... LETS GOOOO \n", test.Props["BackgroundColor"])
	}

}

func AmmoStacks() {
	var count int = 0
	items := database.GetItems()
	for _, item := range items {
		if item.Type != "Item" || item.Type == "Node" {
			break
		} else {
			switch item.Props["Caliber"] {
			case "Caliber9x19PARA":
			case "Caliber1143x23ACP":
			case "Caliber762x25TT":
			case "Caliber9x18PM":
			case "Caliber9x18PMM":
			case "Caliber9x33R":
			case "Caliber57x28":
			case "Caliber46x30":
			case "Caliber9x21":
				item.Props["StackMaxSize"] = 160
				database.SetItemByUID(item.ID, item)
				count = count + 1

			}
		}
	}
}

func ConfigPass(c any) {

}

type Configurable interface {
	GetConfig()
}
