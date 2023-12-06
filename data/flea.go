package data

import (
	"MT-GO/tools"
)

var flea = Flea{
	Offers:           nil,
	OffersCount:      0,
	SelectedCategory: "",
	Categories:       make(map[string]int),
}

type Ragfair struct {
	AllOffers     []Offer
	AllCategories map[string][]string
	Index         map[string]int16
}

// #region Flea getters

func GetFlea() *[]Offer {
	return &db.ragfair.AllOffers
}

func createFleaOffer(userId string, items []AssortItem, scheme []*Scheme) *Offer {
	return nil
}

func GetRequestedFlea() {

}

// #region Flea setters

func setFlea() {
	db.ragfair = &Ragfair{
		AllCategories: make(map[string][]string),
		Index:         make(map[string]int16),
	}

	output := make([]Offer, 0)
	var fleaOffersCount int16

	for tid, trader := range db.trader {
		if trader.Assort == nil {
			continue
		}

		for id, s := range trader.Assort.BarterScheme {
			var scheme []*Scheme
			var items []AssortItem
			var main AssortItem

			if idx, ok := trader.Index.Assort.Items[id]; ok {
				main = *trader.Assort.Items[idx]
				flea.Categories[main.Tpl]++
				scheme = s[0]
				items = []AssortItem{main}
			} else if family, ok := trader.Index.Assort.ParentItems[id]; ok {
				items = make([]AssortItem, 0, len(family))

				scheme = s[0]
				for _, value := range family {
					item := *trader.Assort.Items[value]
					if item.SlotID == "hideout" {
						main = item
						item.SlotID = ""
					}
					flea.Categories[item.Tpl]++
					items = append(items, *trader.Assort.Items[value])
				}
			}

			price, err := GetPriceByID(main.Tpl)
			if err != nil {
				panic(err)
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
				LoyaltyLevel:     trader.Assort.LoyalLevelItems[main.ID],
			}

			if main.Upd.BuyRestrictionMax != 0 {
				offer.BuyRestrictionMax = main.Upd.BuyRestrictionMax
			} else {
				offer.UnlimitedCount = true
			}

			output = append(output, *offer)
			db.ragfair.AllCategories[main.Tpl] = append(db.ragfair.AllCategories[main.Tpl], offer.ID)
			db.ragfair.Index[offer.ID] = fleaOffersCount
			fleaOffersCount++
		}
	}

	db.ragfair.AllOffers = append(make([]Offer, 0, len(output)), output...)
}

// #endregion

// #region Flea structs

type Flea struct {
	Offers           []Offer        `json:"offers"`
	OffersCount      int16          `json:"offersCount"`
	SelectedCategory string         `json:"selectedCategory"` //selected item category
	Categories       map[string]int `json:"categories"`       //categories are the TPL of an offer
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
