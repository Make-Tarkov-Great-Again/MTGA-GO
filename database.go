package main

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"
)

type DatabaseStruct struct {
	core      CoreStruct
	items     map[string]interface{}
	locales   LocaleStruct
	templates TemplatesStruct
	//traders       map[string]interface{}
	//flea          map[string]interface{}
	//quests        map[string]interface{}
	//hideout       map[string]interface{}
	//locations     map[string]interface{}
	//weather       map[string]interface{}
	//customization map[string]interface{}
	editions map[string]interface{}
	//presets       map[string]interface{}
	//bot           map[string]interface{}
	//profiles      map[string]interface{}
	//bundles       []interface{}
}

type CoreStruct struct {
	botTemplate    map[string]interface{}
	clientSettings map[string]interface{}
	serverConfig   map[string]interface{}
	globals        map[string]interface{}
	locations      map[string]interface{}
	//gameplay        map[string]interface{}
	//hideoutSettings map[string]interface{}
	//blacklist       []interface{}
	matchMetrics map[string]interface{}
	connections  ConnectionStruct
}

type ConnectionStruct struct {
	webSocket      map[string]interface{}
	webSocketPings map[string]interface{}
}

var Database = DatabaseStruct{}

func initializeDatabase() error {
	Database.core = CoreStruct{
		botTemplate:    make(map[string]interface{}),
		clientSettings: make(map[string]interface{}),
		serverConfig:   make(map[string]interface{}),
		globals:        make(map[string]interface{}),
		locations:      make(map[string]interface{}),
		matchMetrics:   make(map[string]interface{}),
		connections: ConnectionStruct{
			webSocket:      make(map[string]interface{}),
			webSocketPings: make(map[string]interface{}),
		},
	}
	Database.items = make(map[string]interface{})
	Database.locales = LocaleStruct{
		locales:   make(map[string]interface{}),
		extras:    make(map[string]interface{}),
		languages: make(map[string]interface{}),
	}
	Database.templates = TemplatesStruct{
		Handbook: HandbookStruct{
			Items:      []interface{}{},
			Categories: []interface{}{},
		},
		Prices: make(map[string]interface{}),
		TplLookup: TplLookupStruct{
			Items: ItemsLookupStruct{
				byId:     make(map[string]interface{}),
				byParent: make(map[string]interface{}),
			},
			Categories: CategoriesLookupStruct{
				byId:     make(map[string]interface{}),
				byParent: make(map[string]interface{}),
			},
		},
	}
	Database.editions = make(map[string]interface{})

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

	if err := setItems(); err != nil {
		return err
	}

	if err := setLocales(); err != nil {
		return err
	}

	if err := setTemplates(); err != nil {
		return err
	}

	return nil
}

func setDatabaseCore() error {
	core := &Database.core

	if err := setServerConfigCore(core); err != nil {
		return fmt.Errorf("error setting server config: %w", err)
	}
	if err := setMatchMetricsCore(core); err != nil {
		return fmt.Errorf("error setting match metrics: %w", err)
	}
	if err := setGlobalsCore(core); err != nil {
		return fmt.Errorf("error setting globals: %w", err)
	}
	if err := setClientSettingsCore(core); err != nil {
		return fmt.Errorf("error setting client settings: %w", err)
	}
	if err := setLocationsCore(core); err != nil {
		return fmt.Errorf("error setting locations: %w", err)
	}
	if err := setBotTemplateCore(core); err != nil {
		return fmt.Errorf("error setting bot template: %w", err)
	}

	return nil
}

// checkAndReturnIfDataPropertyExists checks if the data property exists and returns it if it does.
func checkAndReturnIfDataPropertyExists(data interface{}) map[string]interface{} {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}
	dataValue, ok := dataMap["data"]
	if !ok {
		return dataMap
	}
	dataType, ok := dataValue.(map[string]interface{})
	if !ok {
		return dataMap
	}
	return dataType
}

const CORE_FILE_PATH = "database/core"
const (
	BOT_TEMPLATE_FILE_PATH = CORE_FILE_PATH + "/botTemplate.json"
	CLIENT_SETTINGS_PATH   = CORE_FILE_PATH + "/client.settings.json"
	GLOBALS_FILE_PATH      = CORE_FILE_PATH + "/globals.json"
	LOCATIONS_FILE_PATH    = CORE_FILE_PATH + "/locations.json"
	MATCH_METRICS_PATH     = CORE_FILE_PATH + "/matchMetrics.json"
	SERVER_CONFIG_PATH     = CORE_FILE_PATH + "/server.json"
)

func setServerConfigCore(core *CoreStruct) error {
	serverConfig, err := tools.ReadParsed(SERVER_CONFIG_PATH)

	if err != nil {
		return fmt.Errorf("error reading server.json: %w", err)
	}

	core.serverConfig = checkAndReturnIfDataPropertyExists(serverConfig)
	return nil
}

func setMatchMetricsCore(core *CoreStruct) error {
	matchMetrics, err := tools.ReadParsed(MATCH_METRICS_PATH)

	if err != nil {
		return fmt.Errorf("error reading matchMetrics.json: %w", err)
	}

	core.matchMetrics = checkAndReturnIfDataPropertyExists(matchMetrics)
	return nil
}

func setGlobalsCore(core *CoreStruct) error {
	globals, err := tools.ReadParsed(GLOBALS_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading globals.json: %w", err)
	}

	core.globals = checkAndReturnIfDataPropertyExists(globals)
	return nil
}

func setClientSettingsCore(core *CoreStruct) error {
	clientSettings, err := tools.ReadParsed(CLIENT_SETTINGS_PATH)

	if err != nil {
		return fmt.Errorf("error reading client.settings.json: %w", err)
	}
	core.clientSettings = checkAndReturnIfDataPropertyExists(clientSettings)
	return nil
}

func setLocationsCore(core *CoreStruct) error {
	locations, err := tools.ReadParsed(LOCATIONS_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading locations.json: %w", err)
	}

	core.locations = checkAndReturnIfDataPropertyExists(locations)
	return nil
}

func setBotTemplateCore(core *CoreStruct) error {
	botTemplate, err := tools.ReadParsed(BOT_TEMPLATE_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading botTemplate.json: %w", err)
	}

	core.botTemplate = checkAndReturnIfDataPropertyExists(botTemplate)
	return nil
}

const EDITIONS_FILE_PATH = "database/editions"

type EditionStruct struct {
	bear    map[string]interface{}
	usec    map[string]interface{}
	storage map[string]interface{}
}

func setEditions() error {
	editionsDirectory := tools.GetDirectoriesFrom(EDITIONS_FILE_PATH)

	for _, edition := range editionsDirectory {
		editionPath := filepath.Join(EDITIONS_FILE_PATH, edition)

		// Read character_bear.json and store in bear map
		bearData, err := tools.ReadParsed(filepath.Join(editionPath, "character_bear.json"))
		if err != nil {
			return fmt.Errorf("error reading character_bear.json for edition %s: %w", edition, err)
		}

		// Read character_usec.json and store in usec map
		usecData, err := tools.ReadParsed(filepath.Join(editionPath, "character_usec.json"))
		if err != nil {
			return fmt.Errorf("error reading character_usec.json for edition %s: %w", edition, err)
		}

		// Read storage.json and store in storage map
		storageData, err := tools.ReadParsed(filepath.Join(editionPath, "storage.json"))
		if err != nil {
			return fmt.Errorf("error reading storage.json for edition %s: %w", edition, err)
		}

		// Create an EditionStruct for this edition and populate the maps
		editionData := EditionStruct{
			bear:    bearData.(map[string]interface{}),
			usec:    usecData.(map[string]interface{}),
			storage: storageData.(map[string]interface{}),
		}

		// Add the editionData to the Database.editions map
		Database.editions[edition] = editionData
	}
	return nil
}

const LOCALES_FILE_PATH = "database/locales"

type LocaleStruct struct {
	locales   map[string]interface{}
	extras    map[string]interface{}
	languages map[string]interface{}
}
type LanguageStruct struct {
	locale map[string]interface{}
	menu   map[string]interface{}
}

func setLocales() error {
	localesDirectory := tools.GetDirectoriesFrom(LOCALES_FILE_PATH)
	locales := &Database.locales

	for _, locale := range localesDirectory {
		localePath := filepath.Join(LOCALES_FILE_PATH, locale)
		language := filepath.Join(localePath, "locale.json")
		menu := filepath.Join(localePath, "menu.json")

		// Read language.json and store in language map
		localeData, err := tools.ReadParsed(language)
		if err != nil {
			return fmt.Errorf("error reading locale.json for locale %s: %w", locale, err)
		}

		// Read menu.json and store in menu map
		menuData, err := tools.ReadParsed(menu)
		if err != nil {
			return fmt.Errorf("error reading menu.json for locale %s: %w", locale, err)
		}

		// Create a LanguageStruct for this locale and populate the maps
		languageData := LanguageStruct{
			locale: localeData.(map[string]interface{}),
			menu:   menuData.(map[string]interface{}),
		}

		// Add the localeData to the locales map
		locales.locales[locale] = languageData
	}

	languagesFilePath := filepath.Join(LOCALES_FILE_PATH, "languages.json")
	languages, err := tools.ReadParsed(languagesFilePath)
	if err != nil {
		return fmt.Errorf("error reading languages.json: %w", err)
	}
	locales.languages = checkAndReturnIfDataPropertyExists(languages)

	extrasFilePath := filepath.Join(LOCALES_FILE_PATH, "extras.json")
	if !tools.FileExist(extrasFilePath) {
		return fmt.Errorf("error reading extras.json")
	}
	extras, err := tools.ReadParsed(extrasFilePath)
	if err != nil {
		return fmt.Errorf("error reading extras.json: %w", err)
	}
	locales.extras = checkAndReturnIfDataPropertyExists(extras)
	return nil
}

func setItems() error {
	items, err := tools.ReadParsed("database/items.json")

	if err != nil {
		return fmt.Errorf("error reading items.json: %w", err)
	}

	Database.items = checkAndReturnIfDataPropertyExists(items)
	return nil
}

type TemplatesStruct struct {
	Handbook  HandbookStruct
	Prices    map[string]interface{}
	TplLookup TplLookupStruct
}

type HandbookStruct struct {
	Items      []interface{}
	Categories []interface{}
}

type TplLookupStruct struct {
	Items      ItemsLookupStruct
	Categories CategoriesLookupStruct
}

type ItemsLookupStruct struct {
	byId     map[string]interface{}
	byParent map[string]interface{}
}

type CategoriesLookupStruct struct {
	byId     map[string]interface{}
	byParent map[string]interface{}
}

func setTemplates() error {
	templates := &Database.templates

	templatesData, err := tools.ReadParsed("database/templates.json")
	if err != nil {
		return fmt.Errorf("error reading templates.json: %w", err)
	}

	handbookData := checkAndReturnIfDataPropertyExists(templatesData)
	setHandbookItems(handbookData, templates)
	setHandbookCategories(handbookData, templates)

	return nil
}

func setHandbookCategories(handbookData map[string]interface{}, templates *TemplatesStruct) error {
	categories := handbookData["Categories"].([]interface{})
	categoriesSize := len(categories)

	templates.Handbook.Categories = make([]interface{}, 0, categoriesSize) // Create a new slice with capacity of the old slice length
	for i := 0; i < categoriesSize; i++ {
		templates.Handbook.Categories = append(templates.Handbook.Categories, categories[i]) // Append the value from old slice to the new slice
	}

	byCategory := &templates.TplLookup.Categories // pointer to the CategoriesLookupStruct
	for _, value := range templates.Handbook.Categories {

		category := value.(map[string]interface{})
		Id := category["Id"].(string)
		ParentId, hasParent := category["ParentId"].(string)

		byCategory.byId[Id] = ParentId
		if hasParent {
			if _, ok := byCategory.byParent[ParentId]; ok {
				byCategory.byParent[ParentId] = append(byCategory.byParent[ParentId].([]interface{}), Id)
			} else {
				byCategory.byParent[ParentId] = []interface{}{Id}
			}
		}
	}
	return nil
}

func setHandbookItems(handbookData map[string]interface{}, templates *TemplatesStruct) error {
	items := handbookData["Items"].([]interface{})
	itemsSize := len(items)

	templates.Handbook.Items = make([]interface{}, 0, itemsSize) // Create a new slice with capacity of the old slice length
	for i := 0; i < itemsSize; i++ {
		templates.Handbook.Items = append(templates.Handbook.Items, items[i]) // Append the value from old slice to the new slice
	}

	pricesData, err := tools.ReadParsed("database/liveflea.json")
	if err != nil {
		return fmt.Errorf("error reading liveflea.json: %w", err)
	}
	prices := pricesData.(map[string]interface{}) // prices is a map[string]interface{}

	byItem := &templates.TplLookup.Items // pointer to the ItemsLookupStruct
	for _, item := range templates.Handbook.Items {
		itemMap := item.(map[string]interface{})
		Id := itemMap["Id"].(string)
		ParentId := itemMap["ParentId"].(string)

		// set the item price to the price from liveflea.json if it exists, otherwise use the item from items.json
		var price interface{} // store price in this variable
		if p, ok := prices[Id]; ok {
			price = p
		} else {
			price = item
		}
		byItem.byId[Id] = price
		templates.Prices[Id] = price

		// add the item to the byParent map
		if _, ok := byItem.byParent[ParentId]; ok {
			// if the key exists, append the value to the slice
			byItem.byParent[ParentId] = append(byItem.byParent[ParentId].([]interface{}), Id)
		} else {
			byItem.byParent[ParentId] = []interface{}{Id}
		}
	}

	return nil
}
