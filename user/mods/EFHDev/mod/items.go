// Items go brrrrr
package items

import (
	"MT-GO/database"
	"fmt"
)

func Modify() {
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
