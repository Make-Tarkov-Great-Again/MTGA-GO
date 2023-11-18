package data

import "MT-GO/tools"

var flea = Flea{
	Offers:           nil,
	OffersCount:      0,
	SelectedCategory: "",
	Categories:       make(map[string]int),
}

// #region Flea getters

func GetFlea() *Flea {
	return &flea
}

var traderFleaOffer *Offer

func createFleaOffer(userId string, items []AssortItem, scheme []*Scheme) *Offer {
	return nil
}

// #region Flea setters

// TODO: TraderID, check if items > 1 for trader.Index.Assort.ParentItems
var fleaOffersCount int16

func setFlea() {
	output := make([]Offer, 0)
	for tid, trader := range traders {
		var scheme []*Scheme
		var items []AssortItem
		for _, idx := range trader.Index.Assort.Items {
			item := *trader.Assort.Items[idx]

			if _, ok := flea.Categories[item.Tpl]; !ok {
				flea.Categories[item.Tpl] = 1
			} else {
				flea.Categories[item.Tpl]++
			}

			scheme = trader.Assort.BarterScheme[item.ID][0]
			items = []AssortItem{item}
		}
		for parentId, family := range trader.Index.Assort.ParentItems {
			items = make([]AssortItem, 0, len(family))
			scheme = trader.Assort.BarterScheme[parentId][0]
			for _, value := range family {
				item := *trader.Assort.Items[value]

				if _, ok := flea.Categories[item.Tpl]; !ok {
					flea.Categories[item.Tpl] = 1
				} else {
					flea.Categories[item.Tpl]++
				}
				items = append(items, *trader.Assort.Items[value])
			}

		}

		offer := traderFleaOffer
		offer.ID = tools.GenerateMongoID()
		offer.IntID = fleaOffersCount
		offer.User = OfferUser{
			ID:         tid,
			MemberType: 4,
		}

		offer.SummaryCost = offer.RequirementsCost
		offer.SellInOnePiece = false
		offer.Root = items[0].ID
		offer.LoyaltyLevel = trader.Assort.LoyalLevelItems[offer.Root]
		if items[0].Upd.BuyRestrictionMax != 0 {
			offer.BuyRestrictionMax = items[0].Upd.BuyRestrictionMax
		}

		offer.Items = items
		offer.Requirements = scheme
		offer.RequirementsCost = scheme[0].Count
		offer.StartTime = int32(tools.GetCurrentTimeInSeconds())
		offer.EndTime = int32(trader.Assort.NextResupply)

		output = append(output, *offer)

		fleaOffersCount++
	}

	flea.Offers = output
	flea.OffersCount = fleaOffersCount

	//TODO: Set Trader offers as flea offers
	// Create Flea Index to match to Trader Offers?
	// Cry
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
	ID                string       `json:"_id"`
	IntID             int16        `json:"intId"`
	User              OfferUser    `json:"user"`
	Root              string       `json:"root"`
	Items             []AssortItem `json:"items"`
	ItemsCost         float32      `json:"itemsCost"` // handbook.GetPriceByID()
	Requirements      []*Scheme    `json:"requirements"`
	RequirementsCost  float32      `json:"requirementsCost"` // Requirements[0].Count, this, SummaryCost are all the same
	SummaryCost       float32      `json:"summaryCost"`
	SellInOnePiece    bool         `json:"sellInOnePiece"`
	StartTime         int32        `json:"startTime"` // current time
	EndTime           int32        `json:"endTime"`   //nextResupply
	UnlimitedCount    bool         `json:"unlimitedCount"`
	BuyRestrictionMax int16        `json:"buyRestrictionMax"`
	LoyaltyLevel      int8         `json:"loyaltyLevel"`
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
