package services

import (
	"MT-GO/database"
	"MT-GO/structs"
)

func GetTraderLoyaltyLevel(traderID string, character *structs.PlayerTemplate) int {
	loyaltyLevels := database.GetTraderByID(traderID).Base["loyaltyLevels"].([]interface{})

	if traderID == "638f541a29ffd1183d187f57" {
		_, ok := character.TradersInfo["638f541a29ffd1183d187f57"]
		if !ok {
			return -1
		}

		zero := loyaltyLevels[0].(map[string]interface{})
		if character.Info.Level < int(zero["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(zero["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(zero["minStanding"].(float64)) {
			return 1
		}
	}

	length := len(loyaltyLevels)

	if length > 0 && loyaltyLevels[1] != nil {
		one := loyaltyLevels[1].(map[string]interface{})
		if character.Info.Level < int(one["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(one["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(one["minStanding"].(float64)) {
			return 1
		}
	}

	// Check if index 1 exists and is not nil
	if length > 1 && loyaltyLevels[2] != nil {
		two := loyaltyLevels[2].(map[string]interface{})
		if character.Info.Level < int(two["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(two["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(two["minStanding"].(float64)) {
			return 2
		}
	}

	// Check if index 2 exists and is not nil
	if length > 2 && loyaltyLevels[3] != nil {
		three := loyaltyLevels[3].(map[string]interface{})
		if character.Info.Level < int(three["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(three["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(three["minStanding"].(float64)) {
			return 3
		}
	}

	return 4
}
