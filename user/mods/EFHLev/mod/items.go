// Items go brrrrr
package items

import (
	"MT-GO/database"
	"fmt"
	"time"
)

var startTime time.Time
var config any

func Modify(passed any) {
	AP63 := database.GetItemByUID("5c925fa22e221601da359b7b")
	AP63.Props["BackgroundColor"] = "blue"
	err := database.SetItemByUID("5c925fa22e221601da359b7b", AP63)
	if err != nil {
		fmt.Printf("Error setting item %s: %s", AP63.ID, err)
	} else {
		fmt.Println("Set AP 6.3 to have blue background lol")
		test := database.GetItemByUID("5c925fa22e221601da359b7b")
		fmt.Printf("Fuck that red guy. Blu for the win. I made AP 6.3's background %s  \n", test.Props["BackgroundColor"])
	}

}

func AmmoStacks() {
	startTime = time.Now()
	var count int = 0
	var chngcount int = 0

	items := database.GetItems()
	fmt.Println("Modifying Ammo...")
	for _, item := range items {
		if item.Type == "Item" {
			switch item.Props["Caliber"] {
			case "Caliber9x19PARA",
				"Caliber1143x23ACP",
				"Caliber762x25TT",
				"Caliber9x18PM",
				"Caliber9x18PMM",
				"Caliber9x33R",
				"Caliber57x28",
				"Caliber46x30",
				"Caliber9x21":
				item.Props["StackMaxSize"] = 160
				database.SetItemByUID(item.ID, item)
				count = count + 1
				chngcount = chngcount + 1
				fmt.Println(item.Name, item.Type)
				break
			default:
				{
					count = count + 1
					break
				}

			}
		}
	}
	endTime := time.Now()
	fmt.Printf("Finished itterating over %d items, changing %d of them, in %s", count, chngcount, endTime.Sub(startTime))

}

func ConfigPass(c any) {

}

type Configurable interface {
	GetConfig()
}
