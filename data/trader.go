package data

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"log"
	"path/filepath"
	"time"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Trader getters

func GetTraders() *haxmap.Map[string, *Trader] {
	return db.trader.Traders
}

// GetTraderByUID returns trader by UID
func GetTraderByUID(UID string) (*Trader, error) {
	trader, ok := db.trader.Traders.Get(UID)
	if ok {
		return trader, nil
	}
	return nil, fmt.Errorf("trader %s does not exist, returning nil", UID)
}

// GetTraderIDByName returns the TID by their name
//
// Prapor, Therapist, Fence, Skier, PeaceKeeper, Mechanic, Ragman, Jaeger, LighthouseKeeper
func GetTraderIDByName(name string) (*string, error) {
	tid, ok := db.trader.Names.Get(name)
	if !ok {
		return nil, fmt.Errorf("trader %s does not exist, returning nil", name)
	}
	return &tid, nil
}

// GetTraderByName returns Trader by their name
//
// Prapor, Therapist, Fence, Skier, PeaceKeeper, Mechanic, Ragman, Jaeger, LighthouseKeeper
func GetTraderByName(name string) (*Trader, error) {
	tid, ok := db.trader.Names.Get(name)
	if !ok {
		return nil, fmt.Errorf("trader with name %s does not exist, returning nil", name)
	}

	trader, ok := db.trader.Traders.Get(tid)
	if !ok {
		return nil, fmt.Errorf("trader %s does not exist, returning nil", name)
	}

	return trader, nil
}

func CloneTrader(name string) *Trader {
	nt := new(Trader)
	tc, err := GetTraderByName(name)
	if err != nil {
		log.Printf("Error Cloning Trader %s: %s\n", name, err)
		return nil
	}

	TraderJSON, err := json.Marshal(tc)
	if err != nil {
		log.Printf("Error Cloning Trader %s: %s\n", name, err)
		return nil
	}

	if err := json.Unmarshal(TraderJSON, &nt); err != nil {
		log.Printf("Error Cloning Trader %s: %s\n", tc.Base.ID, err)
		return nil
	}
	return nt

}

// GetAssortItemByID returns entire item from assort as a slice (to get parent item use [0] when calling)
func (t *Trader) GetAssortItemByID(id string) []*AssortItem {
	item, ok := t.Index.Assort.Items.Get(id)
	if ok {
		return []*AssortItem{t.Assort.Items[item]}
	}

	parentItems, parentOK := t.Index.Assort.ParentItems.Get(id)
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

func (t *Trader) GetIndexOfItem(id string) int16 {
	return 0
}

func (t *Trader) GetStrippedAssort(character *Character[map[string]PlayerTradersInfo]) (*Assort, error) {
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
		Items:       haxmap.New[string, int16](),            //map[string]int16{},
		ParentItems: haxmap.New[string, map[string]int16](), //map[string]map[string]int16{},
	}

	assort := Assort{
		NextResupply:    0,
		BarterScheme:    haxmap.New[string, [][]*Scheme](), //make(map[string][][]*Scheme),
		Items:           make([]*AssortItem, 0, len(t.Assort.Items)),
		LoyalLevelItems: haxmap.New[string, int8](),
	}

	// TODO: add quest checks

	var counter int16
	t.Assort.LoyalLevelItems.ForEach(func(loyalID string, loyalLevel int8) bool {
		if loyaltyLevel >= loyalLevel {
			assort.LoyalLevelItems.Set(loyalID, loyalLevel)

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
		return true
	})

	assort.LoyalLevelItems.ForEach(func(loyalID string, _ int8) bool {
		if index, ok := t.Index.Assort.Items.Get(loyalID); ok {
			barterScheme, ok := t.Assort.BarterScheme.Get(loyalID)
			if !ok {
				log.Fatal("ya momma")
			}

			assort.BarterScheme.Set(loyalID, barterScheme)
			assortIndex.Items.Set(loyalID, counter)
			counter++
			assort.Items = append(assort.Items, t.Assort.Items[index])
		} else if family, ok := t.Index.Assort.ParentItems.Get(loyalID); ok {
			barterScheme, ok := t.Assort.BarterScheme.Get(loyalID)
			if !ok {
				log.Fatal("ya momma")
			}
			assort.BarterScheme.Set(loyalID, barterScheme)
			assortIndex.ParentItems.Set(loyalID, make(map[string]int16))
			parentItems, _ := assortIndex.ParentItems.Get(loyalID)
			for k, v := range family {
				parentItems[k] = counter
				counter++
				assort.Items = append(assort.Items, t.Assort.Items[v])
			}
		}
		return true
	})

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
	Profiles              []string
}

var rs = &ResupplyTimer{
	TimerResupplyTime:     0,
	ResupplyTimeInSeconds: 3600, //1 hour
	NextResupplyTime:      0,
	TimerSet:              false,
	Profiles:              make([]string, 0),
}

func SetResupplyTimer() int {
	if rs.TimerSet {
		return rs.NextResupplyTime
	}

	rs.NextResupplyTime = int(tools.GetCurrentTimeInSeconds()) + rs.ResupplyTimeInSeconds
	rs.TimerResupplyTime = time.Duration(rs.ResupplyTimeInSeconds) * time.Second

	profiles := GetProfiles()
	profiles.ForEach(func(key string, value *Profile) bool {
		rs.Profiles = append(rs.Profiles, value.Account.UID)
		return true
	})

	rs.TimerSet = true

	go func() {
		timer := time.NewTimer(rs.TimerResupplyTime)
		for {
			<-timer.C
			rs.NextResupplyTime += rs.ResupplyTimeInSeconds

			for _, pid := range rs.Profiles {
				traders, err := GetTraderCacheByID(pid)
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
func (t *Trader) SetTraderLoyaltyLevel(character *Character[map[string]PlayerTradersInfo]) {
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
	traderInfo.LoyaltyLevel = length
	character.TradersInfo[traderID] = traderInfo
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

func GetSupplyData() *SupplyData {
	supplyData := db.trader.LogisticData
	if supplyData == nil {
		prices := GetPrices()

		rub, _ := prices.Get(*GetCurrencyByName("RUB"))
		eur, _ := prices.Get(*GetCurrencyByName("EUR"))
		usd, _ := prices.Get(*GetCurrencyByName("USD"))

		db.trader.LogisticData = &SupplyData{
			SupplyNextTime: 0,
			Prices:         make(map[string]int32),
			CurrencyCourses: CurrencyCourses{
				RUB: rub,
				EUR: eur,
				DOL: usd,
			},
		}

		prices.ForEach(func(key string, value int32) bool {
			db.trader.LogisticData.Prices[key] = value
			return true
		})

		return db.trader.LogisticData
	}
	return supplyData
}

func setTraders() {
	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		log.Println(err)
		return
	}
	db.trader = &Traders{
		LogisticData: nil,
		Names:        haxmap.New[string, string](),
		Traders:      haxmap.New[string, *Trader](uintptr(len(directory))),
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
					msg := tools.CheckParsingError(raw, err)
					log.Fatalln(msg)
				}

				db.trader.Names.Set(trader.Base.Nickname, trader.Base.ID)
				done <- true
			}()
		}

		assortPath := filepath.Join(currentTraderPath, "assort.json")
		if tools.FileExist(assortPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(assortPath)
				trader.Assort = &Assort{
					NextResupply:    0,
					BarterScheme:    haxmap.New[string, [][]*Scheme](),
					Items:           make([]*AssortItem, 0),
					LoyalLevelItems: haxmap.New[string, int8](),
				}
				if err := json.Unmarshal(raw, &trader.Assort); err != nil {
					msg := tools.CheckParsingError(raw, err)
					log.Fatalln(msg)
				}

				done <- true
			}()
		}

		questsPath := filepath.Join(currentTraderPath, "questassort.json")
		if tools.FileExist(questsPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(questsPath)
				trader.QuestAssort = haxmap.New[string, map[string]string]() //make(map[string]map[string]string)
				if err := json.UnmarshalNoEscape(raw, &trader.QuestAssort); err != nil {
					msg := tools.CheckParsingError(raw, err)
					log.Fatalln(msg)
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
					msg := tools.CheckParsingError(raw, err)
					log.Fatalln(msg)
				}
				done <- true
			}()
		}

		dialoguesPath := filepath.Join(currentTraderPath, "dialogue.json")
		if tools.FileExist(dialoguesPath) {
			count++
			go func() {
				raw := tools.GetJSONRawMessage(dialoguesPath)
				trader.Dialogue = haxmap.New[string, []string]() //make(map[string][]string)
				if err := json.UnmarshalNoEscape(raw, &trader.Dialogue); err != nil {
					msg := tools.CheckParsingError(raw, err)
					log.Fatalln(msg)
				}
				done <- true
			}()
		}

		for i := 0; i < count; i++ {
			<-done
		}
		db.trader.Traders.Set(dir, trader)
	}
}

func setTraderOfferLookup() {
	db.trader.Traders.ForEach(func(_ string, trader *Trader) bool {
		if trader.Assort != nil && len(trader.Assort.Items) != 0 {
			trader.Index.Assort = &AssortIndex{
				Items:       haxmap.New[string, int16](),            //make(map[string]int16),
				ParentItems: haxmap.New[string, map[string]int16](), //make(map[string]map[string]int16),
			}

			for index, item := range trader.Assort.Items {
				itemChildren := GetItemFamilyTree(trader.Assort.Items, item.ID)
				if len(itemChildren) == 1 {
					trader.Index.Assort.Items.Set(item.ID, int16(index))
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
				trader.Index.Assort.ParentItems.Set(item.ID, family)
			}
		}

		if trader.Suits != nil {
			trader.Index.Suits = make(map[string]int8)
			for index, suit := range trader.Suits {
				trader.Index.Suits[suit.ID] = int8(index)
			}
		}
		return true
	})
}

type Traders struct {
	LogisticData *SupplyData
	Names        *haxmap.Map[string, string]
	Traders      *haxmap.Map[string, *Trader]
}

type SupplyData struct {
	SupplyNextTime  int              `json:"supplyNextTime"`
	Prices          map[string]int32 `json:"prices"`
	CurrencyCourses CurrencyCourses  `json:"currencyCourses"`
}

type CurrencyCourses struct {
	RUB int32 `json:"5449016a4bdc2d6f028b456f"`
	EUR int32 `json:"569668774bdc2da2298b4568"`
	DOL int32 `json:"5696686a4bdc2da3298b456a"`
}

type Trader struct {
	Index       TraderIndex                            `json:",omitempty"`
	Base        *TraderBase                            `json:",omitempty"`
	Assort      *Assort                                `json:",omitempty"`
	QuestAssort *haxmap.Map[string, map[string]string] `json:",omitempty"` //map[string]map[string]string `json:",omitempty"`
	Suits       []TraderSuits                          `json:",omitempty"`
	Dialogue    *haxmap.Map[string, []string]          `json:",omitempty"` //map[string][]string          `json:",omitempty"`
}

type TraderIndex struct {
	Assort *AssortIndex    `json:",omitempty"`
	Suits  map[string]int8 `json:",omitempty"`
}

type AssortIndex struct {
	Items       *haxmap.Map[string, int16]            //map[string]int16
	ParentItems *haxmap.Map[string, map[string]int16] `json:",omitempty"` //map[string]map[string]int16
}

type TraderBase struct {
	ID                             string               `json:"_id"`
	AvailableInRaid                bool                 `json:"availableInRaid"`
	Avatar                         string               `json:"avatar"`
	BalanceDol                     int32                `json:"balance_dol"`
	BalanceEur                     int32                `json:"balance_eur"`
	BalanceRub                     int32                `json:"balance_rub"`
	BuyerUp                        bool                 `json:"buyer_up"`
	Currency                       string               `json:"currency"`
	CustomizationSeller            bool                 `json:"customization_seller"`
	Discount                       int8                 `json:"discount"`
	DiscountEnd                    int8                 `json:"discount_end"`
	GridHeight                     int16                `json:"gridHeight"`
	Insurance                      TraderInsurance      `json:"insurance"`
	ItemsBuy                       ItemsBuy             `json:"items_buy"`
	ItemsBuyProhibited             ItemsBuy             `json:"items_buy_prohibited"`
	Location                       string               `json:"location"`
	LoyaltyLevels                  []TraderLoyaltyLevel `json:"loyaltyLevels"`
	Medic                          bool                 `json:"medic"`
	Name                           string               `json:"name"`
	NextResupply                   int32                `json:"nextResupply"`
	Nickname                       string               `json:"nickname"`
	Repair                         TraderRepair         `json:"repair"`
	SellCategory                   []string             `json:"sell_category"`
	Surname                        string               `json:"surname"`
	UnlockedByDefault              bool                 `json:"unlockedByDefault"`
	SellModifierForProhibitedItems int8                 `json:"sell_modifier_for_prohibited_items"`
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
	NextResupply    int                              `json:"nextResupply"`
	BarterScheme    *haxmap.Map[string, [][]*Scheme] `json:"barter_scheme"` //map[string][][]*Scheme
	Items           []*AssortItem                    `json:"items"`
	LoyalLevelItems *haxmap.Map[string, int8]        `json:"loyal_level_items"` //map[string]int8
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
