package MTGACoop

import (
	common "MT-GO/user/mods/MTGA-Coop/mod/Common"
	coopmatch "MT-GO/user/mods/MTGA-Coop/mod/Functionality/CoopMatch"
)

func Mod() {
	Load()
}

func Load() {

	//TODO: [RE-USE]
	// Setting the same instance in both CoopMatch and WebSocket
	// doesn't make sense just make them share the same one
	// or if they don't both need it, just move this commonGround
	// thing into the package that needs it

	//TODO: [DATA STRUCTURE]
	// We need to comb through the existing structs and assign
	// the correct specific types (especially int/float)
	// Generics (interface{}) are disgusting and we should know
	// exactly what data is being sent

	//TODO: [LOOPING]
	// If we can avoid a loop, do so. Always try to avoid a loop.

	commonGround := common.NewWebSocketHandler() //Create a common ground between

	coopmatch.Init(commonGround) //Coopmatch
	//websocket.Init(commonGround)                 //And websocket, both of these should be synced. this way.
}

/* ---------------------- Boring mod bindings below lol --------------------- */
