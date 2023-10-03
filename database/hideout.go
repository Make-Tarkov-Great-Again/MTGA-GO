package database

import (
	"fmt"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var hideout = Hideout{}

const (
	areasPath           string = hideoutPath + "areas.json"
	productionPath      string = hideoutPath + "production.json"
	qtePath             string = hideoutPath + "qte.json"
	scavcasePath        string = hideoutPath + "scavcase.json"
	hideoutSettingsPath string = hideoutPath + "settings.json"
)

// #region Hideout getters

// GetHideout retrieves the current hideout configuration.
func GetHideout() *Hideout {
	return &hideout
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

	index, ok := hideout.Index.Areas[area]
	if !ok {
		fmt.Println("Area Type ", area, " does not exist")
		return nil
	}

	hideoutArea := hideout.Areas[index]
	return &hideoutArea
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

// GetScavCaseRecipeByID retrieves a scavcase production by its ID.
func GetScavCaseRecipeByID(rid string) *map[string]interface{} {
	index, ok := hideout.Index.ScavCase[rid]
	if ok {
		recipe := hideout.ScavCase[index]
		return &recipe
	}
	fmt.Println("ScavCase recipe ", rid, " does not exist")
	return nil
}

// #endregion

// #region Hideout setters

// setHideoutScavcase sets the hideout scavcase items and their indexes.
func setHideout() {

	if tools.FileExist(areasPath) {
		areas := tools.GetJSONRawMessage(areasPath)
		var areasMap []map[string]interface{}
		err := json.Unmarshal(areas, &areasMap)
		if err != nil {
			log.Fatalln(err)
		}
		setHideoutAreas(areasMap)
	}

	if tools.FileExist(productionPath) {
		recipes := tools.GetJSONRawMessage(productionPath)
		var productionsMap []map[string]interface{}
		err := json.Unmarshal(recipes, &productionsMap)
		if err != nil {
			log.Fatalln(err)
		}
		setHideoutRecipes(productionsMap)
	}

	if tools.FileExist(scavcasePath) {
		scavcase := tools.GetJSONRawMessage(scavcasePath)
		var scavcaseReturns []map[string]interface{}
		err := json.Unmarshal(scavcase, &scavcaseReturns)
		if err != nil {
			log.Fatalln(err)
		}
		setHideoutScavcase(scavcaseReturns)
	}

	if tools.FileExist(qtePath) {
		qte := tools.GetJSONRawMessage(qtePath)
		hideout.QTE = []map[string]interface{}{}

		err := json.Unmarshal(qte, &hideout.QTE)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if tools.FileExist(hideoutSettingsPath) {
		settings := tools.GetJSONRawMessage(hideoutSettingsPath)
		err := json.Unmarshal(settings, &hideout.Settings)
		if err != nil {
			log.Fatalln(err)
		}
	}
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

// setHideoutRecipes sets the hideout production recipes and their indexes.
func setHideoutRecipes(recipes []map[string]interface{}) {
	hideout.Recipes = make([]map[string]interface{}, 0, len(recipes))
	hideout.Index.Recipes = make(map[string]int16)

	for index, recipe := range recipes {
		pid := recipe["_id"].(string)

		hideout.Index.Recipes[pid] = int16(index)
		hideout.Recipes = append(hideout.Recipes, recipe)
	}
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

// #endregion

// #region Hideout structs

type Hideout struct {
	Index    HideoutIndex
	Areas    []map[string]interface{}
	Recipes  []map[string]interface{}
	QTE      []map[string]interface{}
	ScavCase []map[string]interface{}
	Settings HideoutSettings
}

type HideoutIndex struct {
	Areas    map[int8]int8
	ScavCase map[string]int8
	Recipes  map[string]int16
}

type HideoutSettings struct {
	GeneratorSpeedWithoutFuel float64 `json:"generatorSpeedWithoutFuel"`
	GeneratorFuelFlowRate     float64 `json:"generatorFuelFlowRate"`
	AirFilterUnitFlowRate     float64 `json:"airFilterUnitFlowRate"`
	GPUBoostRate              float64 `json:"gpuBoostRate"`
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

// #endregion
