package data

var flea = Flea{}

// #region Flea getters

func GetFlea() *Flea {
	return &flea
}

func ConvertTraderOfferToFleaOffer() {

}

// #region Flea setters

func setFlea() {

}

// #endregion

// #region Flea structs

type Flea struct {
	Offers           []any          `json:"offers"`
	OffersCount      int16          `json:"offersCount"`
	SelectedCategory string         `json:"selectedCategory"` //selected item category
	Categories       map[string]int `json:"categories"`       //categories are the TPL of an offer
}
type MemberCategory int

type Offer struct { //nolint:maligned
	ID               string       `json:"_id"`
	IntID            int16        `json:"intId"`
	User             OfferUser    `json:"user"`
	Root             string       `json:"root"`
	Items            []AssortItem `json:"items"`
	ItemsCost        int32        `json:"itemsCost"` // handbook.GetPriceByID()
	Requirements     []Scheme     `json:"requirements"`
	RequirementsCost int32        `json:"requirementsCost"` // Requirements[0].Count, this, SummaryCost are all the same
	SummaryCost      int32        `json:"summaryCost"`
	SellInOnePiece   bool         `json:"sellInOnePiece"`
	StartTime        int32        `json:"startTime"` // current time
	EndTime          int32        `json:"endTime"`   //nextResupply
	UnlimitedCount   bool         `json:"unlimitedCount"`
	LoyaltyLevel     int8         `json:"loyaltyLevel"`
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
