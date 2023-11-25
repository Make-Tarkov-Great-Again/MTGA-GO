package data

import (
	"fmt"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const (
	areasPath           = hideoutPath + "areas.json"
	productionPath      = hideoutPath + "production.json"
	qtePath             = hideoutPath + "qte.json"
	scavcasePath        = hideoutPath + "scavcase.json"
	hideoutSettingsPath = hideoutPath + "settings.json"

	areasNotExist          = "Hideout Areas does not exist"
	areaNotExist           = "Hideout Area Type %s does not exist"
	qteNotExist            = "Hideout QTE does not exist"
	settingsNotExist       = "Hideout Settings does not exist"
	recipesNotExist        = "Hideout Recipes does not exist"
	recipeNotExist         = "Hideout Recipe %s does not exist"
	scavCaseNotExist       = "Hideout ScavCase does not exist"
	scavCaseRecipeNotExist = "ScavCase recipe %s does not exist"
)

// #region Hideout getters

func GetHideoutAreas() ([]map[string]any, error) {
	if db.hideout.Areas != nil {
		return db.hideout.Areas, nil
	}

	return nil, fmt.Errorf(areasNotExist)
}

func GetHideoutQTE() ([]map[string]any, error) {
	if db.hideout.QTE != nil {
		return db.hideout.QTE, nil
	}

	return nil, fmt.Errorf(qteNotExist)
}

func GetHideoutSettings() (*HideoutSettings, error) {
	if db.hideout.Settings != nil {
		return db.hideout.Settings, nil
	}

	return nil, fmt.Errorf(settingsNotExist)
}

func GetHideoutRecipes() ([]map[string]any, error) {
	if db.hideout.Recipes != nil {
		return db.hideout.Recipes, nil
	}

	return nil, fmt.Errorf(recipesNotExist)
}

func GetHideoutScavcase() ([]map[string]any, error) {
	if db.hideout.ScavCase != nil {
		return db.hideout.ScavCase, nil
	}

	return nil, fmt.Errorf(scavCaseNotExist)
}

// GetHideoutAreaByAreaType retrieves a hideout area by its type int8.
func GetHideoutAreaByAreaType(_type int8) *map[string]any {
	index, ok := db.hideout.Index.Areas[_type]
	if !ok {
		log.Println(areaNotExist, _type)
		return nil
	}

	hideoutArea := db.hideout.Areas[index]
	return &hideoutArea
}

// GetHideoutAreaByName retrieves a hideout area by its name.
func GetHideoutAreaByName(name string) *map[string]any {
	area, ok := HideoutAreaNames[name]
	if !ok {
		log.Println(areaNotExist, name)
		return nil
	}

	index, ok := db.hideout.Index.Areas[area]
	if !ok {
		log.Println(areaNotExist, area)
		return nil
	}

	hideoutArea := db.hideout.Areas[index]
	return &hideoutArea
}

// GetHideoutRecipeByID retrieves a hideout production by its ID.
func GetHideoutRecipeByID(rid string) *map[string]any {
	index, ok := db.hideout.Index.Recipes[rid]
	if ok {
		recipe := db.hideout.Recipes[index]
		return &recipe
	}
	log.Println(recipeNotExist, rid)
	return nil
}

// GetScavCaseRecipeByID retrieves a scavcase production by its ID.
func GetScavCaseRecipeByID(rid string) *map[string]any {
	index, ok := db.hideout.Index.ScavCase[rid]
	if ok {
		recipe := db.hideout.ScavCase[index]
		return &recipe
	}
	log.Println(scavCaseRecipeNotExist, rid)
	return nil
}

// #endregion

// #region Hideout setters

// setHideout sets in-memory database entries for Hideout
func setHideout() {
	db.hideout = &Hideout{
		Index: HideoutIndex{
			Areas:    make(map[int8]int8),
			ScavCase: make(map[string]int8),
			Recipes:  make(map[string]int16),
		},
		Areas:    make([]map[string]any, 0),
		Recipes:  make([]map[string]any, 0),
		QTE:      make([]map[string]any, 0),
		ScavCase: make([]map[string]any, 0),
		Settings: new(HideoutSettings),
	}

	done := make(chan struct{})

	go func() {
		areas := tools.GetJSONRawMessage(areasPath)
		if err := json.UnmarshalNoEscape(areas, &db.hideout.Areas); err != nil {
			log.Println(err)
		}
		done <- struct{}{}
	}()

	go func() {
		recipes := tools.GetJSONRawMessage(productionPath)
		if err := json.UnmarshalNoEscape(recipes, &db.hideout.Recipes); err != nil {
			log.Println(err)
		}
		done <- struct{}{}
	}()

	go func() {
		scavcase := tools.GetJSONRawMessage(scavcasePath)
		if err := json.UnmarshalNoEscape(scavcase, &db.hideout.ScavCase); err != nil {
			log.Println(err)
		}
		done <- struct{}{}
	}()

	go func() {
		qte := tools.GetJSONRawMessage(qtePath)
		if err := json.UnmarshalNoEscape(qte, &db.hideout.QTE); err != nil {
			log.Println(err)
		}
		done <- struct{}{}
	}()

	go func() {
		settings := tools.GetJSONRawMessage(hideoutSettingsPath)
		if err := json.UnmarshalNoEscape(settings, &db.hideout.Settings); err != nil {
			log.Println(err)
		}
		done <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		<-done
	}
}

func setHideoutAreaLookup() {
	for index, area := range db.hideout.Areas {
		areaType := int8(area["type"].(float64))
		db.hideout.Index.Areas[areaType] = int8(index)
	}
}

func setHideoutRecipeLookup() {
	for index, recipe := range db.hideout.Recipes {
		pid := recipe["_id"].(string)
		db.hideout.Index.Recipes[pid] = int16(index)
	}
}

func setScavcaseRecipeLookup() {
	for index, item := range db.hideout.ScavCase {
		pid := item["_id"].(string)
		db.hideout.Index.ScavCase[pid] = int8(index)
	}
}

// #endregion

// #region Hideout structs

type Hideout struct {
	Index    HideoutIndex
	Areas    []map[string]any
	Recipes  []map[string]any
	QTE      []map[string]any
	ScavCase []map[string]any
	Settings *HideoutSettings
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
