package services

import (
	"MT-GO/database"
	"MT-GO/structs"
)

func GetTraderLoyaltyLevel(traderID string, character *structs.PlayerTemplate) int {

	loyaltyLevels := database.GetTraderByID(traderID).Base["loyaltyLevels"].([]interface{})

	//one := loyaltyLevels[0].(map[string]interface{})
	two := loyaltyLevels[1].(map[string]interface{})
	three := loyaltyLevels[2].(map[string]interface{})
	four := loyaltyLevels[3].(map[string]interface{})

	if character.Info.Level < two["minLevel"].(int) ||
		character.TradersInfo[traderID].SalesSum < float32(two["minSalesSum"].(float64)) ||
		character.TradersInfo[traderID].Standing < float32(two["minStanding"].(float64)) {
		return 1
	}

	if character.Info.Level < three["minLevel"].(int) ||
		character.TradersInfo[traderID].SalesSum < float32(three["minSalesSum"].(float64)) ||
		character.TradersInfo[traderID].Standing < float32(three["minStanding"].(float64)) {
		return 2
	}

	if character.Info.Level < four["minLevel"].(int) ||
		character.TradersInfo[traderID].SalesSum < float32(four["minSalesSum"].(float64)) ||
		character.TradersInfo[traderID].Standing < float32(four["minStanding"].(float64)) {
		return 3
	}

	return 4
}
