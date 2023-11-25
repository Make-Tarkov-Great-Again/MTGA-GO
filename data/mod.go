package data

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

type ModInfo struct {
	PackageName  string
	PackageAlias string
	MtgaVersion  string `json:"MTGA_Version"`
	Parameters   *advancedModInfo
	Config       map[string]any `json:",omitempty"`
}

type advancedModInfo struct {
	CustomRoutes  bool `json:",omitempty"`
	CustomItems   bool `json:",omitempty"`
	CustomClothes bool `json:",omitempty"`
	CustomTraders bool `json:",omitempty"`
}

type Manifests struct {
	Manifests []*Manifest `json:"manifest"`
}

type Manifest struct {
	ModPath        string   `json:"modPath"`
	Key            string   `json:"key"`
	Path           string   `json:"path"`
	FilePath       string   `json:"filePath,omitempty"`
	DependencyKeys []string `json:"dependencyKeys"`
}

var modBundleDirPaths = make([]string, 0)

var bundleManifests []*Manifest

func (m *ModInfo) GetConfig() map[string]any {
	return m.Config
}

func GetBundleManifests() []*Manifest {
	return bundleManifests
}

func ClearBundleManifests() {
	bundleManifests = make([]*Manifest, 0)
}

func AddModBundleDirPath(modBundleDirPath string) {
	modBundleDirPaths = append(modBundleDirPaths, modBundleDirPath)
}

func LoadBundleManifests() {
	if len(modBundleDirPaths) == 0 {
		return
	}

	startTime := time.Now()
	bundleLoaded := 0
	totalBundles := 0

	isLocal := GetServerConfig().IP == "127.0.0.1"
	var mainAddress string
	if !isLocal {
		mainAddress = GetMainAddress()
	}

	for _, path := range modBundleDirPaths {
		bundlesSubDirectories, err := tools.GetDirectoriesFrom(path)
		if err != nil {
			log.Println(err)
			return
		}

		for subDir := range bundlesSubDirectories {
			bundleMainDirPath := filepath.Join(path, subDir)
			bundlesJSONPath := filepath.Join(bundleMainDirPath, "bundles.json")
			if !tools.FileExist(bundlesJSONPath) {
				err = fmt.Errorf("bundles.json file not located in %s, returning", path)
				log.Println(err)
				return
			}

			manifests := new(Manifests)
			data := tools.GetJSONRawMessage(bundlesJSONPath)
			if err := json.Unmarshal(data, &manifests); err != nil {
				log.Println(err)
				return
			}

			totalBundles += len(manifests.Manifests)
			for _, manifest := range manifests.Manifests {
				bundlesFolder := filepath.Join(bundleMainDirPath, "bundles")

				split := strings.Split(manifest.Key, "/")
				var name string
				if len(split) == 1 {
					name = split[0]
				} else {
					name = strings.Join(split[2:], "/")
				}

				bundlePath := filepath.Join(bundlesFolder, name)
				if !tools.FileExist(bundlePath) {
					err := fmt.Sprintf("bundle %s does not exist in %s", manifest.Key, bundlesFolder)
					modCritiqueLog[subDir] = append(modCritiqueLog[subDir], err)
					log.Println(err)
					continue
				}

				manifest.ModPath = bundlesFolder
				if isLocal {
					manifest.Path = bundlePath
				} else {
					manifest.Path = filepath.Join(mainAddress, "files", "bundle", manifest.Key)
					manifest.FilePath = manifest.Path
				}

				bundleManifests = append(bundleManifests, manifest)
				bundleLoaded++
			}
		}
	}

	endTime := time.Now()
	fmt.Printf("[BUNDLE LOADER : COMPLETE] %d of %d bundles loaded in %s\n", bundleLoaded, totalBundles, endTime.Sub(startTime))

}

//TODO: When data item is being modified (not cloned) add here with mod-name
// If it's being overwritten then add to `mod-compatibility.log` for launcher
// to notify creators to adjust, then skip current mod trying to modify

var itemModificationLog = map[string]string{}
var modCritiqueLog = make(map[string][]string)

var itemsClone = make(map[string]*ModdingAPI)
var itemsEdit = make(map[string]*ModdingAPI)

func ParseCustomItemAPI(customDirectory string) map[string]*ModdingAPI {
	itemFilesDir := filepath.Join(customDirectory, "items")
	itemFiles, err := tools.GetFilesFrom(itemFilesDir)
	if err != nil {
		log.Println(err)
		return nil
	}

	customItems := make(map[string]*ModdingAPI)
	for file := range itemFiles {
		filePath := filepath.Join(itemFilesDir, file)
		if err := json.Unmarshal(tools.GetJSONRawMessage(filePath), &customItems); err != nil {
			log.Println(err)
			return nil
		}
	}
	return customItems
}

const overwriteNotification = "%s is trying to overwrite %s for item %s. The mod developers will be notified to apply changes!"

func SortAndQueueCustomItems(modName string, items map[string]*ModdingAPI) {
	if _, ok := modCritiqueLog[modName]; !ok {
		modCritiqueLog[modName] = make([]string, 0)
	}

	for key, customItem := range items {
		if customItem.Parameters.ModifierType != "edit" {
			itemModificationLog[key] = modName
			itemsClone[key] = customItem
			continue
		}

		//TODO: Separate Parameters by API? We'll see...
		//TODO: improve mod improvement logs for modders

		if customItem.Parameters.ItemParameters != nil {
			if mod, ok := itemModificationLog[customItem.Parameters.ItemParameters.ReferenceItemTPL]; ok {
				err := fmt.Sprintf(overwriteNotification, modName, mod, customItem.Parameters.ItemParameters.ReferenceItemTPL)
				log.Println(err)

				modCritiqueLog[modName] = append(modCritiqueLog[modName], err)

				if _, ok := modCritiqueLog[mod]; !ok {
					modCritiqueLog[mod] = make([]string, 0)
				}
				modCritiqueLog[mod] = append(modCritiqueLog[mod], err)

				log.Println("Skipping overwrite, continuing...")
				continue
			}
			itemModificationLog[customItem.Parameters.ItemParameters.ReferenceItemTPL] = modName
			itemsEdit[key] = customItem
		} else {
			//TODO: Custom whatever.... probably Trader or some shit
			continue
		}
	}
	cachedResponses.Save = true
	cachedResponses.Overwrite["/client/items"] = nil
	cachedResponses.Overwrite["/client/handbook/templates"] = nil
	cachedResponses.Overwrite["/client/locale/"] = nil
}

func (i *DatabaseItem) GenerateTraderAssortSingleItem() []*AssortItem {
	itemID := tools.GenerateMongoID()
	assortItem := &AssortItem{
		ID:       itemID,
		Tpl:      i.ID,
		ParentID: "hideout",
		SlotID:   "hideout",
	}

	upd, err := i.CreateItemUPD()
	if err != nil {
		return []*AssortItem{assortItem}
	}
	assortItem.Upd = upd

	return []*AssortItem{assortItem}
}

func (i *DatabaseItem) GenerateTraderAssortParentItem() []*AssortItem {
	itemID := tools.GenerateMongoID()
	assortItem := &AssortItem{
		ID:       itemID,
		Tpl:      i.ID,
		ParentID: "hideout",
		SlotID:   "hideout",
	}

	presetItem := &itemPresetItem{
		ID:  itemID,
		Tpl: i.ID,
	}

	upd, err := i.CreateItemUPD()
	if err != nil {
		return []*AssortItem{assortItem}
	}
	assortItem.Upd = upd
	presetItem.Upd = upd
	if presetItem.Upd.StackObjectsCount != 0 {
		presetItem.Upd.StackObjectsCount = 0
	}
	if presetItem.Upd.BuyRestrictionMax != 0 {
		presetItem.Upd.BuyRestrictionMax = 0
	}
	if presetItem.Upd.BuyRestrictionCurrent != 0 {
		presetItem.Upd.BuyRestrictionCurrent = 0
	}

	globalPresetItems = append(globalPresetItems, presetItem)

	return []*AssortItem{assortItem}
}

func ProcessCustomItemSet(parentId string, set map[string]any) []*AssortItem {
	if set == nil {
		return make([]*AssortItem, 0)
	}
	output := make([]*AssortItem, 0, len(set))

	for slotID, value := range set {
		setData, ok := value.(map[string]any)
		if !ok {
			log.Println()
		}

		id := tools.GenerateMongoID()
		preset := &itemPresetItem{
			ID:       id,
			Tpl:      setData["_tpl"].(string),
			ParentID: parentId,
			SlotID:   slotID,
		}

		attachment := &AssortItem{
			ID:       id,
			Tpl:      setData["_tpl"].(string),
			ParentID: parentId,
			SlotID:   slotID,
		}

		output = append(output, attachment)
		globalPresetItems = append(globalPresetItems, preset)

		attachments, ok := setData["attachments"].(map[string]any)
		if !ok {
			continue
		}

		subAttachments := ProcessCustomItemSet(attachment.ID, attachments)
		output = append(output, subAttachments...)
	}

	return output
}

var globalPresetItems []*itemPresetItem

type globalItemPreset struct {
	Id               string            `json:"_id"`
	Type             string            `json:"_type"`
	ChangeWeaponName bool              `json:"_changeWeaponName"`
	Name             string            `json:"_name"`
	Encyclopedia     string            `json:"_encyclopedia"`
	Parent           string            `json:"_parent"`
	Items            []*itemPresetItem `json:"_items"`
}

type itemPresetItem struct {
	ID       string      `json:"_id"`
	Tpl      string      `json:"_tpl"`
	ParentID string      `json:"parentId,omitempty"`
	SlotID   string      `json:"slotId,omitempty"`
	Upd      *ItemUpdate `json:"upd,omitempty"`
}

func GenerateGlobalsItemPresetEntry(preset *CustomItemPreset) {
	presetId := tools.GenerateMongoID()

	gIP := &globalItemPreset{
		Id:               presetId,
		Type:             "Preset",
		ChangeWeaponName: preset.ChangeWeaponName,
		Name:             preset.Name,
		Items:            globalPresetItems,
	}

	AddToItemPresets(presetId, *gIP)
	globalPresetItems = nil
}

func (i *DatabaseItem) GenerateTraderAssortPresetItem(preset *CustomItemPreset) []*AssortItem {
	globalPresetItems = make([]*itemPresetItem, 0, len(preset.Items))

	family := i.GenerateTraderAssortParentItem()
	children := ProcessCustomItemSet(family[0].ID, preset.Items)
	family = append(family, children...)

	return family
}

func GetTraderFromNameOrUID(input string) (*Trader, error) {
	var trader *Trader
	trader, err := GetTraderByUID(input)
	if err != nil {
		trader, err = GetTraderByName(input)
		if err != nil {
			return nil, err
		}
	}
	return trader, nil
}

func (i *DatabaseItem) GenerateTraderAssortEntry(params *CustomItemParams) {
	for tid, traderParams := range params.AddToTrader {
		trader, err := GetTraderFromNameOrUID(tid)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, barter := range traderParams {
			var assortItem []*AssortItem
			if barter.Preset != nil {
				assortItem = i.GenerateTraderAssortPresetItem(barter.Preset)
				GenerateGlobalsItemPresetEntry(barter.Preset)
			} else {
				assortItem = i.GenerateTraderAssortSingleItem()
			}

			schemes := make([]*Scheme, 0, len(barter.BarterScheme))

			parent := assortItem[0]
			if parent.Upd == nil {
				parent.Upd = new(ItemUpdate)
			}
			parent.Upd.StackObjectsCount = int32(barter.AmountInStock)

			if trader.Assort.BarterScheme[parent.ID] == nil {
				trader.Assort.BarterScheme[parent.ID] = make([][]*Scheme, 0, 1)
			}

			for bid, value := range barter.BarterScheme {
				scheme := &Scheme{
					Tpl:   bid,
					Count: value,
				}

				schemes = append(schemes, scheme)
			}

			trader.Assort.LoyalLevelItems[parent.ID] = barter.LoyaltyLevel
			trader.Assort.BarterScheme[parent.ID] = append(trader.Assort.BarterScheme[parent.ID], schemes)
			trader.Assort.Items = append(trader.Assort.Items, assortItem...)
		}
	}
}

func (i *DatabaseItem) SetCustomOverrides(overrides map[string]any) {
	for key, value := range overrides {
		switch key {
		case "_name":
			i.Name = value.(string)
			continue
		case "_parent":
			i.Parent = value.(string)
			continue
		case "_type":
			i.Type = value.(string)
			continue
		case "_proto":
			i.Proto = value.(string)
			continue
		default:
			if _, ok := i.Props[key]; !ok {
				log.Println("Could not override property", key, "because it does not exist on the item")
			} else {
				i.Props[key] = value
			}
			continue
		}
	}
}

const (
	localeName        = "%s Name"
	localeShortName   = "%s ShortName"
	localeDescription = "%s Description"
)

func LoadCustomItems() {
	startTime := time.Now()

	itemsDatabase := GetItems()
	handbookItemsDatabase := GetHandbook().Items

	custBody, err := GetCustomizationByID("64ef3efdb63b74469b6c1499")
	if err != nil {
		log.Println(err)
		return
	}

	custHands, err := GetCustomizationByID("64ef3efdb63b74469b6c1499")
	if err != nil {
		log.Println(err)
		return
	}

	custUpper, err := GetCustomizationByID("64ef3efdb63b74469b6c1499")
	if err != nil {
		log.Println(err)
		return
	}

	custFeet, err := GetCustomizationByID("5d5e7f4986f7746956659f8a")
	if err != nil {
		log.Println(err)
		return
	}

	custLower, err := GetCustomizationByID("5cd946231388ce000d572fe3")
	if err != nil {
		log.Println(err)
		return
	}

	customization := GetCustomizations()

	for uid, api := range itemsClone {
		switch {
		case api.Parameters.ItemParameters != nil:
			if api.Locale == nil || len(api.Locale) == 0 {
				log.Println(uid, "does not have a locale, skipping...")
				continue
			}

			original := itemsDatabase[api.Parameters.ItemParameters.ReferenceItemTPL]
			if original == nil {
				log.Println("ReferenceItemTPL:", api.Parameters.ItemParameters.ReferenceItemTPL, "for UID:", uid, "is invalid, skipping...")
				continue
			}
			itemClone := original.Clone()
			itemClone.ID = uid

			itemClone.GenerateTraderAssortEntry(api.Parameters.ItemParameters)
			if api.Overrides != nil || len(api.Overrides) != 0 {
				itemClone.SetCustomOverrides(api.Overrides)
			}

			handbookEntry, err := original.GetHandbookItemEntry()
			if err != nil {
				log.Println(err)
				continue
			}

			itemsDatabase[uid] = itemClone

			handbookItemsDatabase = append(handbookItemsDatabase, HandbookItem{
				ID:       uid,
				ParentID: handbookEntry.ParentID,
				Price:    int32(api.Parameters.HandbookPrice),
			})

			setCustomItemLocale(uid, api.Locale)
		case api.Parameters.ClothingParameters != nil:
			if api.Parameters.ClothingParameters.Side == nil || len(api.Parameters.ClothingParameters.Side) == 0 {
				log.Println(uid, "does not have Side, skipping...")
				continue
			}

			params := api.Parameters.ClothingParameters
			if params.LowerBody == nil && params.UpperBody == nil {
				log.Println(uid, "does not have LowerBody or UpperBody entries, skipping...")
				continue
			}

			locales := map[string]string{}
			if params.UpperBody != nil {
				locales[params.UpperBody.ID] = params.UpperBody.Name
				upperBodySuit := TraderSuits{
					ID:       params.UpperBody.Body.Id,
					SuiteID:  params.UpperBody.ID,
					IsActive: true,
				}

				for tid, scheme := range params.UpperBody.AddToTrader {
					trader, err := GetTraderFromNameOrUID(tid)
					if err != nil {
						log.Println(err)
						continue
					} else if trader.Suits == nil {
						log.Println(tid, "does not have suits assort, skipping...")
						continue
					}

					upperBodySuit.Tid = trader.Base.ID

					for _, required := range scheme {
						requirements := SuitRequirements{
							LoyaltyLevel:         required.LoyaltyLevel,
							ProfileLevel:         required.PMCLevel,
							Standing:             required.TraderStanding,
							SkillRequirements:    required.SkillRequirements,
							QuestRequirements:    required.QuestRequirements,
							SuitItemRequirements: make([]SuitItemRequirements, 0, len(required.ItemRequirements)),
						}

						for key, value := range required.ItemRequirements {
							itemRequirement := SuitItemRequirements{
								Count:          int(value),
								Tpl:            key,
								OnlyFunctional: true,
							}
							requirements.SuitItemRequirements = append(requirements.SuitItemRequirements, itemRequirement)
						}

						upperBodySuit.Requirements = requirements
						trader.Suits = append(trader.Suits, upperBodySuit)
					}
				}

				body := custBody.Clone()
				hands := custHands.Clone()
				upper := custUpper.Clone()

				body.ID = params.UpperBody.Body.Id
				body.Name = params.UpperBody.Body.Id
				body.Props.Prefab = map[string]any{
					"path": params.UpperBody.Body.Prefab,
					"rcid": "",
				}
				body.Props.Side = params.Side

				hands.ID = params.UpperBody.Hands.Id
				hands.Name = params.UpperBody.Hands.Id
				hands.Props.Side = params.Side
				hands.Props.Prefab = map[string]any{
					"path": params.UpperBody.Hands.Prefab,
					"rcid": "",
				}

				upper.ID = params.UpperBody.ID
				upper.Name = params.UpperBody.ID
				upper.Props.Body = body.ID
				upper.Props.Hands = hands.ID
				upper.Props.Side = params.Side

				customization[body.ID] = body
				customization[hands.ID] = hands
				customization[upper.ID] = upper
			}

			if params.LowerBody != nil {
				locales[params.LowerBody.Id] = params.LowerBody.Name
				lowerBodySuit := TraderSuits{
					ID:       params.LowerBody.Feet.Id,
					SuiteID:  params.LowerBody.Id,
					IsActive: true,
				}

				for tid, scheme := range params.LowerBody.AddToTrader {
					trader, err := GetTraderFromNameOrUID(tid)
					if err != nil {
						log.Println(err)
						continue
					} else if trader.Suits == nil {
						log.Println(tid, "does not have suits assort, skipping...")
						continue
					}

					lowerBodySuit.Tid = trader.Base.ID

					for _, required := range scheme {
						requirements := SuitRequirements{
							LoyaltyLevel:         required.LoyaltyLevel,
							ProfileLevel:         required.PMCLevel,
							Standing:             required.TraderStanding,
							SkillRequirements:    required.SkillRequirements,
							QuestRequirements:    required.QuestRequirements,
							SuitItemRequirements: make([]SuitItemRequirements, 0, len(required.ItemRequirements)),
						}

						for key, value := range required.ItemRequirements {
							itemRequirement := SuitItemRequirements{
								Count:          int(value),
								Tpl:            key,
								OnlyFunctional: true,
							}
							requirements.SuitItemRequirements = append(requirements.SuitItemRequirements, itemRequirement)
						}

						lowerBodySuit.Requirements = requirements
						trader.Suits = append(trader.Suits, lowerBodySuit)
					}
				}

				feet := custFeet.Clone()
				feet.ID = params.LowerBody.Feet.Id
				feet.Name = params.LowerBody.Feet.Id
				feet.Props.Side = params.Side
				feet.Props.Prefab = map[string]any{
					"path": params.LowerBody.Feet.Prefab,
					"rcid": "",
				}
				customization[feet.ID] = feet

				lower := custLower.Clone()
				lower.ID = params.LowerBody.Id
				lower.Name = params.LowerBody.Id
				lower.Props.Side = params.Side
				lower.Props.Feet = feet.ID

				customization[lower.ID] = lower
			}

			setCustomClothingLocation(locales)
		//case api.Parameters.TraderParameters != nil
		//case api.Parameters.RoutingParameters != nil
		//case api.Parameters.???Parameters != nil
		//case api.Parameters.???Parameters != nil
		default:
			log.Println("HAHHAHAHAHAHAHA")
			continue
		}
	}

	//TODO: we do massive recursion because hahahahahah

	/*	for uid, api := range itemsEdit {

		}*/

	endTime := time.Now()
	fmt.Printf("[CUSTOM ITEM LOADER : COMPLETE] in %s\n\n", endTime.Sub(startTime))
}

var mainLocales = [4]string{"en", "ge", "fr", "ru"}

func setCustomClothingLocation(ids map[string]string) {
	formatted := make(map[string]string)
	for key, value := range ids {
		formatted[fmt.Sprintf(localeName, key)] = value
		formatted[fmt.Sprintf(localeShortName, key)] = ""
	}

	for _, lang := range mainLocales {
		data, err := GetLocalesGlobalByName(lang)
		if err != nil {
			log.Println(err)
			continue
		}

		for name, value := range formatted {
			data[name] = value
		}
	}
}

func setCustomItemLocale(uid string, apiLocale map[string]*CustomItemLocale) {
	localeName := fmt.Sprintf(localeName, uid)
	localeShortName := fmt.Sprintf(localeShortName, uid)
	localeDescription := fmt.Sprintf(localeDescription, uid)

	if len(apiLocale) == 1 {
		var nameValue string
		var shortNameValue string
		var descriptionValue string

		for _, value := range apiLocale {
			nameValue = value.Name
			shortNameValue = value.ShortName
			descriptionValue = value.Description
		}

		for _, lang := range mainLocales {
			data, err := GetLocalesGlobalByName(lang)
			if err != nil {
				log.Println(err)
				continue
			}
			data[localeName] = nameValue
			data[localeShortName] = shortNameValue
			data[localeDescription] = descriptionValue
		}
		return
	}

	for lang, value := range apiLocale {
		locale, err := GetLocalesGlobalByName(lang)
		if err != nil {
			log.Println(err)
			continue
		}

		locale[localeName] = value.Name
		locale[localeShortName] = value.ShortName
		locale[localeDescription] = value.Description
	}

}

type ModdingAPI struct {
	Parameters Parameters
	Overrides  map[string]any               `json:"overrides,omitempty"`
	Locale     map[string]*CustomItemLocale `json:"locale,omitempty"`
}

type Parameters struct {
	HandbookPrice      int                   `json:",omitempty"`
	ModifierType       string                `json:",omitempty"`
	ItemParameters     *CustomItemParams     `json:",omitempty"`
	ClothingParameters *CustomClothingParams `json:",omitempty"`
}

type CustomItemParams struct {
	ReferenceItemTPL            string                        `json:",omitempty"`
	AddToTrader                 map[string][]*ItemAddToTrader `json:",omitempty"`
	AdditionalItemCompatibility []*string                     `json:",omitempty"`
}

type CustomItemPreset struct {
	Id               string         `json:"_id"`
	Type             string         `json:"_type"`
	ChangeWeaponName bool           `json:"_changeWeaponName"`
	Name             string         `json:"_name"`
	Encyclopedia     string         `json:"_encyclopedia"`
	Parent           string         `json:"_parent"`
	Items            map[string]any `json:"_items"`
}

type ItemAddToTrader struct {
	//Set allows you to create a full weapon for the trader
	Preset        *CustomItemPreset `json:",omitempty"`
	Set           map[string]any    `json:",omitempty"`
	LoyaltyLevel  int8
	BarterScheme  map[string]float32
	AmountInStock int16
}

type CustomClothingParams struct {
	Head      *any
	UpperBody *UpperBodySuite `json:",omitempty"`
	LowerBody *LowerBodySuite `json:",omitempty"`
	Side      []string        `json:",omitempty"`
}

type UpperBodySuite struct {
	ID          string
	Name        string
	Body        CustomSuite
	Hands       CustomSuite
	AddToTrader map[string][]*ClothingAddToTrader
}
type LowerBodySuite struct {
	Id          string
	Name        string
	Feet        CustomSuite
	AddToTrader map[string][]*ClothingAddToTrader
}
type CustomSuite struct {
	Id     string
	Prefab string
}

type ClothingAddToTrader struct {
	Id                string
	SuiteId           string
	LoyaltyLevel      int8
	TraderStanding    float32
	PMCLevel          int8
	QuestRequirements []string
	SkillRequirements []map[string]int8
	ItemRequirements  map[string]int32
}

type CustomItemLocale struct {
	Name        string
	ShortName   string
	Description string
}
