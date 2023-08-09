package structs

type DatabaseItem struct {
	ID         string             `json:"_id,omitempty"`
	Name       string             `json:"_name,omitempty"`
	Parent     string             `json:"_parent,omitempty"`
	Type       string             `json:"_type,omitempty"`
	Properties DatabaseProperties `json:"_props,omitempty"`
	Proto      string             `json:"_proto,omitempty"`
}

type DatabaseProperties struct {
	Name                                   string     `json:"Name,omitempty"`
	ShortName                              string     `json:"ShortName,omitempty"`
	Description                            string     `json:"Description,omitempty"`
	Weight                                 float32    `json:"Weight,omitempty"`
	BackgroundColor                        string     `json:"BackgroundColor,omitempty"`
	Width                                  int        `json:"Width,omitempty"`
	Height                                 int        `json:"Height,omitempty"`
	StackMaxSize                           int        `json:"StackMaxSize,omitempty"`
	ItemSound                              string     `json:"ItemSound,omitempty"`
	Prefab                                 Prefab     `json:"Prefab,omitempty"`
	UsePrefab                              Prefab     `json:"UsePrefab,omitempty"`
	StackObjectsCount                      int        `json:"StackObjectsCount,omitempty"`
	NotShownInSlot                         bool       `json:"NotShownInSlot,omitempty"`
	ExaminedByDefault                      bool       `json:"ExaminedByDefault,omitempty"`
	ExamineTime                            int        `json:"ExamineTime,omitempty"`
	IsUndiscardable                        bool       `json:"IsUndiscardable,omitempty"`
	IsUnsaleable                           bool       `json:"IsUnsaleable,omitempty"`
	IsUnbuyable                            bool       `json:"IsUnbuyable,omitempty"`
	IsUngivable                            bool       `json:"IsUngivable,omitempty"`
	IsLockedafterEquip                     bool       `json:"IsLockedafterEquip,omitempty"`
	QuestItem                              bool       `json:"QuestItem,omitempty"`
	LootExperience                         int        `json:"LootExperience,omitempty"`
	ExamineExperience                      int        `json:"ExamineExperience,omitempty"`
	HideEntrails                           bool       `json:"HideEntrails,omitempty"`
	RepairCost                             int        `json:"RepairCost,omitempty"`
	RepairSpeed                            int        `json:"RepairSpeed,omitempty"`
	ExtraSizeLeft                          int        `json:"ExtraSizeLeft,omitempty"`
	ExtraSizeRight                         int        `json:"ExtraSizeRight,omitempty"`
	ExtraSizeUp                            int        `json:"ExtraSizeUp,omitempty"`
	ExtraSizeDown                          int        `json:"ExtraSizeDown,omitempty"`
	ExtraSizeForceAdd                      bool       `json:"ExtraSizeForceAdd,omitempty"`
	MergesWithChildren                     bool       `json:"MergesWithChildren,omitempty"`
	CanSellOnRagfair                       bool       `json:"CanSellOnRagfair,omitempty"`
	CanRequireOnRagfair                    bool       `json:"CanRequireOnRagfair,omitempty"`
	ConflictingItems                       []string   `json:"ConflictingItems,omitempty"`
	Unlootable                             bool       `json:"Unlootable,omitempty"`
	UnlootableFromSlot                     string     `json:"UnlootableFromSlot,omitempty"`
	UnlootableFromSide                     []string   `json:"UnlootableFromSide,omitempty"`
	AnimationVariantsNumber                int        `json:"AnimationVariantsNumber,omitempty"`
	DiscardingBlock                        bool       `json:"DiscardingBlock,omitempty"`
	RagFairCommissionModifier              int        `json:"RagFairCommissionModifier,omitempty"`
	IsAlwaysAvailableForInsurance          bool       `json:"IsAlwaysAvailableForInsurance,omitempty"`
	DiscardLimit                           int        `json:"DiscardLimit,omitempty"`
	DropSoundType                          string     `json:"DropSoundType,omitempty"`
	InsuranceDisabled                      bool       `json:"InsuranceDisabled,omitempty"`
	QuestStashMaxCount                     int        `json:"QuestStashMaxCount,omitempty"`
	IsSpecialSlotOnly                      bool       `json:"IsSpecialSlotOnly,omitempty"`
	IsUnremovable                          bool       `json:"IsUnremovable,omitempty"`
	Grids                                  []Grid     `json:"Grids,omitempty"`
	Slots                                  []Slot     `json:"Slots,omitempty"`
	CanPutIntoDuringTheRaid                bool       `json:"CanPutIntoDuringTheRaid,omitempty"`
	CantRemoveFromSlotsDuringRaid          []string   `json:"CantRemoveFromSlotsDuringRaid,omitempty"`
	Durability                             int        `json:"Durability,omitempty"`
	Accuracy                               int        `json:"Accuracy,omitempty"`
	Recoil                                 float32    `json:"Recoil,omitempty"`
	Loudness                               int        `json:"Loudness,omitempty"`
	EffectiveDistance                      int        `json:"EffectiveDistance,omitempty"`
	Ergonomics                             float32    `json:"Ergonomics,omitempty"`
	Velocity                               float32    `json:"Velocity,omitempty"`
	RaidModdable                           bool       `json:"RaidModdable,omitempty"`
	ToolModdable                           bool       `json:"ToolModdable,omitempty"`
	BlocksFolding                          bool       `json:"BlocksFolding,omitempty"`
	BlocksCollapsible                      bool       `json:"BlocksCollapsible,omitempty"`
	IsAnimated                             bool       `json:"IsAnimated,omitempty"`
	HasShoulderContact                     bool       `json:"HasShoulderContact,omitempty"`
	SightingRange                          int        `json:"SightingRange,omitempty"`
	DoubleActionAccuracyPenaltyMult        float32    `json:"DoubleActionAccuracyPenaltyMult,omitempty"`
	HeatFactor                             float32    `json:"HeatFactor,omitempty"`
	CoolFactor                             float32    `json:"CoolFactor,omitempty"`
	KnifeHitDelay                          int        `json:"knifeHitDelay,omitempty"`
	KnifeHitSlashRate                      int        `json:"knifeHitSlashRate,omitempty"`
	KnifeHitStabRate                       int        `json:"knifeHitStabRate,omitempty"`
	KnifeHitRadius                         float32    `json:"knifeHitRadius,omitempty"`
	KnifeHitSlashDam                       int        `json:"knifeHitSlashDam,omitempty"`
	KnifeHitStabDam                        int        `json:"knifeHitStabDam,omitempty"`
	KnifeDurability                        int        `json:"knifeDurab,omitempty"`
	MaxDurability                          int        `json:"MaxDurability,omitempty"`
	PrimaryDistance                        float32    `json:"PrimaryDistance,omitempty"`
	SecondryDistance                       float32    `json:"SecondryDistance,omitempty"`
	SlashPenetration                       int        `json:"SlashPenetration,omitempty"`
	StabPenetration                        int        `json:"StabPenetration,omitempty"`
	MinRepairDegradation                   float32    `json:"MinRepairDegradation,omitempty"`
	MaxRepairDegradation                   float32    `json:"MaxRepairDegradation,omitempty"`
	PrimaryConsumption                     int        `json:"PrimaryConsumption,omitempty"`
	SecondryConsumption                    int        `json:"SecondryConsumption,omitempty"`
	DeflectionConsumption                  int        `json:"DeflectionConsumption,omitempty"`
	StimulatorBuffs                        string     `json:"StimulatorBuffs,omitempty"`
	MaxResource                            int        `json:"MaxResource,omitempty"`
	AppliedTrunkRotation                   XYZ        `json:"AppliedTrunkRotation,omitempty"`
	AppliedHeadRotation                    XYZ        `json:"AppliedHeadRotation,omitempty"`
	DisplayOnModel                         bool       `json:"DisplayOnModel,omitempty"`
	AdditionalAnimationLayer               int        `json:"AdditionalAnimationLayer,omitempty"`
	StaminaBurnRate                        int        `json:"StaminaBurnRate,omitempty"`
	ColliderScaleMultiplier                XYZ        `json:"ColliderScaleMultiplier,omitempty"`
	DogTagQualities                        bool       `json:"DogTagQualities,omitempty"`
	FoodUseTime                            int        `json:"foodUseTime,omitempty"`
	FoodEffectType                         string     `json:"foodEffectType,omitempty"`
	AmmoType                               string     `json:"ammoType,omitempty"`
	InitialSpeed                           int        `json:"InitialSpeed,omitempty"`
	BallisticCoefficient                   float32    `json:"BallisticCoeficient,omitempty"`
	BulletMassGram                         float32    `json:"BulletMassGram,omitempty"`
	BulletDiameterMilimeters               float32    `json:"BulletDiameterMilimeters,omitempty"`
	Damage                                 int        `json:"Damage,omitempty"`
	AmmoAccr                               int        `json:"ammoAccr,omitempty"`
	AmmoRec                                int        `json:"ammoRec,omitempty"`
	AmmoDist                               int        `json:"ammoDist,omitempty"`
	BuckshotBullets                        int        `json:"buckshotBullets,omitempty"`
	PenetrationPower                       int        `json:"PenetrationPower,omitempty"`
	PenetrationPowerDiviation              float32    `json:"PenetrationPowerDiviation,omitempty"`
	AmmoHear                               int        `json:"ammoHear,omitempty"`
	AmmoSfx                                string     `json:"ammoSfx,omitempty"`
	MisfireChance                          float32    `json:"MisfireChance,omitempty"`
	MinFragmentsCount                      int        `json:"MinFragmentsCount,omitempty"`
	MaxFragmentsCount                      int        `json:"MaxFragmentsCount,omitempty"`
	AmmoShiftChance                        int        `json:"ammoShiftChance,omitempty"`
	CasingName                             string     `json:"casingName,omitempty"`
	CasingEjectPower                       int        `json:"casingEjectPower,omitempty"`
	CasingMass                             float32    `json:"casingMass,omitempty"`
	CasingSounds                           string     `json:"casingSounds,omitempty"`
	ProjectileCount                        int        `json:"ProjectileCount,omitempty"`
	PenetrationChance                      float32    `json:"PenetrationChance,omitempty"`
	RicochetChance                         float32    `json:"RicochetChance,omitempty"`
	FragmentationChance                    float32    `json:"FragmentationChance,omitempty"`
	Deterioration                          float32    `json:"Deterioration,omitempty"`
	SpeedRetardation                       float32    `json:"SpeedRetardation,omitempty"`
	Tracer                                 bool       `json:"Tracer,omitempty"`
	TracerColor                            string     `json:"TracerColor,omitempty"`
	TracerDistance                         float32    `json:"TracerDistance,omitempty"`
	ArmorDamage                            int        `json:"ArmorDamage,omitempty"`
	Caliber                                string     `json:"Caliber,omitempty"`
	StaminaBurnPerDamage                   float32    `json:"StaminaBurnPerDamage,omitempty"`
	HeavyBleedingDelta                     float32    `json:"HeavyBleedingDelta,omitempty"`
	LightBleedingDelta                     float32    `json:"LightBleedingDelta,omitempty"`
	ShowBullet                             bool       `json:"ShowBullet,omitempty"`
	HasGrenaderComponent                   bool       `json:"HasGrenaderComponent,omitempty"`
	FuzeArmTimeSec                         float32    `json:"FuzeArmTimeSec,omitempty"`
	ExplosionStrength                      int        `json:"ExplosionStrength,omitempty"`
	MinExplosionDistance                   float32    `json:"MinExplosionDistance,omitempty"`
	MaxExplosionDistance                   float32    `json:"MaxExplosionDistance,omitempty"`
	FragmentsCount                         int        `json:"FragmentsCount,omitempty"`
	FragmentType                           string     `json:"FragmentType,omitempty"`
	ShowHitEffectOnExplode                 bool       `json:"ShowHitEffectOnExplode,omitempty"`
	ExplosionType                          string     `json:"ExplosionType,omitempty"`
	AmmoLifeTimeSec                        int        `json:"AmmoLifeTimeSec,omitempty"`
	Contusion                              XYZ        `json:"Contusion,omitempty"`
	ArmorDistanceDistanceDamage            XYZ        `json:"ArmorDistanceDistanceDamage,omitempty"`
	Blindness                              XYZ        `json:"Blindness,omitempty"`
	IsLightAndSoundShot                    bool       `json:"IsLightAndSoundShot,omitempty"`
	LightAndSoundShotAngle                 int        `json:"LightAndSoundShotAngle,omitempty"`
	LightAndSoundShotSelfContusionTime     int        `json:"LightAndSoundShotSelfContusionTime,omitempty"`
	LightAndSoundShotSelfContusionStrength float32    `json:"LightAndSoundShotSelfContusionStrength,omitempty"`
	MalfMisfireChance                      float32    `json:"MalfMisfireChance,omitempty"`
	DurabilityBurnModificator              float32    `json:"DurabilityBurnModificator,omitempty"`
	MalfFeedChance                         float32    `json:"MalfFeedChance,omitempty"`
	RemoveShellAfterFire                   bool       `json:"RemoveShellAfterFire,omitempty"`
	ErgonomicsModifier                     float32    `json:"ErgonomicsModifier,omitempty"`
	VerticalRecoilModifier                 float32    `json:"VerticalRecoilModifier,omitempty"`
	HorizontalRecoilModifier               float32    `json:"HorizontalRecoilModifier,omitempty"`
	MuzzleVelocityModifier                 float32    `json:"MuzzleVelocityModifier,omitempty"`
	AccuracyModifier                       float32    `json:"AccuracyModifier,omitempty"`
	Heaviness                              float32    `json:"Heaviness,omitempty"`
	RecoilModifier                         float32    `json:"RecoilModifier,omitempty"`
	VelocityModifier                       float32    `json:"VelocityModifier,omitempty"`
	MagAnimationIndex                      int        `json:"magAnimationIndex,omitempty"`
	Cartridges                             []Cartride `json:"Cartridges,omitempty"`
	CanFast                                bool       `json:"CanFast,omitempty"`
	CanHit                                 bool       `json:"CanHit,omitempty"`
	CanAdmin                               bool       `json:"CanAdmin,omitempty"`
	LoadUnloadModifier                     int        `json:"LoadUnloadModifier,omitempty"`
	CheckTimeModifier                      int        `json:"CheckTimeModifier,omitempty"`
	CheckOverride                          int        `json:"CheckOverride,omitempty"`
	ReloadMagType                          string     `json:"ReloadMagType,omitempty"`
	VisibleAmmoRangesString                string     `json:"VisibleAmmoRangesString,omitempty"`
	MalfunctionChance                      float32    `json:"MalfunctionChance,omitempty"`
	TagColor                               int        `json:"TagColor,omitempty"`
	TagName                                string     `json:"TagName,omitempty"`
	LinkedWeapon                           string     `json:"LinkedWeapon,omitempty"`
	UseAmmoWithoutShell                    bool       `json:"UseAmmoWithoutShell,omitempty"`
	Chambers                               []Chamber  `json:"Chambers,omitempty"`
}

type Prefab struct {
	Path string `json:"path,omitempty"`
	Rcid string `json:"rcid,omitempty"`
}

type Chamber struct {
	Name                  string  `json:"_name,omitempty"`
	ID                    string  `json:"_id,omitempty"`
	Parent                string  `json:"_parent,omitempty"`
	Props                 Filters `json:"_props,omitempty"`
	Required              bool    `json:"_required,omitempty"`
	MergeSlotWithChildren bool    `json:"_mergeSlotWithChildren,omitempty"`
	Proto                 string  `json:"_proto,omitempty"`
}

type Cartride struct {
	Name     string  `json:"_name,omitempty"`
	Id       string  `json:"_id,omitempty"`
	Parent   string  `json:"_parent,omitempty"`
	MaxCount int     `json:"_max_count,omitempty"`
	Props    Filters `json:"_props,omitempty"`
	Proto    string  `json:"_proto,omitempty"`
}

type Grid struct {
	ID             string  `json:"_id,omitempty"`
	Name           string  `json:"_name,omitempty"`
	ParentID       string  `json:"_parent,omitempty"`
	Type           string  `json:"_type,omitempty"`
	Filters        Filters `json:"filters,omitempty"`
	CellsH         int     `json:"cellsH,omitempty"`
	CellsV         int     `json:"cellsV,omitempty"`
	MinCount       int     `json:"minCount,omitempty"`
	MaxCount       int     `json:"maxCount,omitempty"`
	MaxWeight      int     `json:"maxWeight,omitempty"`
	IsSortingTable bool    `json:"isSortingTable,omitempty"`
	Proto          string  `json:"_proto,omitempty"`
}

type Slot struct {
	Name                  string  `json:"_name,omitempty"`
	ID                    string  `json:"_id,omitempty"`
	Parent                string  `json:"_parent,omitempty"`
	Properties            Filters `json:"_props,omitempty"`
	Required              bool    `json:"_required,omitempty"`
	MergeSlotWithChildren bool    `json:"_mergeSlotWithChildren,omitempty"`
	Proto                 string  `json:"_proto,omitempty"`
}

type Filters struct {
	Filters [1]FilterEntry `json:"filters,omitempty"`
}

type FilterEntry struct {
	Filters []string `json:"Filters,omitempty"`
}

type ItemPreset struct {
	ID           string       `json:"_id,omitempty"`
	Type         string       `json:"_type,omitempty"`
	ChangeWeapon bool         `json:"_changeWeaponName,omitempty"`
	Name         string       `json:"_name,omitempty"`
	Parent       string       `json:"_parent,omitempty"`
	Items        []PresetItem `json:"_items,omitempty"`
	Encyclopedia string       `json:"_encyclopedia,omitempty"`
}

type PresetItem struct {
	ID       string  `json:"_id,omitempty"`
	Tpl      string  `json:"_tpl,omitempty"`
	Upd      ItemUpd `json:"upd,omitempty"`
	ParentID string  `json:"parentId,omitempty"`
	SlotID   string  `json:"slotId,omitempty"`
}

type InventoryItem struct {
	ID       string                `json:"_id,omitempty"`
	Tpl      string                `json:"_tpl,omitempty"`
	Upd      ItemUpd               `json:"upd,omitempty"`
	ParentID string                `json:"parentId,omitempty"`
	SlotID   string                `json:"slotId,omitempty"`
	Location InventoryItemLocation `json:"location,omitempty"`
}

type InventoryItemLocation struct {
	X          float32 `json:"x,omitempty"`
	Y          float32 `json:"y,omitempty"`
	Z          float32 `json:"z,omitempty"`
	R          float32 `json:"r,omitempty"`
	IsSearched bool    `json:"isSearched,omitempty"`
	Rotation   string  `json:"rotation,omitempty"`
}

type ItemUpd struct {
	BuyRestrictionMax     int `json:"BuyRestrictionMax,omitempty"`
	BuyRestrictionCurrent int `json:"BuyRestrictionCurrent,omitempty"`
	Foldable              struct {
		Folded bool `json:"Folded,omitempty"`
	} `json:"Foldable,omitempty"`
	Togglable struct {
		On bool `json:"On,omitempty"`
	} `json:"Togglable,omitempty"`
	FireMode struct {
		FireMode string `json:"FireMode,omitempty"`
	} `json:"FireMode,omitempty"`
	StackObjectsCount int `json:"StackObjectsCount,omitempty"`
	Repairable        struct {
		MaxDurability float32 `json:"MaxDurability,omitempty"`
		Durability    float32 `json:"Durability,omitempty"`
	} `json:"Repairable,omitempty"`
	Sight struct {
		ScopesCurrentCalibPointIndexes []int `json:"ScopesCurrentCalibPointIndexes,omitempty"`
		ScopesSelectedModes            []int `json:"ScopesSelectedModes,omitempty"`
		SelectedScope                  int   `json:"SelectedScope,omitempty"`
	} `json:"Sight,omitempty"`
	FoodDrink struct {
		HpPercent int `json:"HpPercent,omitempty"`
	} `json:"FoodDrink,omitempty"`
	Resource   Value `json:"Resource,omitempty"`
	SideEffect Value `json:"SideEffect,omitempty"`
	MedKit     struct {
		HpResource int `json:"HpResource,omitempty"`
	} `json:"MedKit,omitempty"`
	RepairKit struct {
		Resource int `json:"Resource,omitempty"`
	} `json:"RepairKit,omitempty"`
	Key struct {
		NumberOfUsages int `json:"NumberOfUsages,omitempty"`
	} `json:"Key,omitempty"`
	SpawnedInSession bool `json:"SpawnedInSession,omitempty"`
	Dogtag           struct {
		AccountId       string `json:"AccountId,omitempty"`
		ProfileId       string `json:"ProfileId,omitempty"`
		Nickname        string `json:"Nickname,omitempty"`
		Side            string `json:"Side,omitempty"`
		Level           int    `json:"Level,omitempty"`
		Time            string `json:"Time,omitempty"`
		Status          string `json:"Status,omitempty"`
		KillerAccountId string `json:"KillerAccountId,omitempty"`
		KillerProfileId string `json:"KillerProfileId,omitempty"`
		KillerName      string `json:"KillerName,omitempty"`
		WeaponName      string `json:"WeaponName,omitempty"`
	} `json:"Dogtag,omitempty"`
	Light struct {
		IsActive     bool `json:"IsActive,omitempty"`
		SelectedMode int  `json:"SelectedMode,omitempty"`
	} `json:"Light,omitempty"`
	Buff struct {
		Rarity              string `json:"rarity,omitempty"`
		BuffType            string `json:"buffType,omitempty"`
		Value               int    `json:"value,omitempty"`
		ThresholdDurability int    `json:"thresholdDurability,omitempty"`
	} `json:"Buff,omitempty"`
	Map struct {
		Markers []MapMarker `json:"Markers,omitempty"`
	} `json:"Map,omitempty"`
	FaceShield struct {
		Hits int `json:"Hits,omitempty"`
	} `json:"FaceShield,omitempty"`
	Tag struct {
		Color float32 `json:"Color,omitempty"`
		Name  string  `json:"Name,omitempty"`
	} `json:"Tag,omitempty"`
}

type MapMarker XY
