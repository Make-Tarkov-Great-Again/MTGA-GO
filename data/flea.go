package data

import (
	"MT-GO/tools"
	"fmt"
	"log"
)

type Ragfair struct {
	Catalog map[string][]Offer
	Market  Flea
}

// #region Flea getters
func GetFlea() *Ragfair {
	return db.ragfair
}

func GetFleaCatalog(id string) ([]Offer, error) {
	catalog, ok := db.ragfair.Catalog[id]
	if !ok {
		return catalog, fmt.Errorf("catalog of %s does not exist", id)
	}
	return catalog, nil
}

// #region Flea setters

func SetFlea() {
	db.ragfair = &Ragfair{
		Catalog: make(map[string][]Offer),
		Market: Flea{
			Offers:           nil,
			OffersCount:      0,
			SelectedCategory: "",
			Categories:       make(map[string]int16),
		},
	}
	var fleaOffersCount int16

	db.trader.ForEach(func(tid string, trader *Trader) bool {
		if trader.Assort == nil {
			return true
		}
		trader.Assort.BarterScheme.ForEach(func(id string, s [][]*Scheme) bool {
			var scheme []*Scheme
			var items []AssortItem
			var main AssortItem

			if idx, ok := trader.Index.Assort.Items.Get(id); ok {
				main = *trader.Assort.Items[idx]
				scheme = s[0]
				items = []AssortItem{main}
			} else if family, ok := trader.Index.Assort.ParentItems.Get(id); ok {
				items = make([]AssortItem, 0, len(family))

				scheme = s[0]
				for _, value := range family {
					item := *trader.Assort.Items[value]
					if item.SlotID == "hideout" {
						main = item
						item.SlotID = ""
					}
					items = append(items, *trader.Assort.Items[value])
				}
			}

			price, err := GetPriceByID(main.Tpl)
			if err != nil {
				panic(err)
			}

			loyalItem, ok := trader.Assort.LoyalLevelItems.Get(main.ID)
			if !ok {
				log.Fatal("loyalitem doesn't exist")
			}

			offer := &Offer{
				ID:    tools.GenerateMongoID(),
				IntID: fleaOffersCount,
				User: OfferUser{
					ID:         tid,
					MemberType: 4,
				},
				Root:             main.ID,
				Items:            items,
				ItemsCost:        price,
				Requirements:     scheme,
				RequirementsCost: int32(scheme[0].Count),
				SummaryCost:      int32(scheme[0].Count),
				SellInOnePiece:   false,
				StartTime:        int32(tools.GetCurrentTimeInSeconds()),
				EndTime:          int32(trader.Assort.NextResupply),
				UnlimitedCount:   false,
				LoyaltyLevel:     loyalItem,
			}

			if main.Upd.BuyRestrictionMax != 0 {
				offer.BuyRestrictionMax = main.Upd.BuyRestrictionMax
			} else {
				offer.UnlimitedCount = true
			}

			if db.ragfair.Catalog[main.Tpl] == nil {
				db.ragfair.Catalog[main.Tpl] = make([]Offer, 0)
			}
			db.ragfair.Catalog[main.Tpl] = append(db.ragfair.Catalog[main.Tpl], *offer)
			db.ragfair.Market.Categories[main.Tpl]++
			return true
		})
		fleaOffersCount++
		return true
	})
}

// #endregion

// #region Flea structs

type Flea struct {
	Offers           []Offer          `json:"offers"`
	OffersCount      int16            `json:"offersCount"`
	SelectedCategory string           `json:"selectedCategory"` //selected item category
	Categories       map[string]int16 `json:"categories"`       //categories are the TPL of an offer
}
type MemberCategory int

type Offer struct { //nolint:maligned
	ID                    string       `json:"_id"`
	IntID                 int16        `json:"intId"`
	User                  OfferUser    `json:"user"`
	Root                  string       `json:"root"`
	Items                 []AssortItem `json:"items"`
	ItemsCost             int32        `json:"itemsCost"` // handbook.GetPriceByID()
	Requirements          []*Scheme    `json:"requirements"`
	RequirementsCost      int32        `json:"requirementsCost"` // Requirements[0].Count, this, SummaryCost are all the same
	SummaryCost           int32        `json:"summaryCost"`
	SellInOnePiece        bool         `json:"sellInOnePiece"`
	StartTime             int32        `json:"startTime"` // current time
	EndTime               int32        `json:"endTime"`   //nextResupply
	UnlimitedCount        bool         `json:"unlimitedCount"`
	BuyRestrictionMax     int16        `json:"buyRestrictionMax"`
	BuyRestrictionCurrent int16        `json:"-"`
	LoyaltyLevel          int8         `json:"loyaltyLevel"`
}

type RagfairFind struct {
	Page              int8           `json:"page"`
	Limit             int8           `json:"limit"`
	SortType          int8           `json:"sortType"`
	SortDirection     int8           `json:"sortDirection"`
	Currency          int8           `json:"currency"`
	PriceFrom         int32          `json:"priceFrom"`
	PriceTo           int32          `json:"priceTo"`
	QuantityFrom      int32          `json:"quantityFrom"`
	QuantityTo        int32          `json:"quantityTo"`
	ConditionFrom     int8           `json:"conditionFrom"`
	ConditionTo       int8           `json:"conditionTo"`
	OneHourExpiration bool           `json:"oneHourExpiration"`
	RemoveBartering   bool           `json:"removeBartering"`
	OfferOwnerType    int8           `json:"offerOwnerType"`
	OnlyFunctional    bool           `json:"onlyFunctional"`
	UpdateOfferCount  bool           `json:"updateOfferCount"`
	HandbookID        string         `json:"handbookId"`
	LinkedSearchID    string         `json:"linkedSearchId"`
	NeededSearchID    string         `json:"neededSearchId"`
	BuildItems        map[string]any `json:"buildItems"`
	BuildCount        int16          `json:"buildCount"`
	Tm                int8           `json:"tm"`
	Reload            int8           `json:"reload"`
}

type OfferUser struct {
	ID         string         `json:"id"`
	MemberType MemberCategory `json:"memberType"`
}

// #endregion

const (
	defaultCategory                  MemberCategory = 0
	developerCategory                MemberCategory = 1
	uniqueIDCategory                 MemberCategory = 2
	traderCategory                   MemberCategory = 4
	groupCategory                    MemberCategory = 8
	systemCategory                   MemberCategory = 16
	chatModeratorCategory            MemberCategory = 32
	chatModeratorWithPermBanCategory MemberCategory = 64
	unitTestCategory                 MemberCategory = 128
	sherpaCategory                   MemberCategory = 256
	emissaryCategory                 MemberCategory = 512
)
