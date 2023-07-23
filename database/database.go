package database

const DATABASE_LIB_PATH string = "database/lib"

/*
import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

type DatabaseStruct struct {
	Core          *CoreStruct
	Connections   *ConnectionStruct
	Items         map[string]map[string]interface{}
	Locales       *LocaleStruct
	Templates     *TemplatesStruct
	Traders       map[string]TraderStruct
	Flea          *FleaStruct
	Quests        map[string]map[string]interface{}
	Hideout       *HideoutStruct
	Locations     *LocationsStruct
	Weather       map[string]interface{}
	Customization map[string]map[string]interface{}
	Editions      map[string]interface{}
	Bot           *BotStruct
	Profiles      map[string]ProfileStruct
	//bundles  []map[string]interface{}
}

type CoreStruct struct {
	BotTemplate    map[string]interface{}
	ClientSettings map[string]interface{}
	ServerConfig   map[string]interface{}
	Globals        map[string]interface{}
	Locations      map[string]interface{}
	Presets        map[string]interface{}
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics map[string]interface{}
}

type ConnectionStruct struct {
	WebSocket      map[string]interface{}
	WebSocketPings map[string]interface{}
}

type FleaStruct struct {
	Offers           []map[string]interface{}
	Offerscount      int
	SelectedCategory string
	Categories       map[string]interface{}
}

type HideoutStruct struct {
	Areas       []map[string]interface{}
	Productions []map[string]interface{}
	Scavcase    []map[string]interface{}
	QTE         []map[string]interface{}
	Settings    map[string]interface{}
}

type LocationsStruct struct {
	Locations map[string]LocationStruct
	LootGen   *LootGenStruct
}

var database = &DatabaseStruct{}

func GetDatabase() *DatabaseStruct {
	return database
}

func InitializeDatabase() error {
	database.Core = &CoreStruct{
		BotTemplate:    map[string]interface{}{},
		ClientSettings: map[string]interface{}{},
		ServerConfig:   map[string]interface{}{},
		Globals:        map[string]interface{}{},
		Locations:      map[string]interface{}{},
		MatchMetrics:   map[string]interface{}{},
		Presets:        map[string]interface{}{},
	}
	database.Connections = &ConnectionStruct{
		WebSocket:      map[string]interface{}{},
		WebSocketPings: map[string]interface{}{},
	}
	database.Items = setItems()
	database.Locales = &LocaleStruct{
		Locales:   make(map[string]LanguageStruct),
		Extras:    map[string]interface{}{},
		Languages: map[string]interface{}{},
	}
	database.Templates = &TemplatesStruct{
		Handbook: HandbookStruct{
			Items:      []map[string]interface{}{},
			Categories: []map[string]interface{}{},
		},
		Prices: map[string]interface{}{},
		TplLookup: TplLookupStruct{
			Items: ItemsLookupStruct{
				ById:     map[string]interface{}{},
				ByParent: map[string]interface{}{},
			},
			Categories: CategoriesLookupStruct{
				ById:     map[string]interface{}{},
				ByParent: map[string]interface{}{},
			},
		},
	}
	database.Editions = map[string]interface{}{}
	database.Traders = map[string]TraderStruct{}
	database.Quests = setQuests()
	database.Flea = &FleaStruct{
		Offers:           []map[string]interface{}{},
		Offerscount:      0,
		SelectedCategory: "",
		Categories:       map[string]interface{}{},
	}
	database.Hideout = &HideoutStruct{
		Areas:       []map[string]interface{}{},
		Productions: []map[string]interface{}{},
		Scavcase:    []map[string]interface{}{},
		QTE:         []map[string]interface{}{},
		Settings:    map[string]interface{}{},
	}
	database.Customization = setCustomization()
	database.Profiles = map[string]ProfileStruct{}
	database.Weather = tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "weather.json"))

	database.Bot = &BotStruct{
		Bots:        map[string]interface{}{},
		Core:        map[string]interface{}{},
		Names:       map[string]interface{}{},
		Appearance:  map[string]interface{}{},
		PlayerScav:  map[string]interface{}{},
		WeaponCache: map[string]interface{}{},
	}
	database.Locations = &LocationsStruct{
		Locations: map[string]LocationStruct{},
		LootGen: &LootGenStruct{
			Containers: map[string]interface{}{},
			Static:     map[string]interface{}{},
		},
	}

	if err := setDatabase(); err != nil {
		return fmt.Errorf("error setting database: %w", err)
	}
	return nil
}


func setDatabase() error {
	if err := setDatabaseCore(); err != nil {
		return err
	}

	if err := setEditions(); err != nil {
		return err
	}

	if err := setLocales(); err != nil {
		return err
	}

	if err := setTemplates(); err != nil {
		return err
	}

	if err := setTraders(); err != nil {
		return err
	}

	if err := setHideout(); err != nil {
		return err
	}

	if err := setProfiles(); err != nil {
		return err
	}

	if err := setBot(); err != nil {
		return err
	}

	if err := setLocations(); err != nil {
		return err
	}

	return nil
}

func setItems() map[string]map[string]interface{} {
	items := tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "items.json"))
	itemsMap := make(map[string]map[string]interface{}, len(items))

	for id, quest := range items {
		itemsMap[id] = quest.(map[string]interface{})
	}

	return itemsMap
}

func setCustomization() map[string]map[string]interface{} {
	customizations := tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "customization.json"))
	customizationsMap := make(map[string]map[string]interface{}, len(customizations))

	for id, customization := range customizations {
		customizationsMap[id] = customization.(map[string]interface{})
	}

	return customizationsMap
}

func setQuests() map[string]map[string]interface{} {
	quests := tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "quests.json"))
	questsMap := make(map[string]map[string]interface{}, len(quests))

	for id, quest := range quests {
		questsMap[id] = quest.(map[string]interface{})
	}

	return questsMap
}

func setDatabaseCore() error {
	core := database.Core

	if err := setServerConfigCore(core); err != nil {
		return fmt.Errorf("error setting server config: %w", err)
	}

	core.MatchMetrics = tools.SetProperObjectDataStructure(MATCH_METRICS_PATH)

	core.Globals = tools.SetProperObjectDataStructure(GLOBALS_FILE_PATH)

	if err := setPresetsCore(core); err != nil {
		return fmt.Errorf("error setting presets: %w", err)
	}
	core.ClientSettings = tools.SetProperObjectDataStructure(CLIENT_SETTINGS_PATH)

	core.Locations = tools.SetProperObjectDataStructure(LOCATIONS_FILE_PATH)

	if err := setBotTemplateCore(core); err != nil {
		return fmt.Errorf("error setting bot template: %w", err)
	}

	return nil
}

func setServerConfigCore(core *CoreStruct) error {
	serverConfig, err := tools.ReadParsed(SERVER_CONFIG_PATH)

	if err != nil {
		return fmt.Errorf("error reading server.json: %w", err)
	}

	serverConfigMap, ok := serverConfig.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in server.json")
	}

	core.ServerConfig = serverConfigMap
	return nil
}

func setPresetsCore(core *CoreStruct) error {
	presets, ok := core.Globals["ItemPresets"]
	if !ok {
		return fmt.Errorf("error reading ItemPresets from globals")
	}

	globalPresets, ok := presets.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in ItemPresets")
	}

	for id, value := range globalPresets {
		preset, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in ItemPresets @ value")
		}

		// Accessing the _items array from preset
		items, ok := preset["_items"].([]interface{})
		if !ok || len(items) == 0 {
			return fmt.Errorf("no items found in preset %s", id)
		}

		// Accessing the first item of the _items array as a map
		item, ok := items[0].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid item found in preset %s", id)
		}

		// Accessing the _tpl string from the first item map
		tpl, ok := item["_tpl"].(string)
		if !ok {
			return fmt.Errorf("tpl not found in preset %s", id)
		}

		if _, ok := core.Presets[tpl]; !ok {
			core.Presets[tpl] = map[string]interface{}{}
		}

		presetId, ok := preset["_id"].(string)
		if !ok {
			return fmt.Errorf("presetId not found in preset %s", id)
		}

		itemPreset, ok := core.Presets[tpl].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in ItemPresets")
		}

		itemPreset[presetId] = preset

	}
	return nil
}

func setBotTemplateCore(core *CoreStruct) error {
	botTemplate, err := tools.ReadParsed(BOT_TEMPLATE_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading botTemplate.json: %w", err)
	}

	botTemplateMap, ok := botTemplate.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in botTemplate.json")
	}

	core.BotTemplate = botTemplateMap
	return nil
}

const EDITIONS_FILE_PATH string = DATABASE_LIB_PATH + "/editions"

type EditionStruct struct {
	BEAR    map[string]interface{}
	USEC    map[string]interface{}
	Storage map[string]interface{}
}

func setEditions() error {
	editionsDirectory, err := tools.GetDirectoriesFrom(EDITIONS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading editions directory: %w", err)
	}

	for _, edition := range editionsDirectory {
		editionPath := filepath.Join(EDITIONS_FILE_PATH, edition)

		// Read character_bear.json and store in bear map
		bearData, err := tools.ReadParsed(filepath.Join(editionPath, "character_bear.json"))
		if err != nil {
			return fmt.Errorf("error reading character_bear.json for edition %s: %w", edition, err)
		}

		bearMap, ok := bearData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in character_bear.json for edition %s", edition)
		}

		// Read character_usec.json and store in usec map
		usecData, err := tools.ReadParsed(filepath.Join(editionPath, "character_usec.json"))
		if err != nil {
			return fmt.Errorf("error reading character_usec.json for edition %s: %w", edition, err)
		}

		usecMap, ok := usecData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in character_usec.json for edition %s", edition)
		}

		// Read storage.json and store in storage map
		storageData, err := tools.ReadParsed(filepath.Join(editionPath, "storage.json"))
		if err != nil {
			return fmt.Errorf("error reading storage.json for edition %s: %w", edition, err)
		}

		storageMap, ok := storageData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in storage.json for edition %s", edition)
		}

		// Create an EditionStruct for this edition and populate the maps
		editionData := EditionStruct{
			BEAR:    bearMap,
			USEC:    usecMap,
			Storage: storageMap,
		}

		// Add the editionData to the database.editions map
		database.Editions[edition] = editionData
	}
	return nil
}

const (
	LOCALES_FILE_PATH        string = DATABASE_LIB_PATH + "/locales"
	LOCALES_EXTRAS_FILE_PATH string = LOCALES_FILE_PATH + "/extras.json"
	LOCALES_LANGUAGES_PATH   string = LOCALES_FILE_PATH + "/languages.json"
)

type LocaleStruct struct {
	Locales   map[string]LanguageStruct
	Extras    map[string]interface{}
	Languages map[string]interface{}
}
type LanguageStruct struct {
	Locale map[string]interface{}
	Menu   map[string]interface{}
}

func setLocales() error {
	localesDirectory, err := tools.GetDirectoriesFrom(LOCALES_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading locales directory: %w", err)
	}

	locales := database.Locales

	locales.Languages = tools.SetProperObjectDataStructure(LOCALES_LANGUAGES_PATH)
	locales.Extras = tools.SetProperObjectDataStructure(LOCALES_EXTRAS_FILE_PATH)

	dataChan := make(chan LanguageStruct, len(localesDirectory))
	for _, locale := range localesDirectory {
		go func(locale string) {

			var data LanguageStruct
			localePath := filepath.Join(LOCALES_FILE_PATH, locale, "locale.json")
			menuPath := filepath.Join(LOCALES_FILE_PATH, locale, "menu.json")

			data.Locale = tools.SetProperObjectDataStructure(localePath)
			data.Menu = tools.SetProperObjectDataStructure(menuPath)

			// Add the localeData to the locales map
			dataChan <- data
		}(locale)
	}

	for i := 0; i < len(localesDirectory); i++ {
		data := <-dataChan
		locales.Locales[localesDirectory[i]] = LanguageStruct{
			Locale: data.Locale,
			Menu:   data.Menu,
		}

	}
	return nil
}

type TemplatesStruct struct {
	Handbook  HandbookStruct
	Prices    map[string]interface{}
	TplLookup TplLookupStruct
}

type HandbookStruct struct {
	Items      []map[string]interface{}
	Categories []map[string]interface{}
}

type TplLookupStruct struct {
	Items      ItemsLookupStruct
	Categories CategoriesLookupStruct
}

type ItemsLookupStruct struct {
	ById     map[string]interface{}
	ByParent map[string]interface{}
}

type CategoriesLookupStruct struct {
	ById     map[string]interface{}
	ByParent map[string]interface{}
}

func setTemplates() error {
	templates := database.Templates
	handbook := tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "templates.json"))

	items := tools.TransformInterfaceIntoMappedArray(handbook["Items"].([]interface{}))
	categories := tools.TransformInterfaceIntoMappedArray(handbook["Categories"].([]interface{}))

	setHandbookItems(items, templates)
	setHandbookCategories(categories, templates)

	return nil
}

func setHandbookCategories(categories []map[string]interface{}, templates *TemplatesStruct) error {
	templates.Handbook.Categories = tools.AuditArrayCapacity(categories)

	byCategory := &templates.TplLookup.Categories // pointer to the CategoriesLookupStruct
	for _, category := range templates.Handbook.Categories {
		Id := category["Id"].(string)
		ParentId, hasParent := category["ParentId"].(string)

		byCategory.ById[Id] = ParentId
		if hasParent {
			if mapSlice, ok := byCategory.ByParent[ParentId].([]string); ok {
				byCategory.ByParent[ParentId] = append(mapSlice, Id)
			} else {
				byCategory.ByParent[ParentId] = []string{Id}
			}
		}
	}
	return nil
}

func setHandbookItems(items []map[string]interface{}, templates *TemplatesStruct) error {
	templates.Handbook.Items = tools.AuditArrayCapacity(items)

	prices := tools.SetProperObjectDataStructure(filepath.Join(DATABASE_LIB_PATH, "liveflea.json"))

	byItem := &templates.TplLookup.Items // pointer to the ItemsLookupStruct
	for _, item := range templates.Handbook.Items {
		Id := item["Id"].(string)
		ParentId := item["ParentId"].(string)

		// set the item price to the price from liveflea.json if it exists, otherwise use the item from items.json
		var price interface{} // store price in this variable
		if p, ok := prices[Id]; ok {
			price = p
		} else {
			price = item
		}
		byItem.ById[Id] = price
		templates.Prices[Id] = price

		// add the item to the byParent map
		if mapSlice, ok := byItem.ByParent[ParentId].([]string); ok {
			// if the key exists, append the value to the slice
			byItem.ByParent[ParentId] = append(mapSlice, Id)
		} else {
			byItem.ByParent[ParentId] = []string{Id}
		}
	}

	return nil
}

const TRADERS_FILE_PATH string = DATABASE_LIB_PATH + "/traders"

type TraderStruct struct {
	Assort      AssortStruct
	Base        map[string]interface{}
	BaseAssort  map[string]interface{}
	Dialogue    map[string]interface{}
	Questassort map[string]interface{}
	Suits       map[string]interface{}
}

func setTraders() error {
	tradersDirectory, err := tools.GetDirectoriesFrom(TRADERS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading traders directory: %w", err)
	}

	type traderData struct {
		id         string
		assort     AssortStruct
		base       map[string]interface{}
		baseassort map[string]interface{}
		dialogue   map[string]interface{}
		quest      map[string]interface{}
		suits      map[string]interface{}
	}
	var wg sync.WaitGroup
	traderDataChan := make(chan traderData, len(tradersDirectory))

	// Launch goroutines to process each trader concurrently
	for _, traderID := range tradersDirectory {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			var traderData traderData

			traderData.id = id
			if assort, err := setTraderAssort(filepath.Join(TRADERS_FILE_PATH, id)); err != nil {
				log.Printf("Error setting trader assort for %s: %v\n", id, err)
			} else {
				traderData.assort = assort
			}

			traderData.base = tools.SetProperObjectDataStructure(filepath.Join(TRADERS_FILE_PATH, id, "base.json"))

			if dialogue, _ := setTraderDialogue(filepath.Join(TRADERS_FILE_PATH, id)); dialogue != nil {
				traderData.dialogue = dialogue
			}

			if questAssort, _ := setTraderQuestAssort(filepath.Join(TRADERS_FILE_PATH, id)); questAssort != nil {
				traderData.quest = questAssort
			}

			if suits, _ := setTraderSuits(filepath.Join(TRADERS_FILE_PATH, id)); suits != nil {
				traderData.suits = suits
			}

			traderDataChan <- traderData
		}(traderID)
	}

	// Wait for all goroutines to finish and update database
	go func() {
		wg.Wait()
		close(traderDataChan)
	}()

	for traderData := range traderDataChan {
		database.Traders[traderData.id] = TraderStruct{traderData.assort, traderData.base, traderData.baseassort, traderData.dialogue, traderData.quest, traderData.suits}
	}

	return nil
}

type AssortStruct struct {
	items             []map[string]interface{}
	barter_scheme     map[string]interface{}
	loyal_level_items map[string]interface{}
}

func setTraderAssort(path string) (AssortStruct, error) {
	data, err := tools.ReadParsed(filepath.Join(path, "assort.json"))
	if err != nil {
		return AssortStruct{}, fmt.Errorf("error reading %s: %w", path, err)
	}

	assort, ok := data.(map[string]interface{})
	if !ok {
		return AssortStruct{}, fmt.Errorf("invalid data structure in assort @ %s", path)
	}

	if assortData, ok := assort["data"].(map[string]interface{}); ok {
		assort = assortData
	}

	itemsData, ok := assort["items"].([]interface{})
	if !ok {
		return AssortStruct{}, fmt.Errorf("invalid data structure in itemsData @ %s", path)
	}
	items := tools.TransformInterfaceIntoMappedArray(itemsData)

	barter_scheme, ok := assort["barter_scheme"].(map[string]interface{})
	if !ok {
		return AssortStruct{}, fmt.Errorf("invalid data structure in barter_scheme @ %s", path)
	}

	loyal_level_items, ok := assort["loyal_level_items"].(map[string]interface{})
	if !ok {
		return AssortStruct{}, fmt.Errorf("invalid data structure in loyal_level_items @ %s", path)
	}

	return AssortStruct{
		items:             items,
		barter_scheme:     barter_scheme,
		loyal_level_items: loyal_level_items,
	}, nil
}

func setTraderDialogue(path string) (map[string]interface{}, error) {
	data, err := tools.ReadParsed(filepath.Join(path, "dialogue.json"))
	if err != nil {
		return nil, fmt.Errorf("error reading dialogue.json: %w", err)
	}

	dialogue, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in dialogue.json")
	}

	if dialogueData, ok := dialogue["data"].(map[string]interface{}); ok {
		dialogue = dialogueData
	}

	return dialogue, nil
}

func setTraderQuestAssort(path string) (map[string]interface{}, error) {
	data, err := tools.ReadParsed(filepath.Join(path, "questassort.json"))
	if err != nil {
		return nil, fmt.Errorf("error reading questassort.json: %w", err)
	}

	questassort, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in questassort.json")
	}

	if questassortData, ok := questassort["data"].(map[string]interface{}); ok {
		questassort = questassortData
	}

	return questassort, nil
}

func setTraderSuits(path string) (map[string]interface{}, error) {
	data, err := tools.ReadParsed(filepath.Join(path, "suits.json"))
	if err != nil {
		return nil, fmt.Errorf("error reading suits.json: %w", err)
	}

	suits, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in suits.json")
	}

	if suitsData, ok := suits["data"].(map[string]interface{}); ok {
		suits = suitsData
	}

	return suits, nil
}

const HIDEOUT_FILE_PATH string = DATABASE_LIB_PATH + "/hideout"
const (
	AREAS_FILE_PATH       string = HIDEOUT_FILE_PATH + "/areas.json"
	PRODUCTIONS_FILE_PATH string = HIDEOUT_FILE_PATH + "/productions.json"
	SCAVCASE_FILE_PATH    string = HIDEOUT_FILE_PATH + "/scavcase.json"
	QTE_FILE_PATH         string = HIDEOUT_FILE_PATH + "/qte.json"
	SETTINGS_FILE_PATH    string = HIDEOUT_FILE_PATH + "/settings.json"
)

func setHideout() error {

	areas, err := setHideoutAreas()
	if err != nil {
		return err
	}

	productions, err := setHideoutProductions()
	if err != nil {
		return err
	}

	scavcase, err := setHideoutScavcase()
	if err != nil {
		return err
	}

	qte, err := setHideoutQTE()
	if err != nil {
		return err
	}

	settings, err := setHideoutSettings()
	if err != nil {
		return err
	}

	database.Hideout = &HideoutStruct{
		Areas:       areas,
		Productions: productions,
		Scavcase:    scavcase,
		QTE:         qte,
		Settings:    settings,
	}

	return nil
}

func setHideoutSettings() (map[string]interface{}, error) {
	settings, err := tools.ReadParsed(SETTINGS_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading settings.json: %w", err)
	}

	settingsMap, ok := settings.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in settings.json")
	}

	settingsData, ok := settingsMap["data"].(map[string]interface{})
	if !ok {
		return settingsMap, nil
	} else {
		return settingsData, nil
	}
}

func setHideoutProductions() ([]map[string]interface{}, error) {
	data, err := tools.ReadParsed(PRODUCTIONS_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading productions.json: %w", err)
	}

	productionMap, ok := data.(map[string]interface{})
	if !ok {
		productionArray, ok := data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data structure in productionMap")
		}
		return tools.TransformInterfaceIntoMappedArray(productionArray), nil
	}
	productionData, ok := productionMap["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in productionData")
	}
	return tools.TransformInterfaceIntoMappedArray(productionData), nil
}

func setHideoutAreas() ([]map[string]interface{}, error) {
	data, err := tools.ReadParsed(AREAS_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading areas.json: %w", err)
	}

	areasMap, ok := data.(map[string]interface{})
	if !ok {
		areasArray, ok := data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data structure in areasMap")
		}
		return tools.TransformInterfaceIntoMappedArray(areasArray), nil
	}
	areasData, ok := areasMap["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in areasData")
	}
	return tools.TransformInterfaceIntoMappedArray(areasData), nil
}

func setHideoutScavcase() ([]map[string]interface{}, error) {
	data, err := tools.ReadParsed(SCAVCASE_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading scavcase.json: %w", err)
	}

	scavcaseMap, ok := data.(map[string]interface{})
	if !ok {
		scavcaseArray, ok := data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data structure in scavcaseMap")
		}
		return tools.TransformInterfaceIntoMappedArray(scavcaseArray), nil
	}
	scavcaseData, ok := scavcaseMap["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in scavcaseData")
	}
	return tools.TransformInterfaceIntoMappedArray(scavcaseData), nil
}

func setHideoutQTE() ([]map[string]interface{}, error) {
	data, err := tools.ReadParsed(QTE_FILE_PATH)
	if err != nil {
		return nil, fmt.Errorf("error reading qte.json: %w", err)
	}

	qteMap, ok := data.(map[string]interface{})
	if !ok {
		qteArray, ok := data.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid data structure in qteMap")
		}
		return tools.TransformInterfaceIntoMappedArray(qteArray), nil
	}
	qteData, ok := qteMap["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in qteData")
	}
	return tools.TransformInterfaceIntoMappedArray(qteData), nil
}

const USER_FILE_PATH string = "user"
const PROFILES_FILE_PATH string = USER_FILE_PATH + "/profiles"

type ProfileStruct struct {
	Account   map[string]interface{}
	Character map[string]interface{}
	Storage   map[string]interface{}
	Dialogues map[string]interface{}
	Raid      RaidProfileStruct
}
type RaidProfileStruct struct {
	LastLocation RaidLocationStruct
	CarExtracts  int
}
type RaidLocationStruct struct {
	Name      string
	Insurance bool
}

func setProfiles() error {
	if !tools.FileExist(USER_FILE_PATH) || !tools.FileExist(PROFILES_FILE_PATH) {
		if err := tools.CreateDirectory(PROFILES_FILE_PATH); err != nil {
			return fmt.Errorf("error creating profiles directory: %w", err)
		}
		return nil
	}

	profilesDirectory, err := tools.GetDirectoriesFrom(PROFILES_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading user directory: %w", err)
	} else if profilesDirectory == nil {
		log.Printf("No profiles found in %s", PROFILES_FILE_PATH)
		return nil
	}

	database.Profiles = make(map[string]ProfileStruct, len(profilesDirectory))

	for _, profileID := range profilesDirectory {
		profilePath := filepath.Join(PROFILES_FILE_PATH, profileID)
		profile := ProfileStruct{
			Account:   setAccount(profilePath, profileID),
			Character: setCharacter(profilePath, profileID),
			Storage:   setStorage(profilePath, profileID),
			Dialogues: setDialogues(profilePath, profileID),
			Raid: RaidProfileStruct{
				LastLocation: RaidLocationStruct{
					Name:      "",
					Insurance: false,
				},
				CarExtracts: 0,
			},
		}
		database.Profiles[profileID] = profile
	}
	return nil
}

func setAccount(path string, profileID string) map[string]interface{} {
	accountPath := filepath.Join(path, "account.json")
	data, err := tools.ReadParsed(accountPath)
	if err != nil {
		log.Printf("Error reading account.json for profile %s: %v", profileID, err)
		return nil
	}

	account, ok := data.(map[string]interface{})
	if !ok {
		log.Printf("Account.json for profile %s has invalid structure", profileID)
		return nil
	}

	if len(account) == 0 {
		log.Printf("Account.json for profile %s is empty", profileID)
		return nil
	}

	return account
}

func setCharacter(path string, profileID string) map[string]interface{} {
	characterPath := filepath.Join(path, "character.json")
	data, err := tools.ReadParsed(characterPath)
	if err != nil {
		log.Printf("Error reading character.json for profile %s: %v", profileID, err)
		return nil
	}

	character, ok := data.(map[string]interface{})
	if !ok {
		log.Printf("Character.json for profile %s has invalid structure", profileID)
		return nil
	}

	if len(character) == 0 {
		log.Printf("Character.json for profile %s is empty", profileID)
		return nil
	}

	return character
}

func setStorage(path string, profileID string) map[string]interface{} {
	storagePath := filepath.Join(path, "storage.json")
	data, err := tools.ReadParsed(storagePath)
	if err != nil {
		log.Printf("Error reading storage.json for profile %s: %v", profileID, err)
		return nil
	}

	storage, ok := data.(map[string]interface{})
	if !ok {
		log.Printf("Storage.json for profile %s has invalid structure", profileID)
		return nil
	}

	if len(storage) == 0 {
		log.Printf("Storage.json for profile %s is empty", profileID)
		return nil
	}

	return storage
}

func setDialogues(path string, profileID string) map[string]interface{} {
	dialoguesPath := filepath.Join(path, "dialogues.json")
	data, err := tools.ReadParsed(dialoguesPath)
	if err != nil {
		log.Printf("Error reading dialogues.json for profile %s: %v", profileID, err)
		return nil
	}

	dialogues, ok := data.(map[string]interface{})
	if !ok {
		log.Printf("Dialogues.json for profile %s has invalid structure", profileID)
		return nil
	}

	if len(dialogues) == 0 {
		log.Printf("Dialogues.json for profile %s is empty", profileID)
		return nil
	}

	return dialogues
}

const BOT_FILE_PATH = DATABASE_LIB_PATH + "/bot"

type BotStruct struct {
	Bots        map[string]interface{}
	Core        map[string]interface{}
	Names       map[string]interface{}
	Appearance  map[string]interface{}
	PlayerScav  map[string]interface{}
	WeaponCache map[string]interface{}
}

func setBot() error {
	bot := database.Bot

	if err := setBotCore(bot); err != nil {
		return err
	}

	if err := setBotNames(bot); err != nil {
		return err
	}

	if err := setBotAppearance(bot); err != nil {
		return err
	}

	if err := setBotPlayerScav(bot); err != nil {
		return err
	}

	if err := setBotWeaponCache(bot); err != nil {
		return err
	}

	if err := setBots(bot); err != nil {
		return err
	}

	return nil
}

func setBotCore(bot *BotStruct) error {
	data, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "__BotGlobalSettings.json"))
	if err != nil {
		return fmt.Errorf("error reading __BotGlobalSettings.json: %w", err)
	}

	botGlobals, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in __BotGlobalSettings.json")
	}

	bot.Core = botGlobals
	return nil
}

func setBotNames(bot *BotStruct) error {
	data, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "names.json"))
	if err != nil {
		return fmt.Errorf("error reading names.json: %w", err)
	}

	names, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in names.json")
	}

	bot.Names = names
	return nil
}

func setBotAppearance(bot *BotStruct) error {
	data, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "appearance.json"))
	if err != nil {
		return fmt.Errorf("error reading appearance.json: %w", err)
	}

	appearance, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in appearance.json")
	}

	bot.Appearance = appearance
	return nil
}

func setBotPlayerScav(bot *BotStruct) error {
	data, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "playerScav.json"))
	if err != nil {
		return fmt.Errorf("error reading playerScav.json: %w", err)
	}

	playerScav, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in playerScav.json")
	}

	bot.PlayerScav = playerScav
	return nil
}

func setBotWeaponCache(bot *BotStruct) error {
	data, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "weaponCache.json"))
	if err != nil {
		return fmt.Errorf("error reading weaponCache.json: %w", err)
	}

	weaponCache, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in weaponCache.json")
	}

	bot.WeaponCache = weaponCache
	return nil
}

const BOTS_FILE_PATH string = BOT_FILE_PATH + "/bots"

type BotTypeStruct struct {
	health     map[string]interface{}
	loadout    map[string]interface{}
	difficulty map[string]interface{}
}

func setBots(bot *BotStruct) error {
	botsFiles, err := tools.GetDirectoriesFrom(BOTS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading bots directory: %w", err)
	}

	type botData struct {
		aiType     string
		health     map[string]interface{}
		loadout    map[string]interface{}
		difficulty map[string]interface{}
	}

	var wg sync.WaitGroup
	botDataChan := make(chan botData, len(botsFiles))

	for _, aiType := range botsFiles {
		wg.Add(1)
		go func(aiType string) {
			defer wg.Done()
			var data botData
			botTypePath := filepath.Join(BOTS_FILE_PATH, aiType)

			data.aiType = aiType
			data.health = setBotTypeHealth(botTypePath)
			data.loadout = setBotTypeLoadout(botTypePath)
			data.difficulty = setBotTypeDifficulty(botTypePath)

			botDataChan <- data
		}(aiType)
	}

	go func() {
		wg.Wait()
		close(botDataChan)
	}()

	for data := range botDataChan {
		bot.Bots[data.aiType] = BotTypeStruct{
			health:     data.health,
			loadout:    data.loadout,
			difficulty: data.difficulty,
		}
	}

	return nil
}

func setBotTypeHealth(path string) map[string]interface{} {
	healthFilePath := filepath.Join(path, "health.json")

	healthData, err := tools.ReadParsed(healthFilePath)
	if err != nil || healthData == nil {
		return nil
	}

	if health, ok := healthData.(map[string]interface{}); ok && len(health) == 1 {
		return health
	}

	healthMap := make(map[string]interface{})
	for key, value := range healthData.(map[string]interface{}) {
		if v, ok := value.(map[string]interface{}); ok {
			healthMap[key] = v
		}
	}

	return healthMap
}

func setBotTypeLoadout(path string) map[string]interface{} {
	loadoutFilePath := filepath.Join(path, "loadout.json")

	data, err := tools.ReadParsed(loadoutFilePath)
	if err != nil || data == nil {
		return nil
	}

	if loadout, ok := data.(map[string]interface{}); ok {
		return loadout
	}

	return nil
}

func setBotTypeDifficulty(path string) map[string]interface{} {
	difficultiesPath := filepath.Join(path, "difficulties")

	difficultiesDirectory, err := tools.GetFilesFrom(difficultiesPath)
	if err != nil {
		log.Panicf("error reading difficulties directory: %v", err)
		return nil
	}

	difficulties := make(map[string]interface{}, len(difficultiesDirectory))
	for _, difficulty := range difficultiesDirectory {
		difficultyName := difficulty[:len(difficulty)-5]
		difficultyPath := filepath.Join(difficultiesPath, difficulty)

		difficultyData, err := tools.ReadParsed(difficultyPath)
		if err != nil {
			log.Panicf("error reading %s: %v", difficulty, err)
			return nil
		}
		difficulties[difficultyName] = difficultyData
	}
	return difficulties
}

type LocationStruct struct {
	Base                   map[string]interface{}
	DynamicAvailableSpawns map[string]interface{}
	LootSpawns             map[string]interface{}
	//waves                  []map[string]interface{}
	//bossWaves              []map[string]interface{}
	Presets map[string]map[string]interface{}
}

func setLocations() error {
	locationsDirectory, err := tools.GetDirectoriesFrom(filepath.Join(DATABASE_LIB_PATH, "locations"))
	if err != nil {
		return fmt.Errorf("error reading locations directory: %w", err)
	}

	locations := database.Locations
	locations.Locations = make(map[string]LocationStruct, len(locationsDirectory))

	for _, location := range locationsDirectory {
		locationPath := filepath.Join(filepath.Join(DATABASE_LIB_PATH, "locations", location))

		baseData, err := tools.ReadParsed(filepath.Join(locationPath, "base.json"))
		if err != nil {
			return fmt.Errorf("error reading base.json for location %s: %w", location, err)
		}

		base, ok := baseData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in base.json for location %s", location)
		}

		baseDataMap, ok := base["data"].(map[string]interface{})
		if ok {
			base = baseDataMap
		}

		dynamicAvailableSpawnsData, err := tools.ReadParsed(filepath.Join(locationPath, "availableSpawns.json"))
		if err != nil {
			return fmt.Errorf("error reading availableSpawns.json for location %s: %w", location, err)
		}
		dynamicAvailableSpawns, ok := dynamicAvailableSpawnsData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in availableSpawns.json for location %s", location)
		}

		locationStruct := LocationStruct{
			Base:                   base,
			DynamicAvailableSpawns: dynamicAvailableSpawns,
			LootSpawns:             setLootSpawns(locationPath),
			Presets:                setLocationPresets(locationPath),
		}

		locations.Locations[location] = locationStruct
	}

	if err := setLocationsLootGen(locations); err != nil {
		return err
	}

	return nil
}

func setLootSpawns(path string) map[string]interface{} {
	lootSpawnsPath := filepath.Join(path, "lootSpawns")
	lootSpawns, err := tools.GetFilesFrom(lootSpawnsPath)
	if err != nil {
		log.Panicf("error reading lootSpawns directory: %v", err)
		return nil
	}
	lootSpawnsMap := make(map[string]interface{}, len(lootSpawns))

	for _, lootSpawn := range lootSpawns {
		lootSpawnName := lootSpawn[:len(lootSpawn)-5]
		lootSpawnPath := filepath.Join(lootSpawnsPath, lootSpawn)

		lootSpawnData, err := tools.ReadParsed(lootSpawnPath)
		if err != nil {
			log.Panicf("error reading %s: %v", lootSpawn, err)
			return nil
		}
		lootSpawnArray, ok := lootSpawnData.([]interface{})
		if !ok {
			log.Panicf("invalid data structure in %s", lootSpawn)
			return nil
		}

		lootSpawnMap := tools.TransformInterfaceIntoMappedArray(lootSpawnArray)
		lootSpawnsMap[lootSpawnName] = lootSpawnMap
	}
	return lootSpawnsMap
}

func setLocationPresets(path string) map[string]map[string]interface{} {
	presetsPath := filepath.Join(path, "#presets")
	presetFiles, err := tools.GetFilesFrom(presetsPath)
	if err != nil {
		log.Panicf("error reading presets directory: %v", err)
		return nil
	}

	presetCount := len(presetFiles)
	presetNames := make([]string, presetCount)
	for i, file := range presetFiles {
		presetNames[i] = file[:len(file)-5]
	}

	presetsMap := make(map[string]map[string]interface{}, presetCount)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(presetCount)

	for i := 0; i < presetCount; i++ {
		go func(i int) {
			defer wg.Done()

			filePath := filepath.Join(presetsPath, presetFiles[i])
			presetData, err := tools.ReadParsed(filePath)
			if err != nil {
				log.Panicf("error reading %s: %v", presetFiles[i], err)
				return
			}

			preset, ok := presetData.(map[string]interface{})
			if !ok {
				log.Panicf("invalid data structure in %s", presetFiles[i])
				return
			}

			mu.Lock()
			presetsMap[presetNames[i]] = preset
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	return presetsMap
}

type LootGenStruct struct {
	Containers map[string]interface{}
	Static     map[string]interface{}
}

func setLocationsLootGen(locations *LocationsStruct) error {
	lootGenPath := filepath.Join(DATABASE_LIB_PATH, "lootGen")
	lootGen := locations.LootGen

	containersData, err := tools.ReadParsed(filepath.Join(lootGenPath, "containersSpawnData.json"))
	if err != nil {
		return fmt.Errorf("error reading containersSpawnData.json: %w", err)
	}
	containers, ok := containersData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in containersSpawnData.json")
	}
	lootGen.Containers = containers

	staticData, err := tools.ReadParsed(filepath.Join(lootGenPath, "staticWeaponsData.json"))
	if err != nil {
		return fmt.Errorf("error reading staticWeaponsData.json: %w", err)
	}
	static, ok := staticData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in staticWeaponsData.json")
	}
	lootGen.Static = static

	return nil
}
*/
