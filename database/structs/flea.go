package structs

type Flea struct {
	Offers           []FleaOffer
	OffersCount      int
	SelectedCategory string
	Categories       map[string]int
}

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

type FleaOffer struct {
	SellResult         []SellResult       `json:"sellResult,omitempty"`
	ID                 string             `json:"_id"`
	Items              []InventoryItem    `json:"items"`
	Requirements       []OfferRequirement `json:"requirements"`
	Root               string             `json:"root"`
	IntID              int                `json:"intId"`
	ItemsCost          int                `json:"itemsCost"`
	RequirementsCost   int                `json:"requirementsCost"`
	StartTime          int                `json:"startTime"`
	EndTime            int                `json:"endTime"`
	SellInOnePiece     bool               `json:"sellInOnePiece"`
	LoyaltyLevel       int                `json:"loyaltyLevel"`
	BuyRestrictionMax  int                `json:"buyRestrictionMax,omitempty"`
	BuyRestrictionCurr int                `json:"buyRestrictionCurrent,omitempty"`
	Locked             bool               `json:"locked"`
	UnlimitedCount     bool               `json:"unlimitedCount"`
	SummaryCost        int                `json:"summaryCost"`
	User               IRagfairOfferUser  `json:"user"`
	NotAvailable       bool               `json:"notAvailable"`
	CurrentItemCount   int                `json:"CurrentItemCount"`
	Priority           bool               `json:"priority"`
}

type OfferRequirement struct {
	Tpl            string `json:"_tpl"`
	Count          int    `json:"count"`
	OnlyFunctional bool   `json:"onlyFunctional"`
}

type IRagfairOfferUser struct {
	ID              string         `json:"id"`
	Nickname        string         `json:"nickname,omitempty"`
	Rating          int            `json:"rating,omitempty"`
	MemberType      MemberCategory `json:"memberType"`
	Avatar          string         `json:"avatar,omitempty"`
	IsRatingGrowing bool           `json:"isRatingGrowing,omitempty"`
}

type SellResult struct {
	SellTime int64 `json:"sellTime"`
	Amount   int   `json:"amount"`
}

type MemberCategory int
