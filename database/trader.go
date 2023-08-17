package database

import (
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strconv"
)

var traders = map[string]map[string]interface{}{}

func GetTraders() map[string]map[string]interface{} {
	return traders
}

func setTraders() {

	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		panic(err)
	}

	for _, dir := range directory {
		trader := map[string]interface{}{}

		currentTraderPath := filepath.Join(traderPath, dir)

		basePath := filepath.Join(currentTraderPath, "base.json")
		if tools.FileExist(basePath) {
			trader["Base"] = processBase(basePath)
		}

		assortPath := filepath.Join(currentTraderPath, "assort.json")
		if tools.FileExist(assortPath) {
			trader["Assort"] = processAssort(assortPath)
			trader["BaseAssort"] = trader["Assort"]
		}

		questsPath := filepath.Join(currentTraderPath, "questassort.json")
		if tools.FileExist(questsPath) {
			trader["QuestAssort"] = processQuestAssort(questsPath)
		}

		suitsPath := filepath.Join(currentTraderPath, "suits.json")
		if tools.FileExist(suitsPath) {
			trader["Suits"] = processSuits(suitsPath)
		}

		dialoguesPath := filepath.Join(currentTraderPath, "dialogue.json")
		if tools.FileExist(dialoguesPath) {
			trader["Dialogue"] = processDialogues(dialoguesPath)
		}

		//traders[dir] = map[string]interface{}{}
		traders[dir] = trader
	}
}

func processBase(basePath string) map[string]interface{} {
	base := map[string]interface{}{}

	var dynamic map[string]interface{} //here we fucking go

	raw := tools.GetJSONRawMessage(basePath)
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

type Assort struct {
	BarterScheme    map[string][][]*Scheme
	Items           []map[string]interface{}
	LoyalLevelItems map[string]int
}

type Scheme struct {
	Tpl   string `json:"_tpl"`
	Count int    `json:"count"`
}

func processAssort(assortPath string) *Assort {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(assortPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	assort := &Assort{}

	items, ok := dynamic["items"].([]interface{})
	if ok {
		assort.Items = make([]map[string]interface{}, 0, len(items))

		for _, item := range items {
			i := item.(map[string]interface{})
			upd, ok := i["upd"].(map[string]interface{})
			if !ok {
				assort.Items = append(assort.Items, i)
				continue
			}

			fireMode, ok := upd["FireMode"].(map[string]interface{})
			if ok {
				fireModeValue, ok := fireMode["FireMode"].(string)

				if ok {
					upd["FireMode"] = map[string]string{
						"FireMode": fireModeValue,
					}
				}
			}

			foldable, ok := upd["Foldable"].(map[string]interface{})
			if ok {
				folded, ok := foldable["Folded"].(bool)
				if ok {
					upd["Foldable"] = map[string]bool{
						"Folded": folded,
					}
				}
			}

			stackCount, ok := upd["StackObjectsCount"].(interface{})
			if ok {
				stackCountValue, ok := stackCount.(float64)
				if ok {
					upd["StackObjectsCount"] = int(stackCountValue)
				}
			}

			buyCurrent, ok := upd["BuyRestrictionCurrent"].(float64)
			if ok {
				upd["BuyRestrictionCurrent"] = int(buyCurrent)
			}

			buyRestrictionMax, ok := upd["BuyRestrictionMax"].(float64)
			if ok {
				upd["BuyRestrictionMax"] = int(buyRestrictionMax)
			}
			assort.Items = append(assort.Items, i)
		}
	} else {
		panic("Items not found")
	}

	barterSchemes, ok := dynamic["barter_scheme"].(map[string]interface{})
	if ok {
		assort.BarterScheme = make(map[string][][]*Scheme)

		for id, scheme := range barterSchemes {
			scheme, ok := scheme.([]interface{})
			if ok {
				trueScheme := make([][]*Scheme, 0, len(scheme))
				for _, value := range scheme {
					value := value.([]interface{})
					template := make([]*Scheme, 0, len(value))

					for _, v := range value {
						if v, ok := v.(map[string]interface{}); ok {
							input := &Scheme{
								Tpl:   v["_tpl"].(string),
								Count: int(v["count"].(float64)),
							}
							template = append(template, input)
						}
					}
					trueScheme = append(trueScheme, template)
				}
				assort.BarterScheme[id] = trueScheme
			}
		}
	} else {
		panic("Barter scheme not found")
	}

	loyalLevelItems, ok := dynamic["loyal_level_items"].(map[string]interface{})
	if ok {
		assort.LoyalLevelItems = map[string]int{}
		for id, item := range loyalLevelItems {
			assort.LoyalLevelItems[id] = int(item.(float64))
		}
	}

	return assort
}

func processQuestAssort(questsPath string) map[string][]string {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(questsPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}
	raw = nil

	quests := map[string][]string{}
	for k, v := range dynamic {
		v := v.(map[string]interface{})

		length := len(v)
		quests[k] = make([]string, 0, len(v))
		if length == 0 {
			continue
		}

		for _, quest := range v {
			quests[k] = append(quests[k], quest.(string))
		}
	}

	return quests
}

func processDialogues(dialoguesPath string) map[string][]string {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(dialoguesPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}
	raw = nil

	dialogues := map[string][]string{}
	for k, v := range dynamic {
		v := v.([]interface{})

		length := len(v)
		dialogues[k] = make([]string, 0, len(v))
		if length == 0 {
			continue
		}

		for _, dialogue := range v {
			dialogues[k] = append(dialogues[k], dialogue.(string))
		}
	}

	return dialogues
}

func processSuits(dialoguesPath string) []map[string]interface{} {
	var dynamic []map[string]interface{}
	raw := tools.GetJSONRawMessage(dialoguesPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}
	raw = nil

	return dynamic
}
