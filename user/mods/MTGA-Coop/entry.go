package MTGACoop

import (
	common "MT-GO/user/mods/MTGA-Coop/mod/Common"
	coopmatch "MT-GO/user/mods/MTGA-Coop/mod/Functionality/CoopMatch"
	websocket "MT-GO/user/mods/MTGA-Coop/mod/WebSocket"
)

func Mod() {
	Load()
}

func Load() {
	commonGround := common.NewWebSocketHandler() //Create a common ground between
	coopmatch.Init(commonGround)                 //Coopmatch
	websocket.Init(commonGround)                 //And websocket, both of these should be synced. this way.
}

/* ---------------------- Boring mod bindings below lol --------------------- */
