package services

import (
	"MT-GO/database"
	"MT-GO/structs"
	"fmt"
)

type Trader struct {
	*structs.Trader
}

func (t *Trader) getAssortItemByID(id string) []*structs.AssortItem {
	item, ok := t.Index.Assort.Items[id]
	if ok {
		return []*structs.AssortItem{t.Assort.Items[item]}
	}

	parentItems, parentOK := t.Index.Assort.ParentItems[id]
	if !parentOK {
		fmt.Println("Assort Item", id, "does not exist for", t.Base["nickname"])
		return nil
	}

	items := make([]*structs.AssortItem, 0, len(parentItems))
	for _, index := range parentItems {
		items = append(items, t.Assort.Items[index])
	}

	return items
}

// GetTraderLoyaltyLevel determines the loyalty level of a trader based on character attributes
func GetTraderLoyaltyLevel(traderID string, character *structs.PlayerTemplate) int {
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
