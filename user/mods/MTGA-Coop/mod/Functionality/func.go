package functionality

import (
	"fmt"
)

func getCoopMatch(serverID string) string {

	if serverID == "" {
		fmt.Println("[getCoopMatch] Server ID is nil.")
		return ""
	}
	//        if (serverId === undefined) {
	//            console.error("getCoopMatch -- no serverId provided");
	//            return undefined;
	//        }
	//        if (CoopMatch_1.CoopMatch.CoopMatches[serverId] === undefined) {
	//            console.error(`getCoopMatch -- no server of ${serverId} exists`);
	//            return undefined;
	//        }
	//        return CoopMatch_1.CoopMatch.CoopMatches[serverId];
}
