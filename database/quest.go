package database

import (
	"MT-GO/tools"
	"encoding/json"
	"strconv"
	"strings"
)

var quests map[string]interface{}

func GetQuests() map[string]interface{} {
	return quests
}

func setQuests() {

	raw := tools.GetJSONRawMessage(questsPath)

	dynamic := make(map[string]interface{})
	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	for _, v := range dynamic {
		quest := v.(map[string]interface{})

		conditions, ok := quest["conditions"].(map[string]interface{})
		if !ok {
			panic("quests.conditions is not a map")
		}

		for _, v := range conditions {
			conditionType := v.([]interface{})

			for _, v := range conditionType {
				condition := v.(map[string]interface{})

				props, ok := condition["_props"].(map[string]interface{})
				if !ok {
					continue
				}

				value, ok := props["value"].(string)
				if ok {
					props["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}

				counter, ok := props["counter"].(map[string]interface{})
				if !ok {
					continue
				}

				conditions := counter["conditions"].([]interface{})

				for _, v := range conditions {
					condition := v.(map[string]interface{})

					props, ok := condition["_props"].(map[string]interface{})
					if !ok {
						continue
					}

					value, ok := props["value"].(string)
					if !ok {
						continue
					}

					props["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}

			}
		}

		rewards, ok := quest["rewards"].(map[string]interface{})
		if !ok {
			panic("quests.rewards is not a map")
		}

		for _, v := range rewards {
			rewardType := v.([]interface{})

			for _, v := range rewardType {
				reward := v.(map[string]interface{})

				value, ok := reward["value"].(string)
				if ok {
					value = strings.TrimSpace(value)
					reward["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}
			}
		}

	}

	jsonData, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &quests)
	if err != nil {
		panic(err)
	}
}
