// Pog :)
package EscapeFromHell

import (
	items "MT-GO/user/mods/EFHDev/mod"
	"fmt"
)

func Mod() {
	fmt.Println("Loading Escape from Hell....")
	items.Modify()
	defer fmt.Println("Loaded mod Escape From hell lol.")

}
