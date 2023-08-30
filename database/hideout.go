package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
)

var hideout = structs.Hideout{}

// GetHideout retrieves the current hideout configuration.
func GetHideout() *structs.Hideout {
	return &hideout
}

const (
	areasPath           string = hideoutPath + "areas.json"
	productionPath      string = hideoutPath + "production.json"
	qtePath             string = hideoutPath + "qte.json"
	scavcasePath        string = hideoutPath + "scavcase.json"
	hideoutSettingsPath string = hideoutPath + "settings.json"
)

// setHideoutScavcase sets the hideout scavcase items and their indexes.
func setHideout() {

	if tools.FileExist(areasPath) {
		areas := tools.GetJSONRawMessage(areasPath)
		areasMap := []map[string]interface{}{}
		err := json.Unmarshal(areas, &areasMap)
		if err != nil {
			panic(err)
		}
		setHideoutAreas(areasMap)
	}

	if tools.FileExist(productionPath) {
		recipies := tools.GetJSONRawMessage(productionPath)
		productionsMap := []map[string]interface{}{}
		err := json.Unmarshal(recipies, &productionsMap)
		if err != nil {
			panic(err)
		}
		setHideoutRecipes(productionsMap)
	}

	if tools.FileExist(scavcasePath) {
		scavcase := tools.GetJSONRawMessage(scavcasePath)
		scavcaseReturns := []map[string]interface{}{}
		err := json.Unmarshal(scavcase, &scavcaseReturns)
		if err != nil {
			panic(err)
		}
		setHideoutScavcase(scavcaseReturns)
	}

	if tools.FileExist(qtePath) {
		qte := tools.GetJSONRawMessage(qtePath)
		hideout.QTE = []map[string]interface{}{}

		err := json.Unmarshal(qte, &hideout.QTE)
		if err != nil {
			panic(err)
		}
	}

	if tools.FileExist(hideoutSettingsPath) {
		settings := tools.GetJSONRawMessage(hideoutSettingsPath)
		err := json.Unmarshal(settings, &hideout.Settings)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println()
}

var HideoutAreaNames = map[string]int8{
	"NotSet":               -1,
	"Vents":                0,
	"Security":             1,
	"Lavatory":             2,
	"Stash":                3,
	"Generator":            4,
	"Heating":              5,
	"WaterCollector":       6,
	"MedStation":           7,
	"NutritionUnit":        8,
	"RestSpace":            9,
	"Workbench":            10,
	"IntelCenter":          11,
	"ShootingRange":        12,
	"Library":              13,
	"ScavCase":             14,
	"Illumination":         15,
	"PlaceOfFame":          16,
	"AirFiltering":         17,
	"SolarPower":           18,
	"BoozeGenerator":       19,
	"BitcoinFarm":          20,
	"ChristmasTree":        21,
	"EmergencyWall":        22,
	"Gym":                  23,
	"WeaponStand":          24,
	"WeaponStandSecondary": 25,
}

// GetHideoutAreaByAreaType retrieves a hideout area by its type int8.
func GetHideoutAreaByAreaType(_type int8) *map[string]interface{} {
	index, ok := hideout.Index.Areas[_type]
	if !ok {
		fmt.Println("Area Type ", _type, " does not exist")
		return nil
	}

	hideoutArea := hideout.Areas[index]
	return &hideoutArea
}

// GetHideoutAreaByName retrieves a hideout area by its name.
func GetHideoutAreaByName(name string) *map[string]interface{} {
	area, ok := HideoutAreaNames[name]
	if !ok {
		fmt.Println("Hideout Area ", name, " does not exist")
		return nil
	}

	hideoutArea := hideout.Areas[hideout.Index.Areas[area]]
	return &hideoutArea
}

// setHideoutAreas sets the hideout areas and their indexes.
func setHideoutAreas(areas []map[string]interface{}) {
	hideout.Areas = make([]map[string]interface{}, 0, len(areas))
	hideout.Index.Areas = make(map[int8]int8)

	for index, area := range areas {
		areaType := int8(area["type"].(float64))

		hideout.Index.Areas[areaType] = int8(index)
		hideout.Areas = append(hideout.Areas, area)
	}
}

// GetHideoutRecipeByID retrieves a hideout production by its ID.
func GetHideoutRecipeByID(rid string) *map[string]interface{} {
	index, ok := hideout.Index.Recipes[rid]
	if ok {
		recipe := hideout.Recipes[index]
		return &recipe
	}
	fmt.Println("Recipe ", rid, " does not exist")
	return nil
}

// setHideoutRecipes sets the hideout production recipes and their indexes.
func setHideoutRecipes(recipies []map[string]interface{}) {
	hideout.Recipes = make([]map[string]interface{}, 0, len(recipies))
	hideout.Index.Recipes = make(map[string]int16)

	for index, recipe := range recipies {
		pid := recipe["_id"].(string)

		hideout.Index.Recipes[pid] = int16(index)
		hideout.Recipes = append(hideout.Recipes, recipe)
	}
}

// GetHideoutRecipeByID retrieves a scavcase production by its ID.
func GetScavCaseRecipeByID(rid string) *map[string]interface{} {
	index, ok := hideout.Index.ScavCase[rid]
	if ok {
		recipe := hideout.ScavCase[index]
		return &recipe
	}
	fmt.Println("ScavCase recipe ", rid, " does not exist")
	return nil
}

// setHideoutScavcase sets the hideout scavcase items and their indexes.
func setHideoutScavcase(scavcase []map[string]interface{}) {
	hideout.ScavCase = make([]map[string]interface{}, 0, len(scavcase))
	hideout.Index.ScavCase = make(map[string]int8)

	for index, item := range scavcase {
		pid := item["_id"].(string)

		hideout.ScavCase = append(hideout.ScavCase, item)
		hideout.Index.ScavCase[pid] = int8(index)
	}
}
