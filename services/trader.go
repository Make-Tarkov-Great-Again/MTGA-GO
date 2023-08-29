package services

import (
	"MT-GO/database"
	"MT-GO/structs"
)

/*
	type TradingService struct {
		Loyalty map[string]int
		Assort map[string]
	}
*/

// GetTraderLoyaltyLevel determines the loyalty level of a trader based on character attributes
func GetTraderLoyaltyLevel(traderID string, character *structs.PlayerTemplate) int {
	loyaltyLevels := database.GetTraderByID(traderID).Base["loyaltyLevels"].([]interface{})

	_, ok := character.TradersInfo["638f541a29ffd1183d187f57"]
	if !ok {
		return -1
	}

	length := len(loyaltyLevels)
	for level := 0; level < length; level++ {
		if loyaltyLevels[level] == nil {
			continue
		}

		loyalty := loyaltyLevels[level].(map[string]interface{})
		if character.Info.Level < int(loyalty["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(loyalty["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(loyalty["minStanding"].(float64)) {
			return level
		}
	}

	return length
}
