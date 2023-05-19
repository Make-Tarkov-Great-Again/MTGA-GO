package main

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

type DatabaseStruct struct {
	core          CoreStruct
	connections   ConnectionStruct
	items         map[string]interface{}
	locales       LocaleStruct
	templates     TemplatesStruct
	traders       map[string]interface{}
	flea          FleaStruct
	quests        map[string]interface{}
	hideout       HideoutStruct
	locations     LocationsStruct
	weather       map[string]interface{}
	customization map[string]interface{}
	editions      map[string]interface{}
	bot           BotStruct
	profiles      map[string]interface{}
	//bundles  []map[string]interface{}
}

type CoreStruct struct {
	botTemplate    map[string]interface{}
	clientSettings map[string]interface{}
	serverConfig   map[string]interface{}
	globals        map[string]interface{}
	locations      map[string]interface{}
	presets        map[string]interface{}
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	matchMetrics map[string]interface{}
}

type ConnectionStruct struct {
	webSocket      map[string]interface{}
	webSocketPings map[string]interface{}
}

type FleaStruct struct {
	offers           []map[string]interface{}
	offerscount      int
	selectedCategory string
	categories       map[string]interface{}
}

type HideoutStruct struct {
	areas       []map[string]interface{}
	productions []map[string]interface{}
	scavcase    []map[string]interface{}
	qte         []map[string]interface{}
	settings    map[string]interface{}
}

type LocationsStruct struct {
	locations map[string]interface{}
	lootGen   LootGenStruct
}

var Database = DatabaseStruct{}

func initializeDatabase() error {
	Database.core = CoreStruct{
		botTemplate:    map[string]interface{}{},
		clientSettings: map[string]interface{}{},
		serverConfig:   map[string]interface{}{},
		globals:        map[string]interface{}{},
		locations:      map[string]interface{}{},
		matchMetrics:   map[string]interface{}{},
		presets:        map[string]interface{}{},
	}
	Database.connections = ConnectionStruct{
		webSocket:      map[string]interface{}{},
		webSocketPings: map[string]interface{}{},
	}
	Database.items = map[string]interface{}{}
	Database.locales = LocaleStruct{
		locales:   map[string]interface{}{},
		extras:    map[string]interface{}{},
		languages: map[string]interface{}{},
	}
	Database.templates = TemplatesStruct{
		Handbook: HandbookStruct{
			Items:      []map[string]interface{}{},
			Categories: []map[string]interface{}{},
		},
		Prices: map[string]interface{}{},
		TplLookup: TplLookupStruct{
			Items: ItemsLookupStruct{
				byId:     map[string]interface{}{},
				byParent: map[string]interface{}{},
			},
			Categories: CategoriesLookupStruct{
				byId:     map[string]interface{}{},
				byParent: map[string]interface{}{},
			},
		},
	}
	Database.editions = map[string]interface{}{}
	Database.traders = map[string]interface{}{}
	Database.quests = map[string]interface{}{}
	Database.flea = FleaStruct{
		offers:           []map[string]interface{}{},
		offerscount:      0,
		selectedCategory: "",
		categories:       map[string]interface{}{},
	}
	Database.hideout = HideoutStruct{
		areas:       []map[string]interface{}{},
		productions: []map[string]interface{}{},
		scavcase:    []map[string]interface{}{},
		qte:         []map[string]interface{}{},
		settings:    map[string]interface{}{},
	}
	Database.customization = map[string]interface{}{}
	Database.profiles = map[string]interface{}{}
	Database.weather = map[string]interface{}{}
	Database.bot = BotStruct{
		bots:        map[string]interface{}{},
		core:        map[string]interface{}{},
		names:       map[string]interface{}{},
		appearance:  map[string]interface{}{},
		playerScav:  map[string]interface{}{},
		weaponCache: map[string]interface{}{},
	}
	Database.locations = LocationsStruct{
		locations: map[string]interface{}{},
		lootGen: LootGenStruct{
			containers: map[string]interface{}{},
			static:     map[string]interface{}{},
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

	if err := setItems(); err != nil {
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

	if err := setQuests(); err != nil {
		return err
	}

	if err := setHideout(); err != nil {
		return err
	}

	if err := setCustomization(); err != nil {
		return err
	}

	if err := setProfiles(); err != nil {
		return err
	}

	if err := setWeather(); err != nil {
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
	if err := setPresetsCore(core); err != nil {
		return fmt.Errorf("error setting presets: %w", err)
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

const CORE_FILE_PATH string = "database/core"
const (
	BOT_TEMPLATE_FILE_PATH string = CORE_FILE_PATH + "/botTemplate.json"
	CLIENT_SETTINGS_PATH   string = CORE_FILE_PATH + "/client.settings.json"
	GLOBALS_FILE_PATH      string = CORE_FILE_PATH + "/globals.json"
	LOCATIONS_FILE_PATH    string = CORE_FILE_PATH + "/locations.json"
	MATCH_METRICS_PATH     string = CORE_FILE_PATH + "/matchMetrics.json"
	SERVER_CONFIG_PATH     string = CORE_FILE_PATH + "/server.json"
)

func setServerConfigCore(core *CoreStruct) error {
	serverConfig, err := tools.ReadParsed(SERVER_CONFIG_PATH)

	if err != nil {
		return fmt.Errorf("error reading server.json: %w", err)
	}

	serverConfigMap, ok := serverConfig.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in server.json")
	}

	core.serverConfig = serverConfigMap
	return nil
}

func setMatchMetricsCore(core *CoreStruct) error {
	matchMetrics, err := tools.ReadParsed(MATCH_METRICS_PATH)

	if err != nil {
		return fmt.Errorf("error reading matchMetrics.json: %w", err)
	}

	matchMetricsMap, ok := matchMetrics.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in matchMetrics.json")
	}

	core.matchMetrics = matchMetricsMap
	return nil
}

func setGlobalsCore(core *CoreStruct) error {
	globals, err := tools.ReadParsed(GLOBALS_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading globals.json: %w", err)
	}

	globalsMap, ok := globals.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in globals.json")
	}

	globalsData, ok := globalsMap["data"].(map[string]interface{})
	if !ok {
		core.globals = globalsMap
	} else {
		core.globals = globalsData
	}

	return nil
}

func setPresetsCore(core *CoreStruct) error {
	presets, ok := core.globals["ItemPresets"]
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

		if _, ok := core.presets[tpl]; !ok {
			core.presets[tpl] = map[string]interface{}{}
		}

		presetId, ok := preset["_id"].(string)
		if !ok {
			return fmt.Errorf("presetId not found in preset %s", id)
		}

		itemPreset, ok := core.presets[tpl].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid data structure in ItemPresets")
		}

		itemPreset[presetId] = preset

	}
	return nil
}

func setClientSettingsCore(core *CoreStruct) error {
	clientSettings, err := tools.ReadParsed(CLIENT_SETTINGS_PATH)

	if err != nil {
		return fmt.Errorf("error reading client.settings.json: %w", err)
	}

	clientSettingsMap, ok := clientSettings.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in client.settings.json")
	}

	clientSettingsData, ok := clientSettingsMap["data"].(map[string]interface{})
	if !ok {
		core.clientSettings = clientSettingsMap
		return nil
	}

	core.clientSettings = clientSettingsData
	return nil
}

func setLocationsCore(core *CoreStruct) error {
	locations, err := tools.ReadParsed(LOCATIONS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading locations.json: %w", err)
	}

	locationsMap, ok := locations.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in locations.json")
	}

	locationsData, ok := locationsMap["data"].(map[string]interface{})
	if !ok {
		core.locations = locationsMap
	} else {
		core.locations = locationsData
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

	core.botTemplate = botTemplateMap
	return nil
}

const EDITIONS_FILE_PATH string = "database/editions"

type EditionStruct struct {
	bear    map[string]interface{}
	usec    map[string]interface{}
	storage map[string]interface{}
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
			bear:    bearMap,
			usec:    usecMap,
			storage: storageMap,
		}

		// Add the editionData to the Database.editions map
		Database.editions[edition] = editionData
	}
	return nil
}

const LOCALES_FILE_PATH string = "database/locales"

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
	localesDirectory, err := tools.GetDirectoriesFrom(LOCALES_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading locales directory: %w", err)
	}
	locales := &Database.locales

	for _, locale := range localesDirectory {
		localePath := filepath.Join(LOCALES_FILE_PATH, locale, "locale.json")
		menuPath := filepath.Join(LOCALES_FILE_PATH, locale, "menu.json")

		// Add the localeData to the locales map
		locales.locales[locale] = LanguageStruct{
			locale: setLocale(localePath),
			menu:   setMenu(menuPath),
		}
	}

	if err := setLocalesLanguages(locales); err != nil {
		return err
	}

	if err := setLocalesExtras(locales); err != nil {
		return err
	}
	return nil
}

func setLocale(languagePath string) map[string]interface{} {
	localeData, err := tools.ReadParsed(languagePath)
	if err != nil {
		log.Panicf("error reading locale.json for locale: %v", err)
		return nil
	}

	localeMap, ok := localeData.(map[string]interface{})
	if !ok {
		log.Panicf("invalid data structure in locale.json for locale")
		return nil
	}

	localeDataMap, ok := localeMap["data"].(map[string]interface{})
	if !ok {
		return localeMap
	} else {
		return localeDataMap
	}
}

func setMenu(menuPath string) map[string]interface{} {
	menuData, err := tools.ReadParsed(menuPath)
	if err != nil {
		log.Panicf("error reading menu.json for locale: %v", err)
		return nil
	}

	menuMap, ok := menuData.(map[string]interface{})
	if !ok {
		log.Panicf("invalid data structure in menu.json for locale")
		return nil
	}
	menuDataMap, ok := menuMap["data"].(map[string]interface{})
	if !ok {
		return menuMap
	} else {
		return menuDataMap
	}
}

func setLocalesExtras(locales *LocaleStruct) error {
	extrasFilePath := filepath.Join(LOCALES_FILE_PATH, "extras.json")
	if !tools.FileExist(extrasFilePath) {
		return fmt.Errorf("error reading extras.json")
	}
	extras, err := tools.ReadParsed(extrasFilePath)
	if err != nil {
		return fmt.Errorf("error reading extras.json: %w", err)
	}

	extrasMap, ok := extras.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in extras.json")
	}

	locales.extras = extrasMap
	return nil
}

func setLocalesLanguages(locales *LocaleStruct) error {
	languagesFilePath := filepath.Join(LOCALES_FILE_PATH, "languages.json")
	languages, err := tools.ReadParsed(languagesFilePath)
	if err != nil {
		return fmt.Errorf("error reading languages.json: %w", err)
	}

	languagesMap, ok := languages.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in languages.json")
	}

	languagesData, ok := languagesMap["data"].(map[string]interface{})
	if !ok {
		locales.languages = languagesMap
	} else {
		locales.languages = languagesData
	}
	return nil
}

func setItems() error {
	items, err := tools.ReadParsed("database/items.json")

	if err != nil {
		return fmt.Errorf("error reading items.json: %w", err)
	}

	Database.items = items.(map[string]interface{})
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

	handbook, ok := templatesData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in templates.json")
	}

	handbookData, ok := handbook["data"].(map[string]interface{})
	if ok {
		handbook = handbookData
	}

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

		byCategory.byId[Id] = ParentId
		if hasParent {
			if mapSlice, ok := byCategory.byParent[ParentId].([]string); ok {
				byCategory.byParent[ParentId] = append(mapSlice, Id)
			} else {
				byCategory.byParent[ParentId] = []string{Id}
			}
		}
	}
	return nil
}

func setHandbookItems(items []map[string]interface{}, templates *TemplatesStruct) error {
	templates.Handbook.Items = tools.AuditArrayCapacity(items)

	pricesData, err := tools.ReadParsed("database/liveflea.json")
	if err != nil {
		return fmt.Errorf("error reading liveflea.json: %w", err)
	}
	prices := pricesData.(map[string]interface{}) // prices is a map[string]interface{}

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
		byItem.byId[Id] = price
		templates.Prices[Id] = price

		// add the item to the byParent map
		if mapSlice, ok := byItem.byParent[ParentId].([]string); ok {
			// if the key exists, append the value to the slice
			byItem.byParent[ParentId] = append(mapSlice, Id)
		} else {
			byItem.byParent[ParentId] = []string{Id}
		}
	}

	return nil
}

const TRADERS_FILE_PATH string = "database/traders"

type TraderStruct struct {
	assort      AssortStruct
	base        map[string]interface{}
	baseAssort  map[string]interface{}
	dialogue    map[string]interface{}
	questassort map[string]interface{}
	suits       map[string]interface{}
}

func setTraders() error {
	tradersDirectory, err := tools.GetDirectoriesFrom(TRADERS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading traders directory: %w", err)
	}

	trader := TraderStruct{}
	for _, traderID := range tradersDirectory {
		trader.assort = AssortStruct{
			items:             []map[string]interface{}{},
			barter_scheme:     map[string]interface{}{},
			loyal_level_items: map[string]interface{}{},
		}

		trader.base = map[string]interface{}{}
		trader.baseAssort = map[string]interface{}{}
		trader.dialogue = map[string]interface{}{}
		trader.questassort = map[string]interface{}{}
		trader.suits = map[string]interface{}{}

		traderPath := filepath.Join(TRADERS_FILE_PATH, traderID)

		if assort, err := setTraderAssort(filepath.Join(traderPath)); err != nil {
			return err
		} else {
			trader.assort = assort
		}

		if base, err := setTraderBase(traderPath); err != nil {
			return fmt.Errorf("error reading base.json for trader %s: %w", traderID, err)
		} else {
			trader.base = base
		}

		if dialogue, err := setTraderDialogue(traderPath); err != nil {
			trader.dialogue = map[string]interface{}{}
		} else {
			trader.dialogue = dialogue
		}

		if questassort, err := setTraderQuestAssort(traderPath); err != nil {
			trader.questassort = map[string]interface{}{}
		} else {
			trader.questassort = questassort
		}

		if suits, err := setTraderSuits(traderPath); err != nil {
			trader.suits = map[string]interface{}{}
		} else {
			trader.suits = suits
		}

		Database.traders[traderID] = trader
	}

	return nil
}

type AssortStruct struct {
	items             []map[string]interface{}
	barter_scheme     map[string]interface{}
	loyal_level_items map[string]interface{}
}

func setTraderBase(path string) (map[string]interface{}, error) {
	data, err := tools.ReadParsed(filepath.Join(path, "base.json"))
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %w", path, err)
	}

	base, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data structure in base @ %s", path)
	}

	if baseData, ok := base["data"].(map[string]interface{}); ok {
		base = baseData
	}

	return base, nil
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

func setQuests() error {
	quests, err := tools.ReadParsed("database/quests.json")

	if err != nil {
		return fmt.Errorf("error reading quests.json: %w", err)
	}

	questsMap, ok := quests.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in quests.json")
	}

	questsData, ok := questsMap["data"].(map[string]interface{})
	if !ok {
		Database.quests = questsMap
	} else {
		Database.quests = questsData
	}

	return nil
}

const HIDEOUT_FILE_PATH string = "database/hideout"
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

	Database.hideout = HideoutStruct{
		areas:       areas,
		productions: productions,
		scavcase:    scavcase,
		qte:         qte,
		settings:    settings,
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

func setCustomization() error {
	data, err := tools.ReadParsed("database/customization.json")
	if err != nil {
		return fmt.Errorf("error reading customization.json: %w", err)
	}

	customization, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("customization.json has invalid structure")
	}

	customizationData, ok := customization["data"].(map[string]interface{})
	if ok {
		customization = customizationData
	}

	Database.customization = customization
	return nil
}

const USER_FILE_PATH string = "user"
const PROFILES_FILE_PATH string = USER_FILE_PATH + "/profiles"

type ProfileStruct struct {
	account   map[string]interface{}
	character map[string]interface{}
	storage   map[string]interface{}
	dialogues map[string]interface{}
	raid      RaidProfileStruct
}
type RaidProfileStruct struct {
	lastLocation RaidLocationStruct
	carExtracts  int
}
type RaidLocationStruct struct {
	name      string
	insurance bool
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

	for _, profileID := range profilesDirectory {
		profilePath := filepath.Join(PROFILES_FILE_PATH, profileID)
		profile := ProfileStruct{
			account:   setAccount(profilePath, profileID),
			character: setCharacter(profilePath, profileID),
			storage:   setStorage(profilePath, profileID),
			dialogues: setDialogues(profilePath, profileID),
			raid: RaidProfileStruct{
				lastLocation: RaidLocationStruct{
					name:      "",
					insurance: false,
				},
				carExtracts: 0,
			},
		}
		Database.profiles[profileID] = profile
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

func setWeather() error {
	data, err := tools.ReadParsed("database/weather.json")
	if err != nil {
		return fmt.Errorf("error reading weather.json: %w", err)
	}

	weather, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in weather.json")
	}

	weatherData, ok := weather["data"].(map[string]interface{})
	if ok {
		weather = weatherData
	}

	Database.weather = weather
	return nil
}

const BOT_FILE_PATH string = "database/bot"

type BotStruct struct {
	bots        map[string]interface{}
	core        map[string]interface{}
	names       map[string]interface{}
	appearance  map[string]interface{}
	playerScav  map[string]interface{}
	weaponCache map[string]interface{}
}

func setBot() error {
	bot := &Database.bot

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

	bot.core = botGlobals
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

	bot.names = names
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

	bot.appearance = appearance
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

	bot.playerScav = playerScav
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

	bot.weaponCache = weaponCache
	return nil
}

const BOTS_FILE_PATH string = "database/bot/bots"

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

	for _, aiType := range botsFiles {
		botTypePath := filepath.Join(BOTS_FILE_PATH, aiType)

		bot.bots[aiType] = BotTypeStruct{
			health:     setBotTypeHealth(botTypePath),
			loadout:    setBotTypeLoadout(botTypePath),
			difficulty: setBotTypeDifficulty(botTypePath),
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
		difficultyName := strings.TrimSuffix(difficulty, ".json")
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
	base                   map[string]interface{}
	dynamicAvailableSpawns map[string]interface{}
	lootSpawns             map[string]interface{}
	//waves                  []map[string]interface{}
	//bossWaves              []map[string]interface{}
	presets map[string]interface{}
}

func setLocations() error {
	locationsDirectory, err := tools.GetDirectoriesFrom("database/locations")
	if err != nil {
		return fmt.Errorf("error reading locations directory: %w", err)
	}

	locations := &Database.locations

	for _, location := range locationsDirectory {
		locationPath := filepath.Join("database/locations", location)

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
			base:                   base,
			dynamicAvailableSpawns: dynamicAvailableSpawns,
			lootSpawns:             setLootSpawns(locationPath),
			presets:                setLocationPresets(locationPath),
		}

		locations.locations[location] = locationStruct
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
		lootSpawnName := strings.TrimSuffix(lootSpawn, ".json")
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

func setLocationPresets(path string) map[string]interface{} {
	presetsPath := filepath.Join(path, "#presets")
	presets, err := tools.GetFilesFrom(presetsPath)
	if err != nil {
		log.Panicf("error reading presets directory: %v", err)
		return nil
	}

	presetsMap := make(map[string]interface{}, len(presets))
	for _, preset := range presets {
		presetName := strings.TrimSuffix(preset, ".json")
		presetPath := filepath.Join(presetsPath, preset)

		presetData, err := tools.ReadParsed(presetPath)
		if err != nil {
			log.Panicf("error reading %s: %v", presetData, err)
			return nil
		}

		preset, ok := presetData.(map[string]interface{})
		if !ok {
			log.Panicf("invalid data structure in %s", preset)
			return nil
		}

		presetsMap[presetName] = preset
	}

	return presetsMap
}

type LootGenStruct struct {
	containers map[string]interface{}
	static     map[string]interface{}
}

func setLocationsLootGen(locations *LocationsStruct) error {
	lootGenPath := "database/lootGen"
	lootGen := &locations.lootGen

	containersData, err := tools.ReadParsed(filepath.Join(lootGenPath, "containersSpawnData.json"))
	if err != nil {
		return fmt.Errorf("error reading containersSpawnData.json: %w", err)
	}
	containers, ok := containersData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in containersSpawnData.json")
	}
	lootGen.containers = containers

	staticData, err := tools.ReadParsed(filepath.Join(lootGenPath, "staticWeaponsData.json"))
	if err != nil {
		return fmt.Errorf("error reading staticWeaponsData.json: %w", err)
	}
	static, ok := staticData.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data structure in staticWeaponsData.json")
	}
	lootGen.static = static

	return nil
}
