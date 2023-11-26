package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"
	"slices"
)

const (
	inventoryCacheNotExist string = "Inventory Cache for %s does not exist"
	traderCacheNotExist    string = "Trader Cache for %s does not exist"
	questCacheNotExist     string = "Quest Cache for %s does not exist"
	cacheNotExist          string = "Cache for %s does not exist"
)

func GetCacheByID(uid string) (*PlayerCache, error) {
	if cache, ok := db.cache.player[uid]; ok {
		return cache, nil
	}
	return nil, fmt.Errorf(cacheNotExist, uid)
}

func GetQuestCacheByID(uid string) (*QuestCache, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Quests != nil {
		return cache.Quests, nil
	}

	return nil, fmt.Errorf(questCacheNotExist, uid)
}

func GetTraderCacheByID(uid string) (*TraderCache, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Traders != nil {
		return cache.Traders, nil
	}
	return nil, fmt.Errorf(traderCacheNotExist, uid)
}

func GetInventoryCacheByID(uid string) (*InventoryContainer, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Inventory != nil {
		return cache.Inventory, nil
	}
	return nil, fmt.Errorf(inventoryCacheNotExist, uid)
}

func SetProfileCache(id string) {
	if _, ok := db.cache.player[id]; !ok && db.profile[id].Character == nil {
		return
	}

	cache := &PlayerCache{
		Traders: &TraderCache{
			Index:      make(map[string]*AssortIndex),
			Assorts:    make(map[string]*Assort),
			Insurances: make(map[string]*Insurances),
		},
	}
	if db.profile[id].Character != nil {
		cache.SetCharacterCache(db.profile[id].Character)
	}
	db.cache.player[id] = cache
}

func (c *PlayerCache) SetCharacterCache(character *Character) {
	done := make(chan struct{})
	go func() {
		if len(character.Quests) != 0 {
			c.Quests = &QuestCache{Index: make(map[string]int8)}
			for index, quest := range character.Quests {
				c.Quests.Index[quest.QID] = int8(index)
			}
		}

		done <- struct{}{}
	}()

	// Define a function to update the common skills map
	go func() {
		if len(character.Skills.Common) != 0 {
			c.Skills = &SkillsCache{Common: make(map[string]int8)}
			for index, commonSkill := range character.Skills.Common {
				c.Skills.Common[commonSkill.ID] = int8(index)
			}
		}
		done <- struct{}{}
	}()

	// Define a function to update the hideout areas map
	go func() {
		if len(character.Hideout.Areas) != 0 {
			c.Hideout = &HideoutCache{Areas: make(map[int8]int8)}
			for index, area := range character.Hideout.Areas {
				c.Hideout.Areas[int8(area.Type)] = int8(index)
			}
		}
		done <- struct{}{}
	}()

	go func() {
		if len(character.Inventory.Items) != 0 {
			c.Inventory = SetInventoryContainer(&character.Inventory)
		}
		done <- struct{}{}
	}()

	for i := 0; i < 4; i++ {
		<-done
	}

	db.cache.player[character.ID] = c
}

func SetInventoryContainer(inventory *Inventory) *InventoryContainer {
	output := &InventoryContainer{}

	output.SetInventoryIndex(inventory)
	output.SetInventoryStash(inventory)

	return output
}

func (ic *InventoryContainer) SetInventoryStash(inventory *Inventory) {
	var stash *Stash
	if ic.Stash == nil {
		ic.Stash = &Stash{}
		stash = ic.Stash

		item, err := GetItemByID(inventory.Items[ic.Lookup.Forward[inventory.Stash]].TPL)
		if err != nil {
			log.Println(err)
			return
		}
		grids := item.GetItemGrids()

		for key, value := range grids {
			stash.SlotID = key

			height := value.Props.CellsV
			width := value.Props.CellsH

			arraySize := int(height) * int(width)

			stash.Container = Map{
				Height:  height,
				Width:   width,
				Map:     make([]string, arraySize),
				FlatMap: make(map[string]FlatMapLookup),
			}
		}
	} else {
		stash = ic.Stash
	}

	var containerMap = ic.Stash.Container.Map
	var containerFlatMap = ic.Stash.Container.FlatMap
	var stride = int16(ic.Stash.Container.Width)
	var itemID string

	for index := range ic.Lookup.Reverse {
		itemInInventory := inventory.Items[index]
		if itemInInventory.ParentID == "" || itemInInventory.ParentID != inventory.Stash || itemInInventory.SlotID != "hideout" || itemInInventory.Location == nil {
			continue
		}

		height, width := ic.MeasureItemForInventoryMapping(inventory.Items, itemInInventory.ID)
		if height == -1 && width == -1 {
			return
		}

		itemFlatMap := *ic.CreateFlatMapLookup(height, width, &itemInInventory)

		if itemFlatMap.Height == 0 && itemFlatMap.Width == 0 {
			if itemID = containerMap[itemFlatMap.StartX]; itemID != "" {
				log.Println("Flat Map Index of", itemFlatMap.StartX, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[itemFlatMap.StartX])
				return
			}
			containerMap[itemFlatMap.StartX] = itemInInventory.ID
			itemFlatMap.Coordinates = append(itemFlatMap.Coordinates, itemFlatMap.StartX)
			containerFlatMap[itemInInventory.ID] = itemFlatMap

			continue
		}

		for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
			if itemID = containerMap[column]; itemID != "" {
				log.Println("Flat Map Index of X position", column, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[column])
				return
			}
			containerMap[column] = itemInInventory.ID
			itemFlatMap.Coordinates = append(itemFlatMap.Coordinates, column)

			for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
				var coordinate = row*stride + column
				if itemID = containerMap[coordinate]; itemID != "" {
					log.Println("Flat Map Index of Y position", row, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[coordinate])
					return
				}
				containerMap[coordinate] = itemInInventory.ID
				itemFlatMap.Coordinates = append(itemFlatMap.Coordinates, coordinate)
			}
		}

		containerFlatMap[itemInInventory.ID] = itemFlatMap
	}
}

func (ic *InventoryContainer) CreateFlatMapLookup(height int8, width int8, itemInInventory *InventoryItem) *FlatMapLookup {
	output := new(FlatMapLookup)

	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	if itemInInventory.Location.R.(float64) == 1 {
		output.Height = width
		output.Width = height
	} else {
		output.Height = height
		output.Width = width
	}

	row := int16(itemInInventory.Location.Y.(float64)) * int16(ic.Stash.Container.Width)
	output.StartX = int16(itemInInventory.Location.X.(float64)) + row
	output.EndX = output.StartX + int16(output.Width)

	return output
}

//TODO: Check consistency of this code
// I might need to check if oldX = newX then check heights

// ResetItemSizeInContainer resets item size in InventoryContainer to reflect item size change
func (ic *InventoryContainer) ResetItemSizeInContainer(itemInInventory *InventoryItem, Inventory *Inventory) {
	var stash = *ic.Stash
	var itemFlatMap = stash.Container.FlatMap[itemInInventory.ID]
	var containerMap = &stash.Container.Map
	var stride = int16(stash.Container.Width)

	height, width := ic.MeasureItemForInventoryMapping(Inventory.Items, itemInInventory.ID)

	newItemFlatMap := *ic.CreateFlatMapLookup(height, width, itemInInventory)

	var toDelete = make([]int16, 0, len(itemFlatMap.Coordinates))
	var toAdd = make([]int16, 0, len(itemFlatMap.Coordinates))

	if newItemFlatMap.EndX < itemFlatMap.EndX {
		for column := newItemFlatMap.EndX + 1; column <= itemFlatMap.EndX; column++ {
			(*containerMap)[column] = ""
			toDelete = append(toDelete, column)

			if newItemFlatMap.Height < itemFlatMap.Height {
				for row := int16(newItemFlatMap.Height) + 1; row <= int16(itemFlatMap.Height); row++ {
					var coordinate = row*stride + itemFlatMap.EndX
					(*containerMap)[coordinate] = ""
					toDelete = append(toDelete, coordinate)
				}
			}
		}
	} else if newItemFlatMap.EndX > itemFlatMap.EndX {
		for column := itemFlatMap.EndX + 1; column <= newItemFlatMap.EndX; column++ {
			(*containerMap)[column] = itemInInventory.ID
			toAdd = append(toAdd, column)

			if newItemFlatMap.Height > itemFlatMap.Height {
				for row := int16(itemFlatMap.Height) + 1; row <= int16(newItemFlatMap.Height); row++ {
					var coordinate = row*stride + itemFlatMap.EndX
					(*containerMap)[coordinate] = itemInInventory.ID
					toAdd = append(toAdd, coordinate)

				}
			}
		}
	}

	if len(toDelete) != 0 {
		coordinates := make([]int16, 0, len(itemFlatMap.Coordinates)-len(toDelete))
		for _, value := range itemFlatMap.Coordinates {
			if slices.Contains(toDelete, value) {
				continue
			}
			coordinates = append(coordinates, value)
		}

		newItemFlatMap.Coordinates = coordinates
	} else if len(toAdd) != 0 {
		coordinates := make([]int16, 0, len(itemFlatMap.Coordinates)+len(toAdd))
		coordinates = append(coordinates, itemFlatMap.Coordinates[0:]...)
		coordinates = append(coordinates, toAdd[0:]...)

		newItemFlatMap.Coordinates = coordinates
	}

	stash.Container.FlatMap[itemInInventory.ID] = newItemFlatMap
	if _, exist := ic.Lookup.Forward[itemInInventory.ID]; !exist {
		ic.SetInventoryIndex(Inventory)
	}
}

func (ic *InventoryContainer) GenerateCoordinatesFromLocation(flatMap FlatMapLookup) []int16 {
	output := make([]int16, 0, flatMap.Height+flatMap.Width)

	for c := flatMap.StartX; c <= flatMap.EndX; c++ {
		output = append(output, c)

		for r := int16(1); r <= int16(flatMap.Height); r++ {
			coordinate := r*int16(ic.Stash.Container.Width) + c
			output = append(output, coordinate)
		}
	}

	return output
}

func (ic *InventoryContainer) SetNewItemFlatMap(items []InventoryItem) {
	height, width := MeasurePurchaseForInventoryMapping(items)
	itemInInventory := items[len(items)-1]
	flatMap := *ic.CreateFlatMapLookup(height, width, &itemInInventory)
	flatMap.Coordinates = ic.GenerateCoordinatesFromLocation(flatMap)
	ic.Stash.Container.FlatMap[itemInInventory.ID] = flatMap
}

// ClearItemFromContainerMap soft-deletes Item from Container.Map by removing its entries
// only and preserves its FlatMapLookup
func (ic *InventoryContainer) ClearItemFromContainerMap(UID string) {
	var stash = *ic.Stash
	itemFlatMap, ok := ic.Stash.Container.FlatMap[UID]
	if !ok {
		fmt.Println(UID)
	}
	var containerMap = &stash.Container.Map

	for _, index := range itemFlatMap.Coordinates {
		(*containerMap)[index] = ""
	}
}

// AddItemToContainerMap sets Item in Container.Map by adding its entries
// from preserved FlatMapLookup
func (ic *InventoryContainer) AddItemToContainerMap(UID string) {
	var stash = *ic.Stash
	itemFlatMap, ok := ic.Stash.Container.FlatMap[UID]
	if !ok {
		fmt.Println(UID)
	}
	var containerMap = &stash.Container.Map

	for _, index := range itemFlatMap.Coordinates {
		(*containerMap)[index] = UID
	}
}

// ClearItemFromContainer wipes item, based on the UID, from the cached InventoryContainer
//
// Warning!
// Only use this function if you're hard-deleting entries from Inventory.Items
func (ic *InventoryContainer) ClearItemFromContainer(UID string) {
	var itemFlatMap = ic.Stash.Container.FlatMap[UID]
	var containerMap = ic.Stash.Container.Map

	for _, index := range itemFlatMap.Coordinates {
		containerMap[index] = ""
	}

	if _, ok := ic.Lookup.Forward[UID]; ok {
		delete(ic.Lookup.Reverse, ic.Lookup.Forward[UID])
		delete(ic.Lookup.Forward, UID)
	}

	delete(ic.Stash.Container.FlatMap, UID)
}

// AddItemToContainer adds item, based on the UID, to the cached InventoryContainer
func (ic *InventoryContainer) AddItemToContainer(UID string, itemFlatMap *FlatMapLookup) {
	var containerMap = ic.Stash.Container.Map

	for _, index := range itemFlatMap.Coordinates {
		containerMap[index] = UID
	}

	ic.Stash.Container.FlatMap[UID] = *itemFlatMap
}

type sizes struct {
	ForcedUp    int8
	ForcedDown  int8
	ForcedRight int8
	ForcedLeft  int8
	SizeUp      int8
	SizeDown    int8
	SizeLeft    int8
	SizeRight   int8
}

// MeasureItemForInventoryMapping gets the correct size of an item within the Character.Inventory for setting in
// Stash.Container
func (ic *InventoryContainer) MeasureItemForInventoryMapping(items []InventoryItem, parent string) (int8, int8) {
	index := ic.Lookup.Forward[parent]
	itemInInventory := items[index]

	itemInDatabase, err := GetItemByID(itemInInventory.TPL) //parent
	if err != nil {
		log.Println(err)
		return -1, -1
	}

	height, width := itemInDatabase.GetItemSize()             //get parent as starting point
	if itemInDatabase.Parent == "5448e53e4bdc2d60728b4567" || //backpack
		itemInDatabase.Parent == "566168634bdc2d144c8b456c" || //searchableItem
		itemInDatabase.Parent == "5795f317245977243854e041" { //simpleContainer
		return height, width
	}

	var parentFolded = itemInInventory.UPD != nil && itemInInventory.UPD.Foldable != nil && itemInInventory.UPD.Foldable.Folded

	canFold, foldablePropertyExists := itemInDatabase.Props["Foldable"].(bool)
	foldedSlotID, foldedSlotPropertyExists := itemInDatabase.Props["FoldedSlot"].(string)

	if (foldablePropertyExists && canFold) && foldedSlotPropertyExists && parentFolded {
		sizeReduceRight, ok := itemInDatabase.Props["SizeReduceRight"].(float64)
		if !ok {
			log.Println("Could not type assert itemInDatabase.Props.SizeReduceRight of UID", itemInInventory.ID)
			return -1, -1
		}
		width -= int8(sizeReduceRight)
	}

	family := GetInventoryItemFamilyTreeIDs(items, parent)
	length := len(family) - 1

	if length == 1 {
		return height, width
	}

	var member string
	sizes := &sizes{}

	var childFolded bool
	for i := 0; i < length; i++ {
		member = family[i]
		index = ic.Lookup.Forward[member]
		itemInInventory = items[index]

		childFolded = itemInInventory.UPD != nil && itemInInventory.UPD.Foldable != nil && itemInInventory.UPD.Foldable.Folded
		if parentFolded || childFolded {
			continue
		} else if (foldablePropertyExists && canFold) &&
			itemInInventory.SlotID == foldedSlotID &&
			(parentFolded || childFolded) {
			continue
		}

		item, err := GetItemByID(itemInInventory.TPL)
		if err != nil {
			log.Println(err)
			return -1, -1
		}
		item.GetItemForcedSize(sizes)
	}

	height += sizes.SizeUp + sizes.SizeDown + sizes.ForcedDown + sizes.ForcedUp
	width += sizes.SizeLeft + sizes.SizeRight + sizes.ForcedRight + sizes.ForcedLeft

	return height, width
}

// SetSingleInventoryIndex sets new item to Lookup based on their index and their item.id
func (ic *InventoryContainer) SetSingleInventoryIndex(UID string, index int16) {
	ic.Lookup.Forward[UID] = index
	ic.Lookup.Reverse[index] = UID
}

// SetInventoryIndex set/reset InventoryContainer.Lookup for fast Inventory.Items lookup
func (ic *InventoryContainer) SetInventoryIndex(inventory *Inventory) {
	if ic.Lookup == nil {
		ic.Lookup = new(Lookup)
	}
	ic.Lookup.Forward = make(map[string]int16)
	ic.Lookup.Reverse = make(map[int16]string)

	var pos int16
	for idx, item := range inventory.Items {
		pos = int16(idx)

		ic.Lookup.Forward[item.ID] = pos
		ic.Lookup.Reverse[pos] = item.ID
	}
}

// GetIndexOfItemByID retrieves cached index of the item in your Inventory by its UID in Lookup.Forward
func (ic *InventoryContainer) GetIndexOfItemByID(UID string) *int16 {
	index, ok := ic.Lookup.Forward[UID]
	if !ok {
		log.Println("Item of UID", UID, "does not exist in cache. Returning -1")
		return nil
	}
	return &index
}

type ValidLocation struct {
	MapInfo []int16
	X       int16
	Y       int16
}

func (ic *InventoryContainer) GetValidLocationForItem(height int8, width int8) *ValidLocation {
	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	var itemID string
	var containerMap = ic.Stash.Container.Map
	var stride = int16(ic.Stash.Container.Width)

	length := int16(len(containerMap))

	position := &ValidLocation{
		MapInfo: make([]int16, 0),
	}

	var counter int8
columnLoop:
	for column := int16(0); column < length; column++ {
		if column%stride == 9 && counter != width {
			position.MapInfo = []int16{}
			counter = 0
			continue
		}

		if itemID = containerMap[column]; itemID != "" {
			position.MapInfo = []int16{}
			counter = 0
			continue
		}

		position.MapInfo = append(position.MapInfo, column)

		var coordinate int16
		for row := int16(1); row <= int16(height); row++ {
			coordinate = row*stride + column
			if itemID = containerMap[coordinate]; itemID != "" {
				position.MapInfo = []int16{}
				counter = 0
				continue columnLoop
			}
			position.MapInfo = append(position.MapInfo, coordinate)
		}

		if counter == width {
			position.Y = position.MapInfo[0] / stride
			position.X = position.MapInfo[0] % stride

			return position
		}
		counter++
	}

	log.Println("There are no positions available for the Item, returning nil")
	return nil
}

type Cache struct {
	serverListings []ServerListing
	response       *ResponseCache
	player         map[string]*PlayerCache
}

func HasGetServerListings() []ServerListing {
	if db.cache.serverListings != nil {
		return db.cache.serverListings
	}
	db.cache.serverListings = make([]ServerListing, 0)
	return db.cache.serverListings
}

type ServerListing struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Ping int    `json:"ping,omitempty"`
}

type PlayerCache struct {
	Inventory *InventoryContainer
	Skills    *SkillsCache
	Hideout   *HideoutCache
	Quests    *QuestCache
	Traders   *TraderCache
}

type ResponseCache struct {
	Save            bool             `json:"-"`
	Overwrite       map[string]*int8 `json:"-"`
	Version         string
	CachedResponses map[string]*[]byte
}

func GetCachedResponses() *ResponseCache {
	return db.cache.response
}

func setCachedResponses() {
	cachePath := filepath.Join(coreFilePath, "response.json")
	if !tools.FileExist(cachePath) {
		db.cache.response = new(ResponseCache)
		db.cache.response.CachedResponses = map[string]*[]byte{
			"/client/settings":                            nil,
			"/client/customization":                       nil,
			"/client/items":                               nil,
			"/client/globals":                             nil,
			"/client/locations":                           nil,
			"/client/game/config":                         nil,
			"/client/languages":                           nil,
			"/client/handbook/templates":                  nil,
			"/client/hideout/areas":                       nil,
			"/client/hideout/qte/list":                    nil,
			"/client/hideout/settings":                    nil,
			"/client/hideout/production/recipes":          nil,
			"/client/hideout/production/scavcase/recipes": nil,
		}
		return
	}

	data := tools.GetJSONRawMessage(cachePath)
	if err := json.UnmarshalNoEscape(data, &db.cache.response); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("This needs to be adjusted in the future for modifications")
}

func (rsc *ResponseCache) SaveIfRequired() {
	if rsc.Save {
		if err := rsc.SaveResponseCache(); err != nil {
			log.Println(err)
		}
		rsc.Save = false
	}
}

func (rsc *ResponseCache) SaveResponseCache() error {
	responseCache := filepath.Join(coreFilePath, "response.json")

	if err := tools.WriteToFile(responseCache, rsc); err != nil {
		return fmt.Errorf("response cache not saved: %w", err)
	}
	log.Println("Response Cache saved")
	db.cache.response.Save = false
	return nil
}

func CheckRequestedResponseCache(route string) bool {
	if _, ok := db.cache.response.Overwrite[route]; ok {
		delete(db.cache.response.Overwrite, route)
		return false
	}
	return db.cache.response.CachedResponses[route] != nil
}

func GetRequestedResponseCache(route string) *[]byte {
	return db.cache.response.CachedResponses[route]
}

func SetResponseCacheForRoute(route string, data []byte) {
	db.cache.response.Save = true
	db.cache.response.CachedResponses[route] = &data
}

type SkillsCache struct {
	Common map[string]int8
}

type HideoutCache struct {
	Areas map[int8]int8
}

type TraderCache struct {
	Index      map[string]*AssortIndex
	Assorts    map[string]*Assort
	Insurances map[string]*Insurances
}

type Insurances struct {
	LoyaltyLevel int8
	PriceCoef    int16
	Items        map[string]int32
}

type QuestCache struct {
	Index map[string]int8
}

type InventoryContainer struct {
	Stash  *Stash
	Lookup *Lookup
}

type Lookup struct {
	Forward map[string]int16
	Reverse map[int16]string
}

type Stash struct {
	SlotID    string
	Container Map
}

type Map struct {
	Height  int8
	Width   int8
	Map     []string
	FlatMap map[string]FlatMapLookup
}

type FlatMapLookup struct {
	Width       int8
	Height      int8
	StartX      int16
	EndX        int16
	Coordinates []int16
}
