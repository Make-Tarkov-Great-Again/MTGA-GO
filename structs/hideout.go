package structs

type Hideout struct {
	Areas       []HideoutArea       `json:"areas"`
	Productions []HideoutProduction `json:"production"`
	QTE         []QTE               `json:"qte"`
	ScavCase    []ScavCase          `json:"scavcase"`
	Settings    HideoutSettings     `json:"settings"`
}

// areas
type AreaBonus struct {
	Value      int    `json:"value"`
	Passive    bool   `json:"passive"`
	Production bool   `json:"production"`
	Visible    bool   `json:"visible"`
	Type       string `json:"type"`
	SkillType  string `json:"skillType,omitempty"`
}

type AreaRequirement struct {
	TemplateID   string `json:"templateId"`
	Count        int    `json:"count"`
	IsFunctional bool   `json:"isFunctional"`
	IsEncoded    bool   `json:"isEncoded"`
	Type         string `json:"type"`
}

type AreaImprovement struct {
	ID              string            `json:"id"`
	Requirements    []AreaRequirement `json:"requirements"`
	Bonuses         []AreaBonus       `json:"bonuses"`
	ImprovementTime int               `json:"improvementTime"`
}

type AreaStage struct {
	Requirements     []AreaRequirement `json:"requirements"`
	Bonuses          []AreaBonus       `json:"bonuses"`
	Slots            int               `json:"slots"`
	ConstructionTime int               `json:"constructionTime"`
	Description      string            `json:"description"`
	AutoUpgrade      bool              `json:"autoUpgrade"`
	DisplayInterface bool              `json:"displayInterface"`
	Improvements     []AreaImprovement `json:"improvements"`
}

type HideoutArea struct {
	ID                     string               `json:"_id"`
	Type                   int                  `json:"type"`
	Enabled                bool                 `json:"enabled"`
	NeedsFuel              bool                 `json:"needsFuel"`
	TakeFromSlotLocked     bool                 `json:"takeFromSlotLocked"`
	CraftGivesExp          bool                 `json:"craftGivesExp"`
	DisplayLevel           bool                 `json:"displayLevel"`
	Requirements           []AreaRequirement    `json:"requirements"`
	Stages                 map[string]AreaStage `json:"stages"`
	EnableAreaRequirements bool                 `json:"enableAreaRequirements"`
}

type HideoutProductionRequirement struct {
	TemplateID    string `json:"templateId"`
	Count         int    `json:"count"`
	IsFunctional  bool   `json:"isFunctional"`
	IsEncoded     bool   `json:"isEncoded"`
	Type          string `json:"type"`
	AreaType      int    `json:"areaType,omitempty"`
	RequiredLevel int    `json:"requiredLevel,omitempty"`
}

type HideoutProduction struct {
	ID                           string                         `json:"_id"`
	AreaType                     int                            `json:"areaType"`
	Requirements                 []HideoutProductionRequirement `json:"requirements"`
	ProductionTime               int                            `json:"productionTime"`
	NeedFuelForAllProductionTime bool                           `json:"needFuelForAllProductionTime"`
	Locked                       bool                           `json:"locked"`
	EndProduct                   string                         `json:"endProduct"`
	Continuous                   bool                           `json:"continuous"`
	Count                        int                            `json:"count"`
	ProductionLimitCount         int                            `json:"productionLimitCount"`
	IsEncoded                    bool                           `json:"isEncoded"`
}

// QTE
type QTE struct {
	ID              string           `json:"id"`
	Area            int              `json:"area"`
	AreaLevel       int              `json:"areaLevel"`
	QuickTimeEvents []QuickTimeEvent `json:"quickTimeEvents"`
	Requirements    []QTERequirement `json:"requirements"`
	Results         QTEResult        `json:"results"`
}

type XY struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type QuickTimeEvent struct {
	Type         string  `json:"type"`
	Position     XY      `json:"position"`
	Speed        float32 `json:"speed"`
	SuccessRange XY      `json:"successRange"`
	Key          string  `json:"key"`
	StartDelay   float32 `json:"startDelay"`
	EndDelay     float32 `json:"endDelay"`
}

type QTERequirement struct {
	Energy     int    `json:"energy,omitempty"`
	Hydration  int    `json:"hydration,omitempty"`
	BodyPart   string `json:"bodyPart,omitempty"`
	EffectName string `json:"effectName,omitempty"`
	Excluded   bool   `json:"excluded,omitempty"`
	Type       string `json:"type"`
}

type LevelMultiplier struct {
	Level      int     `json:"level"`
	Multiplier float32 `json:"multiplier"`
}

type RewardRange struct {
	Weight           int               `json:"weight"`
	Result           string            `json:"result"`
	Time             int               `json:"time,omitempty"`
	SkillID          string            `json:"skillId,omitempty"`
	LevelMultipliers []LevelMultiplier `json:"levelMultipliers,omitempty"`
	Type             string            `json:"type,omitempty"`
}

type QTEResult struct {
	FinishEffect struct {
		Energy       int           `json:"energy"`
		Hydration    int           `json:"hydration"`
		RewardsRange []RewardRange `json:"rewardsRange"`
	} `json:"finishEffect"`

	SingleSuccessEffect struct {
		Energy       int           `json:"energy"`
		Hydration    int           `json:"hydration"`
		RewardsRange []RewardRange `json:"rewardsRange"`
	} `json:"singleSuccessEffect"`

	SingleFailEffect struct {
		Energy       int           `json:"energy"`
		Hydration    int           `json:"hydration"`
		RewardsRange []RewardRange `json:"rewardsRange"`
	} `json:"singleFailEffect"`
}

// scavcase
type ScavCaseEndProducts struct {
	Common struct {
		Max string `json:"max"`
		Min string `json:"min"`
	} `json:"Common"`
	Rare struct {
		Max string `json:"max"`
		Min string `json:"min"`
	} `json:"Rare"`
	Superrare struct {
		Max string `json:"max"`
		Min string `json:"min"`
	} `json:"Superrare"`
}

// ScavCaseRequirement represents each requirement in the JSON data
type ScavCaseRequirement struct {
	Count        int    `json:"count"`
	IsEncoded    bool   `json:"isEncoded"`
	IsFunctional bool   `json:"isFunctional"`
	TemplateID   string `json:"templateId"`
	Type         string `json:"type"`
}

// ScavCase represents each scav case in the JSON data
type ScavCase struct {
	ID             string                `json:"_id"`
	EndProducts    ScavCaseEndProducts   `json:"EndProducts"`
	ProductionTime int                   `json:"ProductionTime"`
	Requirements   []ScavCaseRequirement `json:"Requirements"`
}

type HideoutSettings struct {
	GeneratorSpeedWithoutFuel float64 `json:"generatorSpeedWithoutFuel"`
	GeneratorFuelFlowRate     float64 `json:"generatorFuelFlowRate"`
	AirFilterUnitFlowRate     float64 `json:"airFilterUnitFlowRate"`
	GPUBoostRate              float64 `json:"gpuBoostRate"`
}
