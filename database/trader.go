package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
	"strconv"
)

func setTraders() map[string]*structs.Trader {
	traders := make(map[string]*structs.Trader)

	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		return traders
	}
	for _, dir := range directory {
		trader := structs.Trader{}

		currentTraderPath := traderPath + dir + "/"

		if tools.FileExist(currentTraderPath + "base.json") {
			trader.Base = processBase(currentTraderPath)
		}

		if tools.FileExist(currentTraderPath + "assort.json") {
			trader.Assort = processAssort(currentTraderPath)
			trader.BaseAssort = trader.Assort
		}

		if tools.FileExist(currentTraderPath + "questassort.json") {
			raw := tools.GetJSONRawMessage(currentTraderPath + "questassort.json")
			questAssort := structs.QuestAssort{}

			err = json.Unmarshal(raw, &questAssort)
			if err != nil {
				panic(err)
			}
			trader.QuestAssort = questAssort
		}

		if tools.FileExist(currentTraderPath + "suits.json") {
			suits := []structs.Suit{}

			raw := tools.GetJSONRawMessage(currentTraderPath + "suits.json")
			err = json.Unmarshal(raw, &suits)
			if err != nil {
				panic(err)
			}
			trader.Suits = suits
		}

		if tools.FileExist(currentTraderPath + "dialogue.json") {
			dialogue := structs.Dialogue{}

			raw := tools.GetJSONRawMessage(currentTraderPath + "dialogue.json")
			err = json.Unmarshal(raw, &dialogue)
			if err != nil {
				panic(err)
			}
			trader.Dialogue = dialogue
		}
		traders[dir] = &trader
	}
	return traders
}

func processBase(currentTraderPath string) structs.Base {
	base := structs.Base{}

	var dynamic map[string]interface{} //here we fucking go

	raw := tools.GetJSONRawMessage(currentTraderPath + "base.json")
	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	loyaltyLevels := dynamic["loyaltyLevels"].([]interface{})
	length := len(loyaltyLevels)

	for i := 0; i < length; i++ {
		level := loyaltyLevels[i].(map[string]interface{})

		insurancePriceCoef, ok := level["insurance_price_coef"].(string)
		if !ok {
			continue
		}

		level["insurance_price_coef"], err = strconv.Atoi(insurancePriceCoef)
		if err != nil {
			panic(err)
		}
	}

	repair := dynamic["repair"].(map[string]interface{})

	repairQuality, ok := repair["quality"].(string)
	if ok {
		repair["quality"], err = strconv.ParseFloat(repairQuality, 32)
		if err != nil {
			panic(err)
		}
	}

	sanitized, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(sanitized, &base)
	if err != nil {
		panic(err)
	}

	return base
}

func processAssort(currentTraderPath string) structs.Assort {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(currentTraderPath + "assort.json")

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	items, ok := dynamic["items"].([]interface{})
	if !ok {
		panic("not okay!!!!!!!!!!!!!!!!!!!!")
	}

	for _, item := range items {
		i := item.(map[string]interface{})
		upd, ok := i["upd"].(map[string]interface{})
		if !ok {
			continue
		}

		buyRestrictionMax, ok := upd["BuyRestrictionMax"].(string)
		if !ok {
			continue
		}
		upd["BuyRestrictionMax"], err = strconv.Atoi(buyRestrictionMax)
		if err != nil {
			panic(err)
		}

	}

	sanitized, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	assort := structs.Assort{}
	err = json.Unmarshal(sanitized, &assort)
	if err != nil {
		panic(err)
	}
	return assort
}
