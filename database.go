package main

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

type DatabaseStruct struct {
	core        CoreStruct
	connections ConnectionStruct
	items       map[string]interface{}
	locales     LocaleStruct
	templates   TemplatesStruct
	traders     map[string]interface{}
	flea        FleaStruct
	quests      map[string]interface{}
	hideout     HideoutStruct
	//locations map[string]interface{}
	weather       map[string]interface{}
	customization map[string]interface{}
	editions      map[string]interface{}
	//presets       map[string]interface{}
	bot      BotStruct
	profiles map[string]interface{}
	//bundles  []map[string]interface{}
}

type CoreStruct struct {
	botTemplate    map[string]interface{}
	clientSettings map[string]interface{}
	serverConfig   map[string]interface{}
	globals        map[string]interface{}
	locations      map[string]interface{}
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
	productions map[string]interface{}
	scavcase    []map[string]interface{}
	qte         []map[string]interface{}
	settings    map[string]interface{}
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
	}
	Database.connections = ConnectionStruct{
		webSocket:      make(map[string]interface{}),
		webSocketPings: make(map[string]interface{}),
	}
	Database.items = make(map[string]interface{})
	Database.locales = LocaleStruct{
		locales:   make(map[string]interface{}),
		extras:    make(map[string]interface{}),
		languages: make(map[string]interface{}),
	}
	Database.templates = TemplatesStruct{
		Handbook: HandbookStruct{
			Items:      []map[string]interface{}{},
			Categories: []map[string]interface{}{},
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
	Database.traders = make(map[string]interface{})
	Database.quests = make(map[string]interface{})
	Database.flea = FleaStruct{
		offers:           []map[string]interface{}{},
		offerscount:      0,
		selectedCategory: "",
		categories:       make(map[string]interface{}),
	}
	Database.hideout = HideoutStruct{
		areas:       []map[string]interface{}{},
		productions: make(map[string]interface{}),
		scavcase:    []map[string]interface{}{},
		qte:         []map[string]interface{}{},
		settings:    make(map[string]interface{}),
	}
	Database.customization = make(map[string]interface{})
	Database.profiles = map[string]interface{}{}
	Database.weather = make(map[string]interface{})
	Database.bot = BotStruct{
		bots:        make(map[string]interface{}),
		core:        make(map[string]interface{}),
		names:       make(map[string]interface{}),
		appearance:  make(map[string]interface{}),
		playerScav:  make(map[string]interface{}),
		weaponCache: make(map[string]interface{}),
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

	core.serverConfig = serverConfig.(map[string]interface{})
	return nil
}

func setMatchMetricsCore(core *CoreStruct) error {
	matchMetrics, err := tools.ReadParsed(MATCH_METRICS_PATH)

	if err != nil {
		return fmt.Errorf("error reading matchMetrics.json: %w", err)
	}

	core.matchMetrics = matchMetrics.(map[string]interface{})
	return nil
}

func setGlobalsCore(core *CoreStruct) error {
	globals, err := tools.ReadParsed(GLOBALS_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading globals.json: %w", err)
	}

	core.globals = globals.(map[string]interface{})
	return nil
}

func setClientSettingsCore(core *CoreStruct) error {
	clientSettings, err := tools.ReadParsed(CLIENT_SETTINGS_PATH)

	if err != nil {
		return fmt.Errorf("error reading client.settings.json: %w", err)
	}
	core.clientSettings = clientSettings.(map[string]interface{})["data"].(map[string]interface{})
	return nil
}

func setLocationsCore(core *CoreStruct) error {
	locations, err := tools.ReadParsed(LOCATIONS_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading locations.json: %w", err)
	}

	core.locations = locations.(map[string]interface{})["data"].(map[string]interface{})
	return nil
}

func setBotTemplateCore(core *CoreStruct) error {
	botTemplate, err := tools.ReadParsed(BOT_TEMPLATE_FILE_PATH)

	if err != nil {
		return fmt.Errorf("error reading botTemplate.json: %w", err)
	}

	core.botTemplate = botTemplate.(map[string]interface{})
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
	locales.languages = languages.(map[string]interface{})["data"].(map[string]interface{})

	extrasFilePath := filepath.Join(LOCALES_FILE_PATH, "extras.json")
	if !tools.FileExist(extrasFilePath) {
		return fmt.Errorf("error reading extras.json")
	}
	extras, err := tools.ReadParsed(extrasFilePath)
	if err != nil {
		return fmt.Errorf("error reading extras.json: %w", err)
	}
	locales.extras = extras.(map[string]interface{})
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

	handbookData := templatesData.(map[string]interface{})["data"].(map[string]interface{})

	items := tools.TransformInterfaceIntoMappedArray(handbookData["Items"].([]interface{}))
	categories := tools.TransformInterfaceIntoMappedArray(handbookData["Categories"].([]interface{}))

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

type AssortStruct struct {
	items             []map[string]interface{}
	barter_scheme     map[string]interface{}
	loyal_level_items map[string]interface{}
}

func setTraders() error {
	tradersDirectory, err := tools.GetDirectoriesFrom(TRADERS_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading traders directory: %w", err)
	}

	for i := 0; i < len(tradersDirectory); i++ {
		traderID := tradersDirectory[i]
		traderPath := filepath.Join(TRADERS_FILE_PATH, traderID)

		trader := make(map[string]interface{})

		for _, file := range []string{"base.json", "assort.json", "questassort.json", "suits.json", "dialogue.json"} {
			if parsed, err := tools.ReadParsed(filepath.Join(traderPath, file)); err != nil {
				continue
				//return fmt.Errorf("error reading %s for trader %s: %w", file, traderID, err)
			} else {
				fileName := strings.Split(file, ".")[0]

				if file == "assort.json" {
					assort := parsed.(map[string]interface{})
					trader["assort"] = AssortStruct{}
					trader["baseAssort"] = AssortStruct{
						items:             tools.TransformInterfaceIntoMappedArray(assort["items"].([]interface{})),
						barter_scheme:     assort["barter_scheme"].(map[string]interface{}),
						loyal_level_items: assort["loyal_level_items"].(map[string]interface{}),
					}
				} else if file == "suits.json" {
					trader["suits"] = tools.TransformInterfaceIntoMappedArray(parsed.([]interface{}))
				} else {
					trader[fileName] = parsed.(map[string]interface{})
				}
			}
		}

		Database.traders[traderID] = trader
	}

	return nil
}

func setQuests() error {
	quests, err := tools.ReadParsed("database/quests.json")

	if err != nil {
		return fmt.Errorf("error reading quests.json: %w", err)
	}

	Database.quests = quests.(map[string]interface{})
	return nil
}

const HIDEOUT_FILE_PATH string = "database/hideout"

func setHideout() error {
	hideoutFiles, err := tools.GetFilesFrom(HIDEOUT_FILE_PATH)
	if err != nil {
		return fmt.Errorf("error reading hideout directory: %w", err)
	}

	for _, file := range hideoutFiles {
		fileName := strings.Split(file, ".")[0]
		filePath := filepath.Join(HIDEOUT_FILE_PATH, file)

		if data, err := tools.ReadParsed(filePath); err != nil {
			return fmt.Errorf("error reading %s: %w", file, err)
		} else {
			switch fileName {
			case "areas":
				Database.hideout.areas = tools.TransformInterfaceIntoMappedArray(data.([]interface{}))
			case "productions":
				Database.hideout.productions = data.(map[string]interface{})
			case "scavcase":
				Database.hideout.scavcase = tools.TransformInterfaceIntoMappedArray(data.([]interface{}))
			case "qte":
				Database.hideout.qte = tools.TransformInterfaceIntoMappedArray(data.([]interface{}))
			case "settings":
				Database.hideout.settings = data.(map[string]interface{})
			}
		}
	}

	return nil
}

func setCustomization() error {
	customization, err := tools.ReadParsed("database/customization.json")

	if err != nil {
		return fmt.Errorf("error reading customization.json: %w", err)
	}

	Database.customization = customization.(map[string]interface{})["data"].(map[string]interface{})
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
	if !ok || account == nil {
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
	if !ok || character == nil {
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
	if !ok || storage == nil {
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
	if !ok || dialogues == nil {
		log.Printf("Dialogues.json for profile %s is empty", profileID)
		return nil
	}

	return dialogues
}

func setWeather() error {
	weather, err := tools.ReadParsed("database/weather.json")

	if err != nil {
		return fmt.Errorf("error reading weather.json: %w", err)
	}

	Database.weather = weather.(map[string]interface{})
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
	core, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "__BotGlobalSettings.json"))

	if err != nil {
		return fmt.Errorf("error reading core.json: %w", err)
	}

	bot.core = core.(map[string]interface{})
	return nil
}

func setBotNames(bot *BotStruct) error {
	names, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "names.json"))

	if err != nil {
		return fmt.Errorf("error reading names.json: %w", err)
	}
	bot.names = names.(map[string]interface{})
	return nil
}

func setBotAppearance(bot *BotStruct) error {
	appearance, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "appearance.json"))

	if err != nil {
		return fmt.Errorf("error reading appearance.json: %w", err)
	}
	bot.appearance = appearance.(map[string]interface{})
	return nil
}

func setBotPlayerScav(bot *BotStruct) error {
	playerScav, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "playerScav.json"))

	if err != nil {
		return fmt.Errorf("error reading playerScav.json: %w", err)
	}
	bot.playerScav = playerScav.(map[string]interface{})
	return nil
}

func setBotWeaponCache(bot *BotStruct) error {
	weaponCache, err := tools.ReadParsed(filepath.Join(BOT_FILE_PATH, "weaponCache.json"))

	if err != nil {
		return fmt.Errorf("error reading weaponCache.json: %w", err)
	}
	bot.weaponCache = weaponCache.(map[string]interface{})
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

	healthData, _ := tools.ReadParsed(healthFilePath)
	if healthData != nil {
		health := healthData.(map[string]interface{})
		if len(health) > 1 {
			healthMap := make(map[string]interface{})
			for key, value := range health {
				healthMap[key] = value.(map[string]interface{})
			}
			return healthMap
		} else {
			return health
		}
	}
	//log.Printf("health.json for %s is empty", path)
	return nil
}

func setBotTypeLoadout(path string) map[string]interface{} {
	loadoutFilePath := filepath.Join(path, "loadout.json")

	loadoutData, _ := tools.ReadParsed(loadoutFilePath)
	if loadoutData == nil {
		//log.Printf("loadout.json for %s is empty", path)
		return nil
	} else {
		return loadoutData.(map[string]interface{})
	}
}

func setBotTypeDifficulty(path string) map[string]interface{} {
	difficultiesPath := filepath.Join(path, "difficulties")

	difficultiesDirectory, err := tools.GetFilesFrom(difficultiesPath)
	if err != nil {
		log.Panicf("error reading difficulties directory: %v", err)
		return nil
	}

	difficulties := make(map[string]interface{})
	for _, difficulty := range difficultiesDirectory {
		difficultyPath := filepath.Join(difficultiesPath, difficulty)

		difficultyData, err := tools.ReadParsed(difficultyPath)
		if err != nil {
			log.Panicf("error reading %s: %v", difficulty, err)
			return nil
		}
		difficultyName := strings.Split(difficulty, ".")[0]
		difficulties[difficultyName] = difficultyData.(map[string]interface{})
	}
	return difficulties
}
