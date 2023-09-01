package services

import (
	"MT-GO/database"
)

// GetTraderLoyaltyLevel determines the loyalty level of a trader based on character attributes
func GetTraderLoyaltyLevel(traderID string, character *database.Character) int {
	loyaltyLevels := database.GetTraderByID(traderID).Base["loyaltyLevels"].([]interface{})

	_, ok := character.TradersInfo[traderID]
	if !ok {
		return -1
	}

	length := len(loyaltyLevels)
	for index := 0; index < length; index++ {
		loyalty := loyaltyLevels[index].(map[string]interface{})
		if character.Info.Level < int(loyalty["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(loyalty["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(loyalty["minStanding"].(float64)) {
			return index + 1
		}
	}

	return length
}
