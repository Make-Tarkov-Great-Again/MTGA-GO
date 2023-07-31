package structs

type Trader struct {
	Base        Base
	Assort      Assort
	BaseAssort  Assort
	QuestAssort QuestAssort    `omitempty:""`
	Suits       []Suit         `omitempty:""`
	Dialogue    TraderDialogue `omitempty:""`
}

type Base struct {
	ID                  string `json:"_id"`
	AvailableInRaid     bool   `json:"availableInRaid"`
	Avatar              string `json:"avatar"`
	BalanceDol          int32  `json:"balance_dol"`
	BalanceEur          int32  `json:"balance_eur"`
	BalanceRub          int32  `json:"balance_rub"`
	BuyerUp             bool   `json:"buyer_up"`
	Currency            string `json:"currency"`
	CustomizationSeller bool   `json:"customization_seller"`
	Discount            int32  `json:"discount"`
	DiscountEnd         int32  `json:"discount_end"`
	GridHeight          int16  `json:"gridHeight"`
	Insurance           struct {
		Availability     bool     `json:"availability"`
		ExcludedCategory []string `json:"excluded_category"`
		MaxReturnHour    int16    `json:"max_return_hour"`
		MaxStorageTime   int16    `json:"max_storage_time"`
		MinPayment       int32    `json:"min_payment"`
		MinReturnHour    int16    `json:"min_return_hour"`
	} `json:"insurance"`
	ItemsBuy struct {
		Category []string `json:"category"`
		IDList   []string `json:"id_list"`
	} `json:"items_buy"`
	ItemsBuyProhibited struct {
		Category []string `json:"category"`
		IDList   []string `json:"id_list"`
	} `json:"items_buy_prohibited"`
	Location      string `json:"location"`
	LoyaltyLevels []struct {
		BuyPriceCoef       int16   `json:"buy_price_coef"`
		ExchangePriceCoef  int16   `json:"exchange_price_coef"`
		HealPriceCoef      int16   `json:"heal_price_coef"`
		InsurancePriceCoef int16   `json:"insurance_price_coef"`
		MinLevel           int8    `json:"minLevel"`
		MinSalesSum        int32   `json:"minSalesSum"`
		MinStanding        float32 `json:"minStanding"`
		RepairPriceCoef    int16   `json:"repair_price_coef"`
	} `json:"loyaltyLevels"`
	Medic        bool   `json:"medic"`
	Name         string `json:"name"`
	NextResupply int32  `json:"nextResupply"`
	Nickname     string `json:"nickname"`
	Repair       struct {
		Availability        bool     `json:"availability"`
		Currency            string   `json:"currency"`
		CurrencyCoefficient int8     `json:"currency_coefficient"`
		ExcludedCategory    []string `json:"excluded_category"`
		ExcludedIDList      []string `json:"excluded_id_list"`
		PriceRate           int32    `json:"price_rate"`
		Quality             float32  `json:"quality"`
	} `json:"repair"`
	SellCategory      []string `json:"sell_category"`
	Surname           string   `json:"surname"`
	UnlockedByDefault bool     `json:"unlockedByDefault"`
}

type Assort struct {
	Items           []TraderItem                `json:"items"`
	BarterScheme    map[string][][]BarterScheme `json:"barter_scheme"`
	LoyalLevelItems map[string]int8             `json:"loyal_level_items"`
}

type TraderItem struct {
	ID       string  `json:"_id"`
	Tpl      string  `json:"_tpl"`
	Upd      ItemUpd `json:"upd,omitempty"`
	ParentID string  `json:"parentId"`
	SlotID   string  `json:"slotId"`
}

type BarterScheme struct {
	Count float32 `json:"count"`
	Tpl   string  `json:"_tpl"`
}

type QuestAssort struct {
	Started map[string]string `json:"started"`
	Success map[string]string `json:"success"`
	Fail    map[string]string `json:"fail"`
}

type ItemRequirement struct {
	Tpl            string `json:"_tpl"`
	Count          int32  `json:"count"`
	OnlyFunctional bool   `json:"onlyFunctional"`
}

type Requirements struct {
	ItemRequirements  []ItemRequirement `json:"itemRequirements"`
	LoyaltyLevel      int32             `json:"loyaltyLevel"`
	ProfileLevel      int32             `json:"profileLevel"`
	QuestRequirements []string          `json:"questRequirements"`
	SkillRequirements []interface{}     `json:"skillRequirements"`
	Standing          int32             `json:"standing"`
}

type Suit struct {
	ID           string       `json:"_id"`
	IsActive     bool         `json:"isActive"`
	Requirements Requirements `json:"requirements"`
	SuiteID      string       `json:"suiteId"`
	TID          string       `json:"tid"`
}

type TraderDialogue struct {
	InsuranceStart    []string `json:"insuranceStart"`
	InsuranceFound    []string `json:"insuranceFound"`
	InsuranceExpired  []string `json:"insuranceExpired"`
	InsuranceComplete []string `json:"insuranceComplete"`
	InsuranceFailed   []string `json:"insuranceFailed"`
}
