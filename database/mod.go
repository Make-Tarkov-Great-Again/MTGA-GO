package database

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"path/filepath"
	"strings"
	"time"
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
	bundleManifests = nil
}

func AddModBundleDirPath(modBundleDirPath string) {
	modBundleDirPaths = append(modBundleDirPaths, modBundleDirPath)
}

func SetBundleManifests() {
	if len(modBundleDirPaths) == 0 {
		return
	}
	defer func() {
		modBundleDirPaths = nil
	}()

	startTime := time.Now()
	bundleLoaded := 0
	totalBundles := 0

	fmt.Printf("\n[BUNDLELOADER : BEGIN]\n")

	isLocal := GetServerConfig().IP == "127.0.0.1"
	var mainAddress string
	if !isLocal {
		mainAddress = GetMainAddress()
	}

	for _, path := range modBundleDirPaths {
		bundlesSubDirectories, err := tools.GetDirectoriesFrom(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		for subDir := range bundlesSubDirectories {
			bundleMainDirPath := filepath.Join(path, subDir)
			bundlesJsonPath := filepath.Join(bundleMainDirPath, "bundles.json")
			if !tools.FileExist(bundlesJsonPath) {
				err = fmt.Errorf("bundles.json file not located in %s, returning", path)
				fmt.Println(err)
				return
			}

			manifests := new(Manifests)
			data := tools.GetJSONRawMessage(bundlesJsonPath)
			if err := json.Unmarshal(data, &manifests); err != nil {
				fmt.Println(err)
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
					fmt.Println(err)
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
	fmt.Printf("[BUNDLELOADER : COMPLETE] %d of %d bundles loaded in %s\n\n", bundleLoaded, totalBundles, endTime.Sub(startTime))

}

//TODO: When database item is being modified (not cloned) add here with mod-name
// If it's being overwritten then add to `mod-compatibility.log` for launcher
// to notify creators to adjust, then skip current mod trying to modify

var itemModificationLog = map[string]string{}
var modCritiqueLog = make(map[string][]string)

var itemsClone = make(map[string]*CustomItemAPI)
var itemsEdit = make(map[string]*CustomItemAPI)

const overwriteNotification = "%s is trying to overwrite %s for item %s. The mod developers will be notified to apply changes!"

func SortAndQueueCustomItems(modName string, items map[string]*CustomItemAPI) {
	if _, ok := modCritiqueLog[modName]; !ok {
		modCritiqueLog[modName] = make([]string, 0)
	}

	for key, customItem := range items {
		if customItem.Parameters.ModifierType != "edit" {
			itemModificationLog[key] = modName
			itemsClone[key] = customItem
			continue
		}

		//TODO: improve mod improvement logs for modders

		if mod, ok := itemModificationLog[customItem.Parameters.ReferenceItemTPL]; ok {
			err := fmt.Sprintf(overwriteNotification, modName, mod, customItem.Parameters.ReferenceItemTPL)
			fmt.Println(err)

			modCritiqueLog[modName] = append(modCritiqueLog[modName], err)

			if _, ok := modCritiqueLog[itemModificationLog[customItem.Parameters.ReferenceItemTPL]]; !ok {
				modCritiqueLog[itemModificationLog[customItem.Parameters.ReferenceItemTPL]] = make([]string, 0)
			}
			modCritiqueLog[itemModificationLog[customItem.Parameters.ReferenceItemTPL]] = append(modCritiqueLog[itemModificationLog[customItem.Parameters.ReferenceItemTPL]], err)

			fmt.Println("Skipping overwrite, continuing...")
			continue
		}
		itemModificationLog[customItem.Parameters.ReferenceItemTPL] = modName
		itemsEdit[key] = customItem
	}
}

func (i *DatabaseItem) GenerateTraderAssortItem() *AssortItem {
	assortItem := new(AssortItem)
	assortItem.ID = tools.GenerateMongoID()
	assortItem.Tpl = i.ID
	assortItem.ParentID = "hideout"
	assortItem.SlotID = "hideout"
	assortItem.Upd = *i.GenerateNewUPD()

	return assortItem
}

func (i *DatabaseItem) GenerateTraderAssortEntry(params customItemParams) {
	assortItem := i.GenerateTraderAssortItem()

	for tid, traderParams := range params.AddToTrader {
		schemes := map[string][]*Scheme{}

		assort := GetTraderByUID(tid).Assort
		if assort == nil {
			tid = *GetTraderIDByName(tid)
			assort = GetTraderByUID(tid).Assort
		}

		schemes[tid] = make([]*Scheme, 0, len(traderParams.BarterScheme))
		for bid, value := range traderParams.BarterScheme {
			scheme := new(Scheme)
			scheme.Tpl = bid
			scheme.Count = value

			schemes[tid] = append(schemes[tid], scheme)
		}
		assort.LoyalLevelItems[i.ID] = traderParams.LoyaltyLevel

		assort.Items = append(assort.Items, assortItem)
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
				fmt.Println("Could not override property", key, "because it does not exist on the item")
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

func ProcessCustomItems() {
	itemsDatabase := GetItems()

	for uid, api := range itemsClone {
		if api.Locale == nil || len(api.Locale) == 0 {
			fmt.Println(uid, "does not have a locale, skipping...")
			continue
		}

		itemClone := itemsDatabase[api.Parameters.ReferenceItemTPL].Clone()
		itemClone.ID = uid

		itemClone.GenerateTraderAssortEntry(api.Parameters)
		if api.Overrides != nil || len(api.Overrides) != 0 {
			itemClone.SetCustomOverrides(api.Overrides)
		}

		setCustomItemLocale(uid, api.Locale)
	}

	//TODO: we do massive recursion because hahahahahah

	/*	for uid, api := range itemsEdit {

		}*/
}

func setCustomItemLocale(uid string, apiLocale map[string]*CustomItemLocale) {
	name := fmt.Sprintf(localeName, uid)
	shortName := fmt.Sprintf(localeShortName, uid)
	description := fmt.Sprintf(localeDescription, uid)

	if len(apiLocale) == 1 {
		modname, ok := itemModificationLog[uid]
		if !ok {
			fmt.Println("how")
		}
		if _, ok := modCritiqueLog[modname]; !ok {
			modCritiqueLog[modname] = make([]string, 0)
		}
		modCritiqueLog[modname] = append(modCritiqueLog[modname], "It is encouraged to have en, de, fr, ru locales!")

		locales := GetLocales()
		var nameValue string
		var shortNameValue string
		var descriptionValue string

		for _, value := range apiLocale {
			nameValue = value.Name
			shortNameValue = value.ShortName
			descriptionValue = value.Description
		}

		for _, data := range locales {
			data.Locale[name] = nameValue
			data.Locale[shortName] = shortNameValue
			data.Locale[description] = descriptionValue
		}
	} else {
		for lang, value := range apiLocale {
			locale := GetLocalesLocaleByName(lang)
			locale[name] = value.Name
			locale[shortName] = value.ShortName
			locale[description] = value.Description
		}
	}
}

type CustomItemAPI struct {
	API        string
	Parameters CustomItemParams
	Overrides  map[string]any `json:"overrides,omitempty"`
	Locale     map[string]*CustomItemLocale
}

type CustomItemParams struct {
	ReferenceClothingTpl        string `json:",omitempty"`
	ReferenceItemTPL            string `json:",omitempty"`
	HandbookPrice               int    `json:",omitempty"`
	ModifierType                string
	AddToTrader                 map[string]*CustomItemAddToTrader `json:",omitempty"`
	AdditionalItemCompatibility []*string                         `json:",omitempty"`
	ItemPresets                 map[string]map[string]any         `json:",omitempty"`
}

type CustomItemAddToTrader struct {
	LoyaltyLevel  int8
	BarterScheme  map[string]float32
	AmountInStock int16
}

type CustomItemLocale struct {
	Name        string
	ShortName   string
	Description string
}
