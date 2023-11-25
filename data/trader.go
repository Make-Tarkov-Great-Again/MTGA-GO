package data

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var traders = map[string]*Trader{}

// #region Trader getters

func GetTraders() map[string]*Trader {
	return traders
}

// GetTraderByUID returns trader by UID
func GetTraderByUID(UID string) (*Trader, error) {
	trader, ok := traders[UID]
	if ok {
		return trader, nil
	}
	return nil, fmt.Errorf("Trader with UID", UID, "does not exist, returning nil")
}

var tradersByName = map[string]string{
	"Prapor":           "54cb50c76803fa8b248b4571",
	"Therapist":        "54cb57776803fa99248b456e",
	"Fence":            "579dc571d53a0658a154fbec",
	"Skier":            "58330581ace78e27b8b10cee",
	"PeaceKeeper":      "5935c25fb3acc3127c3d8cd9",
	"Mechanic":         "5a7c2eca46aef81a7ca2145d",
	"Ragman":           "5ac3b934156ae10c4430e83c",
	"Jaeger":           "5c0647fdd443bc2504c2d371",
	"LighthouseKeeper": "638f541a29ffd1183d187f57",
}

// GetTraderIDByName returns the TID by their name
//
// Prapor, Therapist, Fence, Skier, PeaceKeeper, Mechanic, Ragman, Jaeger, LighthouseKeeper
func GetTraderIDByName(name string) (*string, error) {
	tid, ok := tradersByName[name]
	if !ok {
		return nil, fmt.Errorf("Trader with name", name, "does not exist, returning nil")
	}
	return &tid, nil
}

// GetTraderByName returns Trader by their name
//
// Prapor, Therapist, Fence, Skier, PeaceKeeper, Mechanic, Ragman, Jaeger, LighthouseKeeper
func GetTraderByName(name string) (*Trader, error) {
	tid, ok := tradersByName[name]
	if !ok {
		return nil, fmt.Errorf("Trader with name", name, "does not exist, returning nil")
	}
	return traders[tid], nil
}

func CloneTrader(name string) *Trader {
	nt := new(Trader)
	tc, err := GetTraderByName(name)
	if err != nil {
		log.Println("Error Cloning Trader %s: %s", tc.Base.ID, err)
		return nil
	}

	TraderJSON, err := json.Marshal(tc)
	if err != nil {
		log.Println("Error Cloning Trader %s: %s", tc.Base.ID, err)
		return nil
	}

	if err := json.Unmarshal(TraderJSON, &nt); err != nil {
		log.Println("Error Cloning Trader %s: %s", tc.Base.ID, err)
		return nil
	}
	return nt

}

// GetAssortItemByID returns entire item from assort as a slice (to get parent item use [0] when calling)
func (t *Trader) GetAssortItemByID(id string) []*AssortItem {
	item, ok := t.Index.Assort.Items[id]
	if ok {
		return []*AssortItem{t.Assort.Items[item]}
	}

	parentItems, parentOK := t.Index.Assort.ParentItems[id]
	if !parentOK {
		log.Println("Assort Item", id, "does not exist for", t.Base.Nickname)
		return nil
	}

	var parent *AssortItem

	items := make([]*AssortItem, 0, len(parentItems))
	for _, index := range parentItems {
		if t.Assort.Items[index].ID == id {
			parent = t.Assort.Items[index]
			continue
		}
		items = append(items, t.Assort.Items[index])
	}

	items = append(items, parent)
	return items
}

func (t *Trader) GetStrippedAssort(character *Character) (*Assort, error) {
	traderID := t.Base.ID

	cache, err := GetTraderCacheByID(character.ID)
	if err != nil {
		return nil, err
	}

	cachedAssort, ok := cache.Assorts[traderID]
	if ok {
		return cachedAssort, nil
	}

	traderInfo, ok := character.TradersInfo[traderID]
	if !ok || traderInfo.LoyaltyLevel == 0 {
		t.SetTraderLoyaltyLevel(character) // check loyalty level
	}
	loyaltyLevel := character.TradersInfo[traderID].LoyaltyLevel

	assortIndex := AssortIndex{
		Items:       map[string]int16{},
		ParentItems: map[string]map[string]int16{},
	}

	assort := Assort{
		NextResupply:    0,
		BarterScheme:    make(map[string][][]*Scheme),
		Items:           make([]*AssortItem, 0, len(t.Assort.Items)),
		LoyalLevelItems: make(map[string]int8),
	}

	// TODO: add quest checks
	for loyalID, loyalLevel := range t.Assort.LoyalLevelItems {
		if loyaltyLevel >= loyalLevel {
			assort.LoyalLevelItems[loyalID] = loyalLevel
			continue

			/* if t.QuestAssort == nil {
				loyalLevelItems[loyalID] = loyalLevel
				continue
			}

			for _, condition := range t.QuestAssort {
				if len(condition) == 0 {
					continue
				}

				for aid, qid := range condition {


				}
			} */
		}
	}

	var counter int16
	for itemID := range assort.LoyalLevelItems {
		index, ok := t.Index.Assort.Items[itemID]
		if ok {
			assort.BarterScheme[itemID] = t.Assort.BarterScheme[itemID]

			assortIndex.Items[itemID] = counter
			counter++
			assort.Items = append(assort.Items, t.Assort.Items[index])
		} else {
			family, ok := t.Index.Assort.ParentItems[itemID]
			if ok {
				assort.BarterScheme[itemID] = t.Assort.BarterScheme[itemID]

				assortIndex.ParentItems[itemID] = make(map[string]int16)
				for k, v := range family {
					assortIndex.ParentItems[itemID][k] = counter
					counter++
					assort.Items = append(assort.Items, t.Assort.Items[v])
				}
			}
		}
	}

	assort.NextResupply = SetResupplyTimer()

	cache.Index[traderID] = &assortIndex
	cache.Assorts[traderID] = &assort

	return cache.Assorts[traderID], nil
}

type ResupplyTimer struct {
	TimerResupplyTime     time.Duration
	ResupplyTimeInSeconds int
	NextResupplyTime      int
	TimerSet              bool
	Profiles              map[string]*Profile
}

var rs = &ResupplyTimer{
	TimerResupplyTime:     0,
	ResupplyTimeInSeconds: 3600, //1 hour
	NextResupplyTime:      0,
	TimerSet:              false,
	Profiles:              nil,
}

func SetResupplyTimer() int {
	if rs.TimerSet {
		return rs.NextResupplyTime
	}

	//TODO: Adjust rs.ResupplyTimeInSeconds based on a config

	rs.NextResupplyTime = int(tools.GetCurrentTimeInSeconds()) + rs.ResupplyTimeInSeconds
	rs.TimerResupplyTime = time.Duration(rs.ResupplyTimeInSeconds) * time.Second

	rs.TimerSet = true

	go func() {
		timer := time.NewTimer(rs.TimerResupplyTime)
		for {
			<-timer.C
			rs.NextResupplyTime += rs.ResupplyTimeInSeconds
			rs.Profiles = GetProfiles()

			for _, profile := range rs.Profiles {
				traders, err := GetTraderCacheByID(profile.Character.ID)
				if err != nil {
					continue
				}

				for _, assort := range traders.Assorts {
					assort.NextResupply = rs.NextResupplyTime
				}
			}

			timer.Reset(rs.TimerResupplyTime)
		}
	}()

	return rs.NextResupplyTime
}

// SetTraderLoyaltyLevel determines the loyalty level of a trader based on character attributes
func (t *Trader) SetTraderLoyaltyLevel(character *Character) {
	loyaltyLevels := t.Base.LoyaltyLevels
	traderID := t.Base.ID

	traderInfo, ok := character.TradersInfo[traderID]
	if !ok {
		return
	}

	length := int8(len(loyaltyLevels))
	for idx := traderInfo.LoyaltyLevel; idx < length; idx++ {
		loyalty := loyaltyLevels[idx]

		if character.Info.Level < loyalty.MinLevel ||
			character.TradersInfo[traderID].SalesSum < loyalty.MinSalesSum ||
			character.TradersInfo[traderID].Standing < loyalty.MinStanding {

			traderInfo.LoyaltyLevel = idx
			character.TradersInfo[traderID] = traderInfo
			return
		}
	}
}

/*
GetItemFamilyTree returns the family of an item based on parentID if it and the family exists
*/
func GetItemFamilyTree(items []*AssortItem, parent string) []string {
	var list []string

	for _, childItem := range items {
		child := childItem

		if child.ParentID == parent {
			list = append(list, GetItemFamilyTree(items, child.ID)...)
		}
	}

	list = append(list, parent) // required
	return list
}

func setTraders() {
	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		log.Println(err)
		return
	}

	for dir := range directory {
		count := 0
		done := make(chan bool)
		trader := new(Trader)
		currentTraderPath := filepath.Join(traderPath, dir)

		basePath := filepath.Join(currentTraderPath, "base.json")
		if tools.FileExist(basePath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(basePath)
				trader.Base = new(TraderBase)
				if err := json.UnmarshalNoEscape(raw, &trader.Base); err != nil {
					log.Fatalln(err)
				}
				done <- true
			}()
		}

		assortPath := filepath.Join(currentTraderPath, "assort.json")
		if tools.FileExist(assortPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(assortPath)
				trader.Assort = new(Assort)
				if err := json.UnmarshalNoEscape(raw, &trader.Assort); err != nil {
					log.Fatalln(err)
				}
				trader.Assort.NextResupply = 0

				done <- true
			}()
		}

		questsPath := filepath.Join(currentTraderPath, "questassort.json")
		if tools.FileExist(questsPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(questsPath)
				trader.QuestAssort = make(map[string]map[string]string)
				if err := json.UnmarshalNoEscape(raw, &trader.QuestAssort); err != nil {
					log.Fatalln(err)
				}
				done <- true
			}()
		}

		suitsPath := filepath.Join(currentTraderPath, "suits.json")
		if tools.FileExist(suitsPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(suitsPath)
				trader.Suits = make([]TraderSuits, 0)
				if err := json.UnmarshalNoEscape(raw, &trader.Suits); err != nil {
					log.Println(err)
				}
				done <- true
			}()
		}

		dialoguesPath := filepath.Join(currentTraderPath, "dialogue.json")
		if tools.FileExist(dialoguesPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(dialoguesPath)
				trader.Dialogue = make(map[string][]string)
				if err := json.UnmarshalNoEscape(raw, &trader.Dialogue); err != nil {
					log.Fatalln(err)
				}
				done <- true
			}()
		}

		for i := 0; i < count; i++ {
			<-done
		}
		traders[dir] = trader
	}
}

func SetTraderOfferLookup() {
	for _, trader := range traders {
		if trader.Assort != nil {
			trader.Index.Assort = &AssortIndex{}
			parentItems := make(map[string]map[string]int16)
			childlessItems := make(map[string]int16)

			for index, item := range trader.Assort.Items {
				itemChildren := GetItemFamilyTree(trader.Assort.Items, item.ID)
				if len(itemChildren) == 1 {
					childlessItems[item.ID] = int16(index)
					continue
				}

				family := make(map[string]int16)
				for _, child := range itemChildren {
					for k, v := range trader.Assort.Items {
						if child != v.ID {
							continue
						}

						family[child] = int16(k)
						break
					}
				}
				parentItems[item.ID] = family
			}

			trader.Index.Assort.ParentItems = parentItems
			trader.Index.Assort.Items = childlessItems
		}

		if trader.Suits != nil {
			trader.Index.Suits = make(map[string]int8)
			for index, suit := range trader.Suits {
				trader.Index.Suits[suit.SuiteID] = int8(index)
			}
		}
	}
}

type Trader struct {
	Index       TraderIndex                  `json:",omitempty"`
	Base        *TraderBase                  `json:",omitempty"`
	Assort      *Assort                      `json:",omitempty"`
	QuestAssort map[string]map[string]string `json:",omitempty"`
	Suits       []TraderSuits                `json:",omitempty"`
	Dialogue    map[string][]string          `json:",omitempty"`
}

type TraderIndex struct {
	Assort *AssortIndex    `json:",omitempty"`
	Suits  map[string]int8 `json:",omitempty"`
}

type AssortIndex struct {
	Items       map[string]int16
	ParentItems map[string]map[string]int16 `json:",omitempty"`
}

type TraderBase struct {
	ID                  string               `json:"_id"`
	AvailableInRaid     bool                 `json:"availableInRaid"`
	Avatar              string               `json:"avatar"`
	BalanceDol          int32                `json:"balance_dol"`
	BalanceEur          int32                `json:"balance_eur"`
	BalanceRub          int32                `json:"balance_rub"`
	BuyerUp             bool                 `json:"buyer_up"`
	Currency            string               `json:"currency"`
	CustomizationSeller bool                 `json:"customization_seller"`
	Discount            int8                 `json:"discount"`
	DiscountEnd         int8                 `json:"discount_end"`
	GridHeight          int16                `json:"gridHeight"`
	Insurance           TraderInsurance      `json:"insurance"`
	ItemsBuy            ItemsBuy             `json:"items_buy"`
	ItemsBuyProhibited  ItemsBuy             `json:"items_buy_prohibited"`
	Location            string               `json:"location"`
	LoyaltyLevels       []TraderLoyaltyLevel `json:"loyaltyLevels"`
	Medic               bool                 `json:"medic"`
	Name                string               `json:"name"`
	NextResupply        int32                `json:"nextResupply"`
	Nickname            string               `json:"nickname"`
	Repair              TraderRepair         `json:"repair"`
	SellCategory        []string             `json:"sell_category"`
	Surname             string               `json:"surname"`
	UnlockedByDefault   bool                 `json:"unlockedByDefault"`
}

type TraderInsurance struct {
	Availability     bool     `json:"availability"`
	ExcludedCategory []string `json:"excluded_category"`
	MaxReturnHour    int8     `json:"max_return_hour"`
	MaxStorageTime   int32    `json:"max_storage_time"`
	MinPayment       float32  `json:"min_payment"`
	MinReturnHour    int8     `json:"min_return_hour"`
}

type ItemsBuy struct {
	Category []string `json:"category"`
	IdList   []string `json:"id_list"`
}

type TraderLoyaltyLevel struct {
	BuyPriceCoef       int16   `json:"buy_price_coef"`
	ExchangePriceCoef  int16   `json:"exchange_price_coef"`
	HealPriceCoef      int16   `json:"heal_price_coef"`
	InsurancePriceCoef int16   `json:"insurance_price_coef"`
	MinLevel           int8    `json:"minLevel"`
	MinSalesSum        float32 `json:"minSalesSum"`
	MinStanding        float32 `json:"minStanding"`
	RepairPriceCoef    int16   `json:"repair_price_coef"`
}

type TraderRepair struct {
	Availability        bool     `json:"availability"`
	Currency            string   `json:"currency"`
	CurrencyCoefficient int8     `json:"currency_coefficient"`
	ExcludedCategory    []string `json:"excluded_category"`
	ExcludedIdList      []string `json:"excluded_id_list"`
	PriceRate           int8     `json:"price_rate"`
	Quality             float32  `json:"quality"`
}

type TraderSuits struct {
	ID           string           `json:"_id"`
	Tid          string           `json:"tid"`
	SuiteID      string           `json:"suiteId"`
	IsActive     bool             `json:"isActive"`
	Requirements SuitRequirements `json:"requirements"`
}

type SuitItemRequirements struct {
	Count          int    `json:"count"`
	Tpl            string `json:"_tpl"`
	OnlyFunctional bool   `json:"onlyFunctional"`
}

type SuitRequirements struct {
	LoyaltyLevel         int8                   `json:"loyaltyLevel"`
	ProfileLevel         int8                   `json:"profileLevel"`
	Standing             float32                `json:"standing"`
	SkillRequirements    []map[string]int8      `json:"skillRequirements"`
	QuestRequirements    []string               `json:"questRequirements"`
	SuitItemRequirements []SuitItemRequirements `json:"itemRequirements"`
}

type Assort struct {
	NextResupply    int                    `json:"nextResupply"`
	BarterScheme    map[string][][]*Scheme `json:"barter_scheme"`
	Items           []*AssortItem          `json:"items"`
	LoyalLevelItems map[string]int8        `json:"loyal_level_items"`
}

type AssortItem struct {
	ID       string      `json:"_id"`
	Tpl      string      `json:"_tpl"`
	ParentID string      `json:"parentId"`
	SlotID   string      `json:"slotId,omitempty"`
	Upd      *ItemUpdate `json:"upd,omitempty"`
}

type Scheme struct {
	Tpl   string  `json:"_tpl"`
	Count float32 `json:"count"`
}

// #endregion
