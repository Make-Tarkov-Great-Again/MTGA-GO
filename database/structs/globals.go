package structs

type Globals struct {
	Config               GlobalsConfig          `json:"config,omitempty"`
	BotPresets           [18]BotPreset          `json:"bot_presets,omitempty"`
	BotWeaponScatterings [4]BotWeaponScattering `json:"BotWeaponScatterings,omitempty"`
	ItemPresets          map[string]ItemPreset  `json:"ItemPresets,omitempty"`
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

type GlobalsConfig struct {
	Content                   Content        `json:"content,omitempty"`
	AimPunchMagnitude         int            `json:"AimPunchMagnitude,omitempty"`
	WeaponSkillProgressRate   float32        `json:"WeaponSkillProgressRate,omitempty"`
	SkillAtrophy              bool           `json:"SkillAtrophy,omitempty"`
	Exp                       Exp            `json:"exp,omitempty"`
	TBaseLooting              int            `json:"t_base_looting,omitempty"`
	TBaseLockpicking          int            `json:"t_base_lockpicking,omitempty"`
	Armor                     Armor          `json:"armor,omitempty"`
	SessionsToShowHotKeys     int            `json:"SessionsToShowHotKeys,omitempty"`
	MaxBotsAliveOnMap         int            `json:"MaxBotsAliveOnMap,omitempty"`
	SavagePlayCooldown        int            `json:"SavagePlayCooldown,omitempty"`
	SavagePlayCooldownNdaFree int            `json:"SavagePlayCooldownNdaFree,omitempty"`
	MarksmanAccuracy          float32        `json:"MarksmanAccuracy,omitempty"`
	SavagePlayCooldownDevelop int            `json:"SavagePlayCooldownDevelop,omitempty"`
	TODSkyDate                string         `json:"TODSkyDate,omitempty"`
	Mastering                 [70]Mastering  `json:"Mastering,omitempty"`
	GlobalItemPriceModifier   int            `json:"GlobalItemPriceModifier,omitempty"`
	TradingUnlimitedItems     bool           `json:"TradingUnlimitedItems,omitempty"`
	MaxLoyaltyLevelForAll     bool           `json:"MaxLoyaltyLevelForAll,omitempty"`
	GlobalLootChanceModifier  float32        `json:"GlobalLootChanceModifier,omitempty"`
	TimeBeforeDeploy          int            `json:"TimeBeforeDeploy,omitempty"`
	TimeBeforeDeployLocal     int            `json:"TimeBeforeDeployLocal,omitempty"`
	LoadTimeSpeedProgress     int            `json:"LoadTimeSpeedProgress,omitempty"`
	BaseLoadTime              float32        `json:"BaseLoadTime,omitempty"`
	BaseUnloadTime            float32        `json:"BaseUnloadTime,omitempty"`
	BaseCheckTime             int            `json:"BaseCheckTime,omitempty"`
	Customization             Customization  `json:"Customization,omitempty"`
	UncheckOnShot             bool           `json:"UncheckOnShot,omitempty"`
	BotsEnabled               bool           `json:"BotsEnabled,omitempty"`
	ArmorMaterials            ArmorMaterials `json:"ArmorMaterials,omitempty"`
	LegsOverdamage            int            `json:"LegsOverdamage,omitempty"`
	HandsOverdamage           float32        `json:"HandsOverdamage,omitempty"`
	StomachOverdamage         float32        `json:"StomachOverdamage,omitempty"`
	Health                    Health         `json:"Health,omitempty"`
	Rating                    Rating         `json:"rating,omitempty"`
	Tournament                Tournament     `json:"tournament,omitempty"`
	RagFair                   RagFair        `json:"RagFair,omitempty"`
	Handbook                  struct {
		DefaultCategory string `json:"defaultCategory,omitempty"`
	} `json:"handbook,omitempty"`
	FractureCausedByFalling       Probability           `json:"FractureCausedByFalling,omitempty"`
	FractureCausedByBulletHit     Probability           `json:"FractureCausedByBulletHit,omitempty"`
	WAVE_COEF_LOW                 float32               `json:"WAVE_COEF_LOW,omitempty"`
	WAVE_COEF_MID                 float32               `json:"WAVE_COEF_MID,omitempty"`
	WAVE_COEF_HIGH                float32               `json:"WAVE_COEF_HIGH,omitempty"`
	WAVE_COEF_HORDE               float32               `json:"WAVE_COEF_HORDE,omitempty"`
	Stamina                       Stamina               `json:"Stamina,omitempty"`
	StaminaRestoration            StaminaParameters     `json:"StaminaRestoration,omitempty"`
	StaminaDrain                  StaminaParameters     `json:"StaminaDrain,omitempty"`
	RequirementReferences         RequirementReferences `json:"RequirementReferences,omitempty"`
	RestrictionsInRaid            RestrictionsInRaid    `json:"RestrictionsInRaids,omitempty"`
	SkillMinEffectiveness         float32               `json:"SkillMinEffectiveness,omitempty"`
	SkillFatiguePerPoint          float32               `json:"SkillFatiguePerPoint,omitempty"`
	SkillFreshEffectiveness       float32               `json:"SkillFreshEffectiveness,omitempty"`
	SkillFreshPoints              int                   `json:"SkillFreshPoints,omitempty"`
	SkillPointsBeforeFatigue      int                   `json:"SkillPointsBeforeFatigue,omitempty"`
	SkillFatigueReset             int                   `json:"SkillFatigueReset,omitempty"`
	EventType                     [1]string             `json:"EventType,omitempty"`
	WalkSpeed                     XYZ                   `json:"WalkSpeed,omitempty"`
	SprintSpeed                   XYZ                   `json:"SprintSpeed,omitempty"`
	SkillEnduranceWeightThreshold float32               `json:"SkillEnduranceWeightThreshold,omitempty"`
	TeamSearchingTimeout          int                   `json:"TeamSearchingTimeout,omitempty"`
	Insurance                     struct {
		MaxStorageTimeInHour int `json:"MaxStorageTimeInHour,omitempty"`
	} `json:"Insurance,omitempty"`
	SkillExpPerLevel                   int           `json:"SkillExpPerLevel,omitempty"`
	GameSearchingTimeout               int           `json:"GameSearchingTimeout,omitempty"`
	WallContusionAbsorption            XYZ           `json:"WallContusionAbsorption,omitempty"`
	AzimuthPanelShowsPlayerOrientation bool          `json:"AzimuthPanelShowsPlayerOrientation,omitempty"`
	Aiming                             Aiming        `json:"Aiming,omitempty"`
	Malfunction                        Malfunction   `json:"Malfunction,omitempty"`
	Overheat                           Overheat      `json:"Overheat,omitempty"`
	FenceSettings                      FenceSettings `json:"FenceSettings,omitempty"`
	TestValue                          int           `json:"TestValue,omitempty"`
	Inertia                            Inertia       `json:"Inertia,omitempty"`
	Ballistic                          struct {
		GlobalDamageDegradationCoefficient float32 `json:"GlobalDamageDegradationCoefficient,omitempty"`
	} `json:"Ballistic,omitempty"`
	RepairSettings       RepairSettings `json:"RepairSettings,omitempty"`
	DiscardLimitsEnabled bool           `json:"DiscardLimitsEnabled,omitempty"`
	CoopSettings         struct {
		AvailableVersions [2]string `json:"AvailableVersions,omitempty"`
	} `json:"CoopSettings,omitempty"`
	BufferZone      BufferZone `json:"BufferZone,omitempty"`
	TradingSetting  int        `json:"TradingSetting,omitempty"`
	TradingSettings struct {
		BuyoutRestrictions struct {
			MinFoodDrinkResource float32 `json:"MinFoodDrinkResource,omitempty"`
			MinMedsResource      float32 `json:"MinMedsResource,omitempty"`
			MinDurability        float32 `json:"MinDurability,omitempty"`
		} `json:"BuyoutRestrictions,omitempty"`
	} `json:"TradingSettings,omitempty"`
	ItemsCommonSettings struct {
		ItemRemoveAfterInterruptionTime int `json:"ItemRemoveAfterInterruptionTime,omitempty"`
	} `json:"ItemsCommonSettings,omitempty"`
	SquadSettings struct {
		SendRequestDelaySeconds    int `json:"SendRequestDelaySeconds,omitempty"`
		SecondsForExpiredRequest   int `json:"SecondsForExpiredRequest,omitempty"`
		CountOfRequestsToOnePlayer int `json:"CountOfRequestsToOnePlayer,omitempty"`
	}
}

type BotWeaponScattering struct {
	Name                    string  `json:"Name,omitempty"`
	PriorityScatter1meter   float32 `json:"PriorityScatter1meter,omitempty"`
	PriorityScatter10meter  float32 `json:"PriorityScatter10meter,omitempty"`
	PriorityScatter100meter float32 `json:"PriorityScatter100meter,omitempty"`
}

type BotPreset struct {
	UseThis                    bool    `json:"UseThis,omitempty"`
	Role                       string  `json:"Role,omitempty"`
	BotDifficulty              string  `json:"BotDifficulty,omitempty"`
	VisibleAngle               float32 `json:"VisibleAngle,omitempty"`
	VisibleDistance            float32 `json:"VisibleDistance,omitempty"`
	ScatteringPerMeter         float32 `json:"ScatteringPerMeter,omitempty"`
	HearingSense               float32 `json:"HearingSense,omitempty"`
	SCATTERING_DIST_MODIF      float32 `json:"SCATTERING_DIST_MODIF,omitempty"`
	MAX_AIMING_UPGRADE_BY_TIME float32 `json:"MAX_AIMING_UPGRADE_BY_TIME,omitempty"`
	FIRST_CONTACT_ADD_SEC      float32 `json:"FIRST_CONTACT_ADD_SEC,omitempty"`
	COEF_IF_MOVE               float32 `json:"COEF_IF_MOVE,omitempty"`
}

type BufferZone struct {
	CustomerCriticalTimeStart int `json:"CustomerCriticalTimeStart,omitempty"`
	CustomerKickNotifTime     int `json:"CustomerKickNotifTime,omitempty"`
	CustomerAccessTime        int `json:"CustomerAccessTime,omitempty"`
}

type RepairSettings struct {
	ArmorClassDivisor        int                 `json:"armorClassDivisor,omitempty"`
	DurabilityPointCostArmor float32             `json:"durabilityPointCostArmor,omitempty"`
	DurabilityPointCostGuns  float32             `json:"durabilityPointCostGuns,omitempty"`
	RepairStrategies         map[string]Strategy `json:"RepairStrategies,omitempty"`
	MinimumLevelToApplyBuff  int                 `json:"MinimumLevelToApplyBuff,omitempty"`
	ItemEnhancementSettings  EnhancementSettings `json:"ItemEnhancementSettings,omitempty"`
}

type Strategy struct {
	Filter    []string `json:"Filter,omitempty"`
	BuffTypes []string `json:"BuffTypes,omitempty"`
}

type EnhancementSettings struct {
	WeaponSpread           PriceModifier `json:"WeaponSpread,omitempty"`
	DamageReduction        PriceModifier `json:"DamageReduction,omitempty"`
	MalfunctionProtections PriceModifier `json:"MalfunctionProtections,omitempty"`
}

type PriceModifier struct {
	PriceModifier float32 `json:"PriceModifier,omitempty"`
}

type Inertia struct {
	InertiaLimits                      XYZ     `json:"InertiaLimits,omitempty"`
	InertiaLimitsStep                  float32 `json:"InertiaLimitsStep,omitempty"`
	ExitMovementStateSpeedThreshold    XYZ     `json:"ExitMovementStateSpeedThreshold,omitempty"`
	WalkInertia                        XYZ     `json:"WalkInertia,omitempty"`
	FallThreshold                      float32 `json:"FallThreshold,omitempty"`
	SpeedLimitAfterFallMin             XYZ     `json:"SpeedLimitAfterFallMin,omitempty"`
	SpeedLimitAfterFallMax             XYZ     `json:"SpeedLimitAfterFallMax,omitempty"`
	SpeedLimitDurationMin              XYZ     `json:"SpeedLimitDurationMin,omitempty"`
	SpeedLimitDurationMax              XYZ     `json:"SpeedLimitDurationMax,omitempty"`
	SpeedInertiaAfterJump              XYZ     `json:"SpeedInertiaAfterJump,omitempty"`
	BaseJumpPenaltyDuration            float32 `json:"BaseJumpPenaltyDuration,omitempty"`
	DurationPower                      float32 `json:"DurationPower,omitempty"`
	BaseJumpPenalty                    float32 `json:"BaseJumpPenalty,omitempty"`
	PenaltyPower                       float32 `json:"PenaltyPower,omitempty"`
	InertiaTiltCurveMin                XYZ     `json:"InertiaTiltCurveMin,omitempty"`
	InertiaTiltCurveMax                XYZ     `json:"InertiaTiltCurveMax,omitempty"`
	InertiaBackwardCoef                XYZ     `json:"InertiaBackwardCoef,omitempty"`
	TiltInertiaMaxSpeed                XYZ     `json:"TiltInertiaMaxSpeed,omitempty"`
	TiltStartSideBackSpeed             XYZ     `json:"TiltStartSideBackSpeed,omitempty"`
	TiltMaxSideBackSpeed               XYZ     `json:"TiltMaxSideBackSpeed,omitempty"`
	TiltAcceleration                   XYZ     `json:"TiltAcceleration,omitempty"`
	AverageRotationFrameSpan           int     `json:"AverageRotationFrameSpan,omitempty"`
	SprintSpeedInertiaCurveMin         XYZ     `json:"SprintSpeedInertiaCurveMin,omitempty"`
	SprintSpeedInertiaCurveMax         XYZ     `json:"SprintSpeedInertiaCurveMax,omitempty"`
	SprintBrakeInertia                 XYZ     `json:"SprintBrakeInertia,omitempty"`
	SprintTransitionMotionPreservation XYZ     `json:"SprintTransitionMotionPreservation,omitempty"`
	WeaponFlipSpeed                    XYZ     `json:"WeaponFlipSpeed,omitempty"`
	PreSprintAccelerationLimits        XYZ     `json:"PreSprintAccelerationLimits,omitempty"`
	SprintAccelerationLimits           XYZ     `json:"SprintAccelerationLimits,omitempty"`
	SideTime                           XYZ     `json:"SideTime,omitempty"`
	DiagonalTime                       XYZ     `json:"DiagonalTime,omitempty"`
	MinDirectionBlendTime              float32 `json:"MinDirectionBlendTime,omitempty"`
	MinMovementAccelerationRangeRight  XYZ     `json:"MinMovementAccelerationRangeRight,omitempty"`
	MaxMovementAccelerationRangeRight  XYZ     `json:"MaxMovementAccelerationRangeRight,omitempty"`
	MaxTimeWithoutInput                XYZ     `json:"MaxTimeWithoutInput,omitempty"`
	ProneSpeedAccelerationRange        XYZ     `json:"ProneSpeedAccelerationRange,omitempty"`
	ProneDirectionAccelerationRange    XYZ     `json:"ProneDirectionAccelerationRange,omitempty"`
	MoveTimeRange                      XYZ     `json:"MoveTimeRange,omitempty"`
	CrouchSpeedAccelerationRange       XYZ     `json:"CrouchSpeedAccelerationRange,omitempty"`
}

type FenceSettings struct {
	FenceId                   string             `json:"FenceId,omitempty"`
	Levels                    map[int]FenceLevel `json:"Levels,omitempty"`
	PaidExitStandingNumerator float32            `json:"paidExitStandingNumerator,omitempty"`
}

type FenceLevel struct {
	SavageCooldownModifier           float32 `json:"SavageCooldownModifier,omitempty"`
	ScavCaseTimeModifier             float32 `json:"ScavCaseTimeModifier,omitempty"`
	PaidExitCostModifier             float32 `json:"PaidExitCostModifier,omitempty"`
	BotFollowChance                  float32 `json:"BotFollowChance,omitempty"`
	ScavEquipmentSpawnChanceModifier float32 `json:"ScavEquipmentSpawnChanceModifier,omitempty"`
	PriceModifier                    float32 `json:"PriceModifier,omitempty"`
	HostileBosses                    bool    `json:"HostileBosses,omitempty"`
	HostileScavs                     bool    `json:"HostileScavs,omitempty"`
	ScavAttackSupport                bool    `json:"ScavAttackSupport,omitempty"`
	ExfiltrationPriceModifier        float32 `json:"ExfiltrationPriceModifier,omitempty"`
	AvailableExits                   int     `json:"AvailableExits,omitempty"`
}

type Overheat struct {
	MinOverheat                 float32 `json:"MinOverheat,omitempty"`
	MaxOverheat                 float32 `json:"MaxOverheat,omitempty"`
	OverheatProblemsStart       float32 `json:"OverheatProblemsStart,omitempty"`
	ModHeatFactor               float32 `json:"ModHeatFactor,omitempty"`
	ModCoolFactor               float32 `json:"ModCoolFactor,omitempty"`
	MinWearOnOverheat           float32 `json:"MinWearOnOverheat,omitempty"`
	MaxWearOnOverheat           float32 `json:"MaxWearOnOverheat,omitempty"`
	MinWearOnMaxOverheat        float32 `json:"MinWearOnMaxOverheat,omitempty"`
	MaxWearOnMaxOverheat        float32 `json:"MaxWearOnMaxOverheat,omitempty"`
	OverheatWearLimit           float32 `json:"OverheatWearLimit,omitempty"`
	MaxCOIIncreaseMult          float32 `json:"MaxCOIIncreaseMult,omitempty"`
	MinMalfChance               float32 `json:"MinMalfChance,omitempty"`
	MaxMalfChance               float32 `json:"MaxMalfChance,omitempty"`
	DurReduceMinMult            float32 `json:"DurReduceMinMult,omitempty"`
	DurReduceMaxMult            float32 `json:"DurReduceMaxMult,omitempty"`
	BarrelMoveRndDuration       float32 `json:"BarrelMoveRndDuration,omitempty"`
	BarrelMoveMaxMult           float32 `json:"BarrelMoveMaxMult,omitempty"`
	FireratePitchMult           float32 `json:"FireratePitchMult,omitempty"`
	FirerateReduceMinMult       float32 `json:"FirerateReduceMinMult,omitempty"`
	FirerateReduceMaxMult       float32 `json:"FirerateReduceMaxMult,omitempty"`
	FirerateOverheatBorder      float32 `json:"FirerateOverheatBorder,omitempty"`
	EnableSlideOnMaxOverheat    bool    `json:"EnableSlideOnMaxOverheat,omitempty"`
	StartSlideOverheat          float32 `json:"StartSlideOverheat,omitempty"`
	FixSlideOverheat            float32 `json:"FixSlideOverheat,omitempty"`
	AutoshotMinOverheat         float32 `json:"AutoshotMinOverheat,omitempty"`
	AutoshotChance              float32 `json:"AutoshotChance,omitempty"`
	AutoshotPossibilityDuration float32 `json:"AutoshotPossibilityDuration,omitempty"`
	MaxOverheatCoolCoef         float32 `json:"MaxOverheatCoolCoef,omitempty"`
}

type Malfunction struct {
	AmmoMalfChanceMult           float32 `json:"AmmoMalfChanceMult,omitempty"`
	MagazineMalfChanceMult       float32 `json:"MagazineMalfChanceMult,omitempty"`
	MalfRepairHardSlideMult      float32 `json:"MalfRepairHardSlideMult,omitempty"`
	MalfRepairOneHandBrokenMult  float32 `json:"MalfRepairOneHandBrokenMult,omitempty"`
	MalfRepairTwoHandsBrokenMult float32 `json:"MalfRepairTwoHandsBrokenMult,omitempty"`
	AllowMalfForBots             bool    `json:"AllowMalfForBots,omitempty"`
	ShowGlowAttemptsCount        int     `json:"ShowGlowAttemptsCount,omitempty"`
	OutToIdleSpeedMultForPistol  int     `json:"OutToIdleSpeedMultForPistol,omitempty"`
	IdleToOutSpeedMultOnMalf     int     `json:"IdleToOutSpeedMultOnMalf,omitempty"`
	TimeToQuickdrawPistol        int     `json:"TimeToQuickdrawPistol,omitempty"`
	DurRangeToIgnoreMalfs        XYZ     `json:"DurRangeToIgnoreMalfs,omitempty"`
	DurFeedWt                    float32 `json:"DurFeedWt,omitempty"`
	DurMisfireWt                 float32 `json:"DurMisfireWt,omitempty"`
	DurJamWt                     float32 `json:"DurJamWt,omitempty"`
	DurSoftSlideWt               float32 `json:"DurSoftSlideWt,omitempty"`
	DurHardSlideMinWt            float32 `json:"DurHardSlideMinWt,omitempty"`
	DurHardSlideMaxWt            float32 `json:"DurHardSlideMaxWt,omitempty"`
	AmmoMisfireWt                float32 `json:"AmmoMisfireWt,omitempty"`
	AmmoFeedWt                   float32 `json:"AmmoFeedWt,omitempty"`
	AmmoJamWt                    float32 `json:"AmmoJamWt,omitempty"`
	OverheatFeedWt               float32 `json:"OverheatFeedWt,omitempty"`
	OverheatJamWt                float32 `json:"OverheatJamWt,omitempty"`
	OverheatSoftSlideWt          float32 `json:"OverheatSoftSlideWt,omitempty"`
	OverheatHardSlideMinWt       float32 `json:"OverheatHardSlideMinWt,omitempty"`
	OverheatHardSlideMaxWt       float32 `json:"OverheatHardSlideMaxWt,omitempty"`
}

type Aiming struct {
	ProceduralIntensityByPose XYZ     `json:"ProceduralIntensityByPose,omitempty"`
	AimProceduralIntensity    float32 `json:"AimProceduralIntensity,omitempty"`
	HeavyWeight               float32 `json:"HeavyWeight,omitempty"`
	LightWeight               float32 `json:"LightWeight,omitempty"`
	MaxTimeHeavy              float32 `json:"MaxTimeHeavy,omitempty"`
	MinTimeHeavy              float32 `json:"MinTimeHeavy,omitempty"`
	MaxTimeLight              float32 `json:"MaxTimeLight,omitempty"`
	MinTimeLight              float32 `json:"MinTimeLight,omitempty"`
	RecoilScaling             float32 `json:"RecoilScaling,omitempty"`
	RecoilDamping             float32 `json:"RecoilDamping,omitempty"`
	CameraSnapGlobalMult      float32 `json:"CameraSnapGlobalMult,omitempty"`
	RecoilXIntensityByPose    XYZ     `json:"RecoilXIntensityByPose,omitempty"`
	RecoilYIntensityByPose    XYZ     `json:"RecoilYIntensityByPose,omitempty"`
	RecoilZIntensityByPose    XYZ     `json:"RecoilZIntensityByPose,omitempty"`
	RecoilCrank               bool    `json:"RecoilCrank,omitempty"`
	RecoilHandDamping         float32 `json:"RecoilHandDamping,omitempty"`
	RecoilConvergenceMult     float32 `json:"RecoilConvergenceMult,omitempty"`
	RecoilVertBonus           int     `json:"RecoilVertBonus,omitempty"`
	RecoilBackBonus           int     `json:"RecoilBackBonus,omitempty"`
}

type Content struct {
	IP   string `json:"ip,omitempty"`
	Port int    `json:"port,omitempty"`
	Root string `json:"root,omitempty"`
}

type Exp struct {
	Heal struct {
		ExpForHeal      int `json:"expForHeal,omitempty"`
		ExpForHydration int `json:"expForHydration,omitempty"`
		ExpForEnergy    int `json:"expForEnergy,omitempty"`
	} `json:"heal,omitempty"`
	MatchEnd struct {
		README                     string  `json:"README,omitempty"`
		SurvivedExpRequirement     int     `json:"survived_exp_requirement,omitempty"`
		SurvivedSecondsRequirement int     `json:"survived_seconds_requirement,omitempty"`
		SurvivedExpReward          int     `json:"survived_exp_reward,omitempty"`
		MiaExpReward               int     `json:"mia_exp_reward,omitempty"`
		RunnerExpReward            int     `json:"runner_exp_reward,omitempty"`
		LeftMult                   int     `json:"leftMult,omitempty"`
		MiaMult                    float32 `json:"miaMult,omitempty"`
		SurvivedMult               float32 `json:"survivedMult,omitempty"`
		RunnerMult                 float32 `json:"runnerMult,omitempty"`
		KilledMult                 int     `json:"killedMult,omitempty"`
	} `json:"match_end,omitempty"`
	Kill struct {
		Combo []struct {
			Percent int `json:"percent,omitempty"`
		} `json:"combo,omitempty"`
		VictimLevelExp       int     `json:"victimLevelExp,omitempty"`
		HeadShotMult         float32 `json:"headShotMult,omitempty"`
		ExpOnDamageAllHealth int     `json:"expOnDamageAllHealth,omitempty"`
		LongShotDistance     int     `json:"longShotDistance,omitempty"`
		BloodLossToLitre     float32 `json:"bloodLossToLitre,omitempty"`
		VictimBotLevelExp    int     `json:"victimBotLevelExp,omitempty"`
	} `json:"kill,omitempty"`
	Level struct {
		ExpTable []struct {
			Exp int `json:"exp,omitempty"`
		} `json:"exp_table,omitempty"`
		TradeLevel  int `json:"trade_level,omitempty"`
		SavageLevel int `json:"savage_level,omitempty"`
		ClanLevel   int `json:"clan_level,omitempty"`
		Mastering1  int `json:"mastering1,omitempty"`
		Mastering2  int `json:"mastering2,omitempty"`
	} `json:"level,omitempty"`
	LootAttempts []struct {
		K_Exp float32 `json:"k_exp,omitempty"`
	} `json:"loot_attempts,omitempty"`
	ExpForLockedDoorOpen   int `json:"expForLockedDoorOpen,omitempty"`
	ExpForLockedDoorBreach int `json:"expForLockedDoorBreach,omitempty"`
	TriggerMult            int `json:"triggerMult,omitempty"`
}

type Armor struct {
	Class []struct {
		Resistance int `json:"resistance,omitempty"`
	} `json:"class,omitempty"`
}

type Mastering struct {
	Templates []string `json:"Templates,omitempty"`
	Level2    int      `json:"Level2,omitempty"`
	Level3    int      `json:"Level3,omitempty"`
}

type Masteries struct {
	Name      string      `json:"Name,omitempty"`
	Templates []Mastering `json:"Templates,omitempty"`
}

type Customization struct {
	SavageHead         map[string]SavageCustomization `json:"SavageHead,omitempty"`
	SavageBody         map[string]SavageCustomization `json:"SavageBody,omitempty"`
	SavageFeet         map[string]SavageCustomization `json:"SavageFeet,omitempty"`
	CustomizationVoice []SavageCustomization          `json:"CustomizationVoice,omitempty"`
	BodyParts          map[string]string              `json:"BodyParts,omitempty"`
}

type SavageCustomization struct {
	Head        string   `json:"head,omitempty"`
	Body        string   `json:"body,omitempty"`
	Hands       string   `json:"hands,omitempty"`
	Feet        string   `json:"feet,omitempty"`
	Voice       string   `json:"voice,omitempty"`
	Side        []string `json:"side,omitempty"`
	IsNotRandom bool     `json:"isNotRandom,omitempty"`
	NotRandom   bool     `json:"NotRandom,omitempty"`
}

type ArmorMaterials struct {
	UHMWPE       ArmorMaterial `json:"UHMWPE,omitempty"`
	Aramid       ArmorMaterial `json:"Aramid,omitempty"`
	Combined     ArmorMaterial `json:"Combined,omitempty"`
	Titan        ArmorMaterial `json:"Titan,omitempty"`
	Aluminium    ArmorMaterial `json:"Aluminium,omitempty"`
	ArmoredSteel ArmorMaterial `json:"ArmoredSteel,omitempty"`
	Ceramic      ArmorMaterial `json:"Ceramic,omitempty"`
	Glass        ArmorMaterial `json:"Glass,omitempty"`
}

type ArmorMaterial struct {
	Destructibility          float32 `json:"Destructibility,omitempty"`
	MinRepairDegradation     float32 `json:"MinRepairDegradation,omitempty"`
	MaxRepairDegradation     float32 `json:"MaxRepairDegradation,omitempty"`
	ExplosionDestructibility float32 `json:"ExplosionDestructibility,omitempty"`
	MinRepairKitDegradation  float32 `json:"MinRepairKitDegradation,omitempty"`
	MaxRepairKitDegradation  float32 `json:"MaxRepairKitDegradation,omitempty"`
}

type Health struct {
	Falling struct {
		DamagePerMeter int `json:"DamagePerMeter,omitempty"`
		SafeHeight     int `json:"SafeHeight,omitempty"`
	} `json:"Falling,omitempty"`
	Effects               Effect                `json:"Effects,omitempty"`
	HealPrice             HealPrice             `json:"HealPrice,omitempty"`
	ProfileHealthSettings ProfileHealthSettings `json:"ProfileHealthSettings,omitempty"`
}

type Effects struct {
	Existence             Effect     `json:"Existence,omitempty"`
	Dehydration           Effect     `json:"Dehydration,omitempty"`
	BreakPart             Effect     `json:"BreakPart,omitempty"`
	Contusion             Effect     `json:"Contusion,omitempty"`
	Disorientation        Effect     `json:"Disorientation,omitempty"`
	Exhaustion            Effect     `json:"Exhaustion,omitempty"`
	LowEdgeHealth         Effect     `json:"LowEdgeHealth,omitempty"`
	RadExposure           Effect     `json:"RadExposure,omitempty"`
	Stun                  Effect     `json:"Stun,omitempty"`
	Intoxication          Effect     `json:"Intoxication,omitempty"`
	Regeneration          Effect     `json:"Regeneration,omitempty"`
	Wound                 Effect     `json:"Wound,omitempty"`
	Berserk               Effect     `json:"Berserk,omitempty"`
	Flash                 Effect     `json:"Flash,omitempty"`
	MedEffect             Effect     `json:"MedEffect,omitempty"`
	Pain                  Effect     `json:"Pain,omitempty"`
	PainKiller            Effect     `json:"PainKiller,omitempty"`
	SandingScreen         Effect     `json:"SandingScreen,omitempty"`
	Stimulator            Stimulator `json:"Stimulator,omitempty"`
	Tremor                Effect     `json:"Tremor,omitempty"`
	ChronicStaminaFatigue Effect     `json:"ChronicStaminaFatigue,omitempty"`
	HeavyBleeding         Effect     `json:"HeavyBleeding,omitempty"`
	LightBleeding         Effect     `json:"LightBleeding,omitempty"`
	BodyTemperature       Effect     `json:"BodyTemperature,omitempty"`
	MildMusclePain        Effect     `json:"MildMusclePain,omitempty"`
	SevereMusclePain      Effect     `json:"SevereMusclePain,omitempty"`
}

type Stimulator struct {
	BuffLoopTime float32           `json:"BuffLoopTime,omitempty"`
	Buffs        map[string][]Buff `json:"Buffs,omitempty"`
}

type Buff struct {
	BuffType      string   `json:"BuffType,omitempty"`
	Chance        float32  `json:"Chance,omitempty"`
	Delay         int64    `json:"Delay,omitempty"`
	Duration      int64    `json:"Duration,omitempty"`
	Value         float32  `json:"Value,omitempty"`
	AbsoluteValue bool     `json:"AbsoluteValue,omitempty"`
	SkillName     string   `json:"SkillName,omitempty"`
	AppliesTo     []string `json:"AppliesTo,omitempty"`
}

type Effect struct {
	EnergyLoopTime                      int         `json:"EnergyLoopTime,omitempty"`
	HydrationLoopTime                   int         `json:"HydrationLoopTime,omitempty"`
	EnergyDamage                        float32     `json:"EnergyDamage,omitempty"`
	HydrationDamage                     float32     `json:"HydrationDamage,omitempty"`
	DestroyedStomachEnergyTimeFactor    int         `json:"DestroyedStomachEnergyTimeFactor,omitempty"`
	DestroyedStomachHydrationTimeFactor int         `json:"DestroyedStomachHydrationTimeFactor,omitempty"`
	DefaultDelay                        int         `json:"DefaultDelay,omitempty"`
	DefaultResidueTime                  int         `json:"DefaultResidueTime,omitempty"`
	BleedingHealth                      float32     `json:"BleedingHealth,omitempty"`
	BleedingLoopTime                    int         `json:"BleedingLoopTime,omitempty"`
	BleedingLifeTime                    int         `json:"BleedingLifeTime,omitempty"`
	DamageOnStrongDehydration           int         `json:"DamageOnStrongDehydration,omitempty"`
	StrongDehydrationLoopTime           int         `json:"StrongDehydrationLoopTime,omitempty"`
	HealExperience                      int         `json:"HealExperience,omitempty"`
	OfflineDurationMin                  int         `json:"OfflineDurationMin,omitempty"`
	OfflineDurationMax                  int         `json:"OfflineDurationMax,omitempty"`
	RemovePrice                         int         `json:"RemovePrice,omitempty"`
	RemovedAfterDeath                   bool        `json:"RemovedAfterDeath,omitempty"`
	BulletHitProbability                Probability `json:"BulletHitProbability,omitempty"`
	FallingProbability                  Probability `json:"FallingProbability,omitempty"`
	Dummy                               int         `json:"Dummy,omitempty"`
	LoopTime                            int         `json:"LoopTime,omitempty"`
	MinimumHealthPercentage             int         `json:"MinimumHealthPercentage,omitempty"`
	Energy                              int         `json:"Energy,omitempty"`
	Hydration                           int         `json:"Hydration,omitempty"`
	BodyHealth                          BodyHealth  `json:"BodyHealth,omitempty"`
	Influences                          Influences  `json:"Influences,omitempty"`
	WorkingTime                         int         `json:"WorkingTime,omitempty"`
	ThresholdMin                        int         `json:"ThresholdMin,omitempty"`
	ThresholdMax                        int         `json:"ThresholdMax,omitempty"`
	GymEffectivity                      float32     `json:"GymEffectivity,omitempty"`
	TraumaChance                        int         `json:"TraumaChance,omitempty"`
}

type Probability struct {
	FunctionType string  `json:"FunctionType,omitempty"`
	K            float32 `json:"K,omitempty"`
	B            float32 `json:"B,omitempty"`
	Threshold    float32 `json:"Threshold,omitempty"`
}

type BodyHealth struct {
	Head     Value `json:"Head,omitempty"`
	Chest    Value `json:"Chest,omitempty"`
	Stomach  Value `json:"Stomach,omitempty"`
	LeftArm  Value `json:"LeftArm,omitempty"`
	RightArm Value `json:"RightArm,omitempty"`
	LeftLeg  Value `json:"LeftLeg,omitempty"`
	RightLeg Value `json:"RightLeg,omitempty"`
}

type Value struct {
	Value         float32 `json:"Value,omitempty"`
	UnitsConsumed int     `json:"UnitsConsumed,omitempty"`
}

type Influences struct {
	LightBleeding SlowDownPercentage `json:"LightBleeding,omitempty"`
	HeavyBleeding SlowDownPercentage `json:"HeavyBleeding,omitempty"`
	Fracture      SlowDownPercentage `json:"Fracture,omitempty"`
	RadExposure   SlowDownPercentage `json:"RadExposure,omitempty"`
	Intoxication  SlowDownPercentage `json:"Intoxication,omitempty"`
}

type SlowDownPercentage struct {
	HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage,omitempty"`
	EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage,omitempty"`
	HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage,omitempty"`
}

type HealPrice struct {
	HealthPointPrice    int `json:"HealthPointPrice,omitempty"`
	HydrationPointPrice int `json:"HydrationPointPrice,omitempty"`
	EnergyPointPrice    int `json:"EnergyPointPrice,omitempty"`
	TrialLevels         int `json:"TrialLevels,omitempty"`
	TrialRaids          int `json:"TrialRaids,omitempty"`
}

type ProfileHealthSettings struct {
	BodyPartsSettings     BodyPartsSettings     `json:"BodyPartsSettings,omitempty"`
	HealthFactorsSettings HealthFactorsSettings `json:"HealthFactorsSettings,omitempty"`
	DefaultStimulatorBuff string                `json:"DefaultStimulatorBuff,omitempty"`
}

type BodyPartsSettings struct {
	Head     HealthFactorsSettings `json:"Head,omitempty"`
	Chest    HealthFactorsSettings `json:"Chest,omitempty"`
	Stomach  HealthFactorsSettings `json:"Stomach,omitempty"`
	LeftArm  HealthFactorsSettings `json:"LeftArm,omitempty"`
	RightArm HealthFactorsSettings `json:"RightArm,omitempty"`
	LeftLeg  HealthFactorsSettings `json:"LeftLeg,omitempty"`
	RightLeg HealthFactorsSettings `json:"RightLeg,omitempty"`
}

type HealthFactorSettings struct {
	Minimum                      int     `json:"Minimum,omitempty"`
	Maximum                      int     `json:"Maximum,omitempty"`
	Default                      float32 `json:"Default,omitempty"`
	OverDamageReceivedMultiplier float32 `json:"OverDamageReceivedMultiplier,omitempty"`
}

type HealthFactorsSettings struct {
	Energy      HealthFactorSettings `json:"Energy,omitempty"`
	Hydration   HealthFactorSettings `json:"Hydration,omitempty"`
	Temperature HealthFactorSettings `json:"Temperature,omitempty"`
	Poisoning   HealthFactorSettings `json:"Poisoning,omitempty"`
	Radiation   HealthFactorSettings `json:"Radiation,omitempty"`
}

type Rating struct {
	LevelRequired int        `json:"levelRequired,omitempty"`
	Limit         int        `json:"limit,omitempty"`
	Categories    Categories `json:"categories,omitempty"`
}

type Categories struct {
	Experience        bool `json:"experience,omitempty"`
	KD                bool `json:"kd,omitempty"`
	SurviveRatio      bool `json:"surviveRatio,omitempty"`
	AvgEarnings       bool `json:"avgEarnings,omitempty"`
	Kills             bool `json:"kills,omitempty"`
	RaidCount         bool `json:"raidCount,omitempty"`
	LongestShot       bool `json:"longestShot,omitempty"`
	TimeOnline        bool `json:"timeOnline,omitempty"`
	InventoryFullCost bool `json:"inventoryFullCost,omitempty"`
	RagFairStanding   bool `json:"ragFairStanding,omitempty"`
}

type Tournament struct {
	Categories    TournamentCategories `json:"categories,omitempty"`
	Limit         int                  `json:"limit,omitempty"`
	LevelRequired int                  `json:"levelRequired,omitempty"`
}

type TournamentCategories struct {
	Dogtags bool `json:"dogtags,omitempty"`
}

type RagFair struct {
	Enabled                                 bool                    `json:"enabled,omitempty"`
	PriceStabilizerEnabled                  bool                    `json:"priceStabilizerEnabled,omitempty"`
	IncludePveTraderSales                   bool                    `json:"includePveTraderSales,omitempty"`
	PriceStabilizerStartIntervalInHours     int                     `json:"priceStabilizerStartIntervalInHours,omitempty"`
	MinUserLevel                            int                     `json:"minUserLevel,omitempty"`
	CommunityTax                            int                     `json:"communityTax,omitempty"`
	CommunityItemTax                        int                     `json:"communityItemTax,omitempty"`
	CommunityRequirementTax                 int                     `json:"communityRequirementTax,omitempty"`
	OfferPriorityCost                       int                     `json:"offerPriorityCost,omitempty"`
	OfferDurationTimeInHour                 int                     `json:"offerDurationTimeInHour,omitempty"`
	OfferDurationTimeInHourAfterRemove      float32                 `json:"offerDurationTimeInHourAfterRemove,omitempty"`
	PriorityTimeModifier                    int                     `json:"priorityTimeModifier,omitempty"`
	MaxRenewOfferTimeInHour                 int                     `json:"maxRenewOfferTimeInHour,omitempty"`
	RenewPricePerHour                       float32                 `json:"renewPricePerHour,omitempty"`
	MaxActiveOfferCount                     []ActiveOfferCountRange `json:"maxActiveOfferCount,omitempty"`
	BalancerRemovePriceCoefficient          int                     `json:"balancerRemovePriceCoefficient,omitempty"`
	BalancerMinPriceCount                   int                     `json:"balancerMinPriceCount,omitempty"`
	BalancerAveragePriceCoefficient         int                     `json:"balancerAveragePriceCoefficient,omitempty"`
	DelaySinceOfferAdd                      int                     `json:"delaySinceOfferAdd,omitempty"`
	UniqueBuyerTimeoutInDays                int                     `json:"uniqueBuyerTimeoutInDays,omitempty"`
	RatingSumForIncrease                    int                     `json:"ratingSumForIncrease,omitempty"`
	RatingIncreaseCount                     float32                 `json:"ratingIncreaseCount,omitempty"`
	RatingSumForDecrease                    int                     `json:"ratingSumForDecrease,omitempty"`
	RatingDecreaseCount                     float32                 `json:"ratingDecreaseCount,omitempty"`
	MaxSumForIncreaseRatingPerOneSale       int                     `json:"maxSumForIncreaseRatingPerOneSale,omitempty"`
	MaxSumForDecreaseRatingPerOneSale       int                     `json:"maxSumForDecreaseRatingPerOneSale,omitempty"`
	MaxSumForRarity                         RarityMaxSums           `json:"maxSumForRarity,omitempty"`
	ChangePriceCoef                         int                     `json:"ChangePriceCoef,omitempty"`
	BalancerUserItemSaleCooldownEnabled     bool                    `json:"balancerUserItemSaleCooldownEnabled,omitempty"`
	BalancerUserItemSaleCooldown            int                     `json:"balancerUserItemSaleCooldown,omitempty"`
	YouSellOfferMaxStorageTimeInHour        int                     `json:"youSellOfferMaxStorageTimeInHour,omitempty"`
	YourOfferDidNotSellMaxStorageTimeInHour int                     `json:"yourOfferDidNotSellMaxStorageTimeInHour,omitempty"`
	IsOnlyFoundInRaidAllowed                bool                    `json:"isOnlyFoundInRaidAllowed,omitempty"`
	SellInOnePiece                          int                     `json:"sellInOnePiece,omitempty"`
}

type ActiveOfferCountRange struct {
	From  float32 `json:"from,omitempty"`
	To    float32 `json:"to,omitempty"`
	Count int     `json:"count,omitempty"`
}

type RarityMaxSums struct {
	Common    Value `json:"Common,omitempty"`
	Rare      Value `json:"Rare,omitempty"`
	Superrare Value `json:"Superrare,omitempty"`
	NotExist  Value `json:"Not_exist,omitempty"`
}

type Stamina struct {
	Capacity                           int     `json:"Capacity,omitempty"`
	SprintDrainRate                    float32 `json:"SprintDrainRate,omitempty"`
	BaseRestorationRate                float32 `json:"BaseRestorationRate,omitempty"`
	JumpConsumption                    int     `json:"JumpConsumption,omitempty"`
	GrenadeHighThrow                   int     `json:"GrenadeHighThrow,omitempty"`
	GrenadeLowThrow                    int     `json:"GrenadeLowThrow,omitempty"`
	AimDrainRate                       float32 `json:"AimDrainRate,omitempty"`
	AimRangeFinderDrainRate            float32 `json:"AimRangeFinderDrainRate,omitempty"`
	OxygenCapacity                     int     `json:"OxygenCapacity,omitempty"`
	OxygenRestoration                  int     `json:"OxygenRestoration,omitempty"`
	WalkOverweightLimits               XYZ     `json:"WalkOverweightLimits,omitempty"`
	BaseOverweightLimits               XYZ     `json:"BaseOverweightLimits,omitempty"`
	SprintOverweightLimits             XYZ     `json:"SprintOverweightLimits,omitempty"`
	WalkSpeedOverweightLimits          XYZ     `json:"WalkSpeedOverweightLimits,omitempty"`
	CrouchConsumption                  XYZ     `json:"CrouchConsumption,omitempty"`
	WalkConsumption                    XYZ     `json:"WalkConsumption,omitempty"`
	StandupConsumption                 XYZ     `json:"StandupConsumption,omitempty"`
	TransitionSpeed                    XYZ     `json:"TransitionSpeed,omitempty"`
	SprintAccelerationLowerLimit       float32 `json:"SprintAccelerationLowerLimit,omitempty"`
	SprintSpeedLowerLimit              float32 `json:"SprintSpeedLowerLimit,omitempty"`
	SprintSensitivityLowerLimit        float32 `json:"SprintSensitivityLowerLimit,omitempty"`
	AimConsumptionByPose               XYZ     `json:"AimConsumptionByPose,omitempty"`
	RestorationMultiplierByPose        XYZ     `json:"RestorationMultiplierByPose,omitempty"`
	OverweightConsumptionByPose        XYZ     `json:"OverweightConsumptionByPose,omitempty"`
	AimingSpeedMultiplier              float32 `json:"AimingSpeedMultiplier,omitempty"`
	WalkVisualEffectMultiplier         float32 `json:"WalkVisualEffectMultiplier,omitempty"`
	HandsCapacity                      int     `json:"HandsCapacity,omitempty"`
	HandsRestoration                   float32 `json:"HandsRestoration,omitempty"`
	ProneConsumption                   int     `json:"ProneConsumption,omitempty"`
	BaseHoldBreathConsumption          int     `json:"BaseHoldBreathConsumption,omitempty"`
	SoundRadius                        XYZ     `json:"SoundRadius,omitempty"`
	ExhaustedMeleeSpeed                float32 `json:"ExhaustedMeleeSpeed,omitempty"`
	FatigueRestorationRate             float32 `json:"FatigueRestorationRate,omitempty"`
	FatigueAmountToCreateEffect        int     `json:"FatigueAmountToCreateEffect,omitempty"`
	ExhaustedMeleeDamageMultiplier     float32 `json:"ExhaustedMeleeDamageMultiplier,omitempty"`
	FallDamageMultiplier               float32 `json:"FallDamageMultiplier,omitempty"`
	SafeHeightOverweight               float32 `json:"SafeHeightOverweight,omitempty"`
	SitToStandConsumption              int     `json:"SitToStandConsumption,omitempty"`
	StaminaExhaustionCausesJiggle      bool    `json:"StaminaExhaustionCausesJiggle,omitempty"`
	StaminaExhaustionStartsBreathSound bool    `json:"StaminaExhaustionStartsBreathSound,omitempty"`
	StaminaExhaustionRocksCamera       bool    `json:"StaminaExhaustionRocksCamera,omitempty"`
	HoldBreathStaminaMultiplier        XYZ     `json:"HoldBreathStaminaMultiplier,omitempty"`
	PoseLevelIncreaseSpeed             XYZ     `json:"PoseLevelIncreaseSpeed,omitempty"`
	PoseLevelDecreaseSpeed             XYZ     `json:"PoseLevelDecreaseSpeed,omitempty"`
	PoseLevelConsumptionPerNotch       XYZ     `json:"PoseLevelConsumptionPerNotch,omitempty"`
}

type XYZ struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
	Z float32 `json:"z,omitempty"`
}

type StaminaParameters struct {
	LowerLeftPoint  float32 `json:"LowerLeftPoint,omitempty"`
	LowerRightPoint float32 `json:"LowerRightPoint,omitempty"`
	LeftPlatoPoint  float32 `json:"LeftPlatoPoint,omitempty"`
	RightPlatoPoint float32 `json:"RightPlatoPoint,omitempty"`
	RightLimit      float32 `json:"RightLimit,omitempty"`
	ZeroValue       float32 `json:"ZeroValue,omitempty"`
}

type RequirementReferences struct {
	Alpinist []AlpinistRequirement `json:"Alpinist,omitempty"`
}

type AlpinistRequirement struct {
	Requirement    string `json:"Requirement,omitempty"`
	ID             string `json:"Id,omitempty"`
	Count          int    `json:"Count,omitempty"`
	RequiredSlot   string `json:"RequiredSlot,omitempty"`
	RequirementTip string `json:"RequirementTip,omitempty"`
}

type RestrictionsInRaid struct {
	Templates []RaidRestriction `json:"RestrictionsInRaid,omitempty"`
}

type RaidRestriction struct {
	TemplateID string `json:"TemplateId,omitempty"`
	MaxInLobby int    `json:"MaxInLobby,omitempty"`
	MaxInRaid  int    `json:"MaxInRaid,omitempty"`
}

type SkillsSettings struct {
	SkillProgressRate              float32                `json:"SkillProgressRate,omitempty"`
	WeaponSkillProgressRate        int                    `json:"WeaponSkillProgressRate,omitempty"`
	WeaponSkillRecoilBonusPerLevel float32                `json:"WeaponSkillRecoilBonusPerLevel,omitempty"`
	HideoutManagement              HideoutManagementSkill `json:"HideoutManagement,omitempty"`
	Crafting                       CraftingSkill          `json:"Crafting,omitempty"`
	Metabolism                     MetabolismSkill        `json:"Metabolism,omitempty"`
	Immunity                       ImmunitySkill          `json:"Immunity,omitempty"`
	Endurance                      EnduranceSkill         `json:"Endurance,omitempty"`
	Strength                       StrengthSkill          `json:"Strength,omitempty"`
	Vitality                       VitalitySkill          `json:"Vitality,omitempty"`
	Health                         HealthSkill            `json:"Health,omitempty"`
	StressResistance               StressResistanceSkill  `json:"StressResistance,omitempty"`
	Throwing                       ThrowingSkill          `json:"Throwing,omitempty"`
	RecoilControl                  RecoilControlSkill     `json:"RecoilControl,omitempty"`
	Pistol                         WeaponHandlingSkill    `json:"Pistol,omitempty"`
	Revolver                       WeaponHandlingSkill    `json:"Revolver,omitempty"`
	SMG                            []interface{}          `json:"SMG,omitempty"`
	Assault                        WeaponHandlingSkill    `json:"Assault,omitempty"`
	Shotgun                        WeaponHandlingSkill    `json:"Shotgun,omitempty"`
	Sniper                         WeaponHandlingSkill    `json:"Sniper,omitempty"`
	LMG                            []interface{}          `json:"LMG,omitempty"`
	HMG                            []interface{}          `json:"HMG,omitempty"`
	Launcher                       []interface{}          `json:"Launcher,omitempty"`
	AttachedLauncher               []interface{}          `json:"AttachedLauncher,omitempty"`
	Melee                          MeleeSkill             `json:"Melee,omitempty"`
	DMR                            WeaponHandlingSkill    `json:"DMR,omitempty"`
	BearAssaultoperations          []interface{}          `json:"BearAssaultoperations,omitempty"`
	BearAuthority                  []interface{}          `json:"BearAuthority,omitempty"`
	BearAksystems                  []interface{}          `json:"BearAksystems,omitempty"`
	BearHeavycaliber               []interface{}          `json:"BearHeavycaliber,omitempty"`
	BearRawpower                   []interface{}          `json:"BearRawpower,omitempty"`
	UsecArsystems                  []interface{}          `json:"UsecArsystems,omitempty"`
	UsecDeepweaponmodding_Settings []interface{}          `json:"UsecDeepweaponmodding_Settings,omitempty"`
	UsecLongrangeoptics_Settings   []interface{}          `json:"UsecLongrangeoptics_Settings,omitempty"`
	UsecNegotiations               []interface{}          `json:"UsecNegotiations,omitempty"`
	UsecTactics                    []interface{}          `json:"UsecTactics,omitempty"`
	BotReload                      []interface{}          `json:"BotReload,omitempty"`
	CovertMovement                 CovertMovementSkill    `json:"CovertMovement,omitempty"`
	FieldMedicine                  []interface{}          `json:"FieldMedicine,omitempty"`
	Search                         SearchSkill            `json:"Search,omitempty"`
	Sniping                        []interface{}          `json:"Sniping,omitempty"`
	ProneMovement                  []interface{}          `json:"ProneMovement,omitempty"`
	FirstAid                       []interface{}          `json:"FirstAid,omitempty"`
	LightVests                     VestSkill              `json:"LightVests,omitempty"`
	HeavyVests                     VestSkill              `json:"HeavyVests,omitempty"`
	WeaponModding                  []interface{}          `json:"WeaponModding,omitempty"`
	AdvancedModding                []interface{}          `json:"AdvancedModding,omitempty"`
	NightOps                       []interface{}          `json:"NightOps,omitempty"`
	SilentOps                      []interface{}          `json:"SilentOps,omitempty"`
	Lockpicking                    []interface{}          `json:"Lockpicking,omitempty"`
	WeaponTreatment                WeaponTreatmentSkill   `json:"WeaponTreatment,omitempty"`
	MagDrills                      MagDrillsSkill         `json:"MagDrills,omitempty"`
	Freetrading                    []interface{}          `json:"Freetrading,omitempty"`
	Auctions                       []interface{}          `json:"Auctions,omitempty"`
	Cleanoperations                []interface{}          `json:"Cleanoperations,omitempty"`
	Barter                         []interface{}          `json:"Barter,omitempty"`
	Shadowconnections              []interface{}          `json:"Shadowconnections,omitempty"`
	Taskperformance                []interface{}          `json:"Taskperformance,omitempty"`
	Perception                     PerceptionSkill        `json:"Perception,omitempty"`
	Intellect                      IntellectSkill         `json:"Intellect,omitempty"`
	Attention                      AttentionSkill         `json:"Attention,omitempty"`
	Charisma                       CharismaSkill          `json:"Charisma,omitempty"`
	Memory                         MemorySkill            `json:"Memory,omitempty"`
	Surgery                        SurgerySkill           `json:"Surgery,omitempty"`
	AimDrills                      AimDrillsSkill         `json:"AimDrills,omitempty"`
	BotSound                       []interface{}          `json:"BotSound,omitempty"`
	TroubleShooting                TroubleShootingSkill   `json:"TroubleShooting,omitempty"`
}

type WeaponTreatmentSkill struct {
	WearAmountRepairGunsReducePerLevel   float32                 `json:"WearAmountRepairGunsReducePerLevel,omitempty"`
	DurLossReducePerLevel                float32                 `json:"DurLossReducePerLevel,omitempty"`
	SkillPointsPerRepair                 int                     `json:"SkillPointsPerRepair,omitempty"`
	Filter                               []string                `json:"Filter,omitempty"`
	WearChanceRepairGunsReduceEliteLevel float32                 `json:"WearChanceRepairGunsReduceEliteLevel,omitempty"`
	BuffSettings                         WeaponTreatmentBuff     `json:"BuffSettings,omitempty"`
	BuffMaxCount                         int                     `json:"BuffMaxCount,omitempty"`
	Counters                             WeaponTreatmentCounters `json:"Counters,omitempty"`
}

type WeaponTreatmentBuff struct {
	CommonBuffMinChanceValue          float32 `json:"CommonBuffMinChanceValue,omitempty"`
	CommonBuffChanceLevelBonus        float32 `json:"CommonBuffChanceLevelBonus,omitempty"`
	RareBuffChanceCoff                float32 `json:"RareBuffChanceCoff,omitempty"`
	ReceivedDurabilityMaxPercent      int     `json:"ReceivedDurabilityMaxPercent,omitempty"`
	CurrentDurabilityLossToRemoveBuff float32 `json:"CurrentDurabilityLossToRemoveBuff,omitempty"`
	MaxDurabilityLossToRemoveBuff     float32 `json:"MaxDurabilityLossToRemoveBuff,omitempty"`
}

type WeaponTreatmentCounters struct {
	FirearmsDurability CounterPoints `json:"firearmsDurability,omitempty"`
}

type MagDrillsSkill struct {
	RaidLoadedAmmoAction   float32 `json:"RaidLoadedAmmoAction,omitempty"`
	RaidUnloadedAmmoAction float32 `json:"RaidUnloadedAmmoAction,omitempty"`
	MagazineCheckAction    float32 `json:"MagazineCheckAction,omitempty"`
}

type VestSkill struct {
	WearAmountRepairVestsReducePerLevel          float32          `json:"WearAmountRepairVestsReducePerLevel,omitempty"`
	WearChanceRepairVestsReduceEliteLevel        float32          `json:"WearChanceRepairVestsReduceEliteLevel,omitempty"`
	MoveSpeedPenaltyReductionVestsReducePerLevel float32          `json:"MoveSpeedPenaltyReductionVestsReducePerLevel,omitempty"`
	BuffSettings                                 VestBuffSettings `json:"BuffSettings,omitempty"`
	BuffMaxCount                                 int              `json:"BuffMaxCount,omitempty"`
	Counters                                     VestCounters     `json:"Counters,omitempty"`
}

type VestBuffSettings struct {
	CommonBuffMinChanceValue          float32 `json:"CommonBuffMinChanceValue,omitempty"`
	CommonBuffChanceLevelBonus        float32 `json:"CommonBuffChanceLevelBonus,omitempty"`
	RareBuffChanceCoff                float32 `json:"RareBuffChanceCoff,omitempty"`
	ReceivedDurabilityMaxPercent      int     `json:"ReceivedDurabilityMaxPercent,omitempty"`
	CurrentDurabilityLossToRemoveBuff float32 `json:"CurrentDurabilityLossToRemoveBuff,omitempty"`
	MaxDurabilityLossToRemoveBuff     float32 `json:"MaxDurabilityLossToRemoveBuff,omitempty"`
}

type VestCounters struct {
	ArmorDurability CounterPoints `json:"armorDurability,omitempty"`
}

type CounterPoints struct {
	Points  int     `json:"points,omitempty"`
	Divisor float32 `json:"divisor,omitempty"`
}

type SkillPointsRate struct {
	Generator        SkillPoints `json:"Generator,omitempty"`
	AirFilteringUnit SkillPoints `json:"AirFilteringUnit,omitempty"`
	WaterCollector   SkillPoints `json:"WaterCollector,omitempty"`
	SolarPower       SkillPoints `json:"SolarPower,omitempty"`
}

type SkillPoints struct {
	ResourceSpent int `json:"ResourceSpent,omitempty"`
	PointsGained  int `json:"PointsGained,omitempty"`
}

type EliteSlots struct {
	Generator        SkillSlotContainer `json:"Generator,omitempty"`
	AirFilteringUnit SkillSlotContainer `json:"AirFilteringUnit,omitempty"`
	WaterCollector   SkillSlotContainer `json:"WaterCollector,omitempty"`
	BitcoinFarm      SkillSlotContainer `json:"BitcoinFarm,omitempty"`
}

type SkillSlotContainer struct {
	Slots     int `json:"Slots,omitempty"`
	Container int `json:"Container,omitempty"`
}

type HideoutManagementSkill struct {
	SkillPointsPerAreaUpgrade    int             `json:"SkillPointsPerAreaUpgrade,omitempty"`
	SkillPointsPerCraft          int             `json:"SkillPointsPerCraft,omitempty"`
	ConsumptionReductionPerLevel float32         `json:"ConsumptionReductionPerLevel,omitempty"`
	SkillBoostPercent            int             `json:"SkillBoostPercent,omitempty"`
	SkillPointsRate              SkillPointsRate `json:"SkillPointsRate,omitempty"`
	EliteSlots                   EliteSlots      `json:"EliteSlots,omitempty"`
}

type CraftingSkill struct {
	PointsPerCraftingCycle          int     `json:"PointsPerCraftingCycle,omitempty"`
	CraftingCycleHours              int     `json:"CraftingCycleHours,omitempty"`
	PointsPerUniqueCraftCycle       float32 `json:"PointsPerUniqueCraftCycle,omitempty"`
	UniqueCraftsPerCycle            int     `json:"UniqueCraftsPerCycle,omitempty"`
	CraftTimeReductionPerLevel      float32 `json:"CraftTimeReductionPerLevel,omitempty"`
	ProductionTimeReductionPerLevel float32 `json:"ProductionTimeReductionPerLevel,omitempty"`
	EliteExtraProductions           int     `json:"EliteExtraProductions,omitempty"`
	CraftingPointsToIntelligence    int     `json:"CraftingPointsToIntelligence,omitempty"`
}

type MetabolismSkill struct {
	HydrationRecoveryRate              int `json:"HydrationRecoveryRate,omitempty"`
	EnergyRecoveryRate                 int `json:"EnergyRecoveryRate,omitempty"`
	IncreasePositiveEffectDurationRate int `json:"IncreasePositiveEffectDurationRate,omitempty"`
	DecreaseNegativeEffectDurationRate int `json:"DecreaseNegativeEffectDurationRate,omitempty"`
	DecreasePoisonDurationRate         int `json:"DecreasePoisonDurationRate,omitempty"`
}

type ImmunitySkill struct {
	ImmunityMiscEffects    int     `json:"ImmunityMiscEffects,omitempty"`
	ImmunityPoisonBuff     int     `json:"ImmunityPoisonBuff,omitempty"`
	ImmunityPainKiller     int     `json:"ImmunityPainKiller,omitempty"`
	HealthNegativeEffect   float32 `json:"HealthNegativeEffect,omitempty"`
	StimulatorNegativeBuff float32 `json:"StimulatorNegativeBuff,omitempty"`
}

type EnduranceSkill struct {
	MovementAction      float32 `json:"MovementAction,omitempty"`
	SprintAction        float32 `json:"SprintAction,omitempty"`
	GainPerFatigueStack float32 `json:"GainPerFatigueStack,omitempty"`
	QTELevelMultipliers struct {
		Lvl10 struct {
			Multiplier float32 `json:"Multiplier,omitempty"`
		} `json:"10,omitempty"`
		Lvl25 struct {
			Multiplier float32 `json:"Multiplier,omitempty"`
		} `json:"25,omitempty"`
		Lvl50 struct {
			Multiplier float32 `json:"Multiplier,omitempty"`
		} `json:"50,omitempty"`
	} `json:"QTELevelMultipliers,omitempty"`
}

type StrengthSkill struct {
	SprintActionMin     float32 `json:"SprintActionMin,omitempty"`
	SprintActionMax     float32 `json:"SprintActionMax,omitempty"`
	MovementActionMin   float32 `json:"MovementActionMin,omitempty"`
	MovementActionMax   float32 `json:"MovementActionMax,omitempty"`
	PushUpMin           float32 `json:"PushUpMin,omitempty"`
	PushUpMax           float32 `json:"PushUpMax,omitempty"`
	FistfightAction     float32 `json:"FistfightAction,omitempty"`
	ThrowAction         float32 `json:"ThrowAction,omitempty"`
	QTELevelMultipliers []struct {
		Level      int     `json:"Level,omitempty"`
		Multiplier float32 `json:"Multiplier,omitempty"`
	} `json:"QTELevelMultipliers,omitempty"`
}

type VitalitySkill struct {
	DamageTakenAction    float32 `json:"DamageTakenAction,omitempty"`
	HealthNegativeEffect float32 `json:"HealthNegativeEffect,omitempty"`
}

type HealthSkill struct {
	SkillProgress float32 `json:"SkillProgress,omitempty"`
}

type StressResistanceSkill struct {
	HealthNegativeEffect float32 `json:"HealthNegativeEffect,omitempty"`
	LowHPDuration        float32 `json:"LowHPDuration,omitempty"`
}

type ThrowingSkill struct {
	ThrowAction float32 `json:"ThrowAction,omitempty"`
}

type RecoilControlSkill struct {
	RecoilAction        float32 `json:"RecoilAction,omitempty"`
	RecoilBonusPerLevel float32 `json:"RecoilBonusPerLevel,omitempty"`
}

type WeaponHandlingSkill struct {
	WeaponReloadAction  float32 `json:"WeaponReloadAction,omitempty"`
	WeaponShotAction    float32 `json:"WeaponShotAction,omitempty"`
	WeaponFixAction     float32 `json:"WeaponFixAction,omitempty"`
	WeaponChamberAction float32 `json:"WeaponChamberAction,omitempty"`
}

type MeleeSkill struct {
	BuffSettings struct {
		CommonBuffMinChanceValue     int     `json:"CommonBuffMinChanceValue,omitempty"`
		CommonBuffChanceLevelBonus   int     `json:"CommonBuffChanceLevelBonus,omitempty"`
		RareBuffChanceCoff           int     `json:"RareBuffChanceCoff,omitempty"`
		ReceivedDurabilityMaxPercent float32 `json:"ReceivedDurabilityMaxPercent,omitempty"`
	} `json:"BuffSettings,omitempty"`
}

type CovertMovementSkill struct {
	MovementAction float32 `json:"MovementAction,omitempty"`
}

type SearchSkill struct {
	SearchAction float32 `json:"SearchAction,omitempty"`
	FindAction   float32 `json:"FindAction,omitempty"`
}

type PerceptionSkill struct {
	OnlineAction         float32      `json:"OnlineAction,omitempty"`
	UniqueLoot           float32      `json:"UniqueLoot,omitempty"`
	DependentSkillRatios []SkillRatio `json:"DependentSkillRatios,omitempty"`
}

type IntellectSkill struct {
	ExamineAction              int               `json:"ExamineAction,omitempty"`
	SkillProgress              float32           `json:"SkillProgress,omitempty"`
	RepairAction               float32           `json:"RepairAction,omitempty"`
	WearAmountReducePerLevel   float32           `json:"WearAmountReducePerLevel,omitempty"`
	WearChanceReduceEliteLevel float32           `json:"WearChanceReduceEliteLevel,omitempty"`
	RepairPointsCostReduction  float32           `json:"RepairPointsCostReduction,omitempty"`
	Counters                   IntellectCounters `json:"Counters,omitempty"`
	DependentSkillRatios       []SkillRatio      `json:"DependentSkillRatios,omitempty"`
}

type AttentionSkill struct {
	ExamineWithInstruction float32      `json:"ExamineWithInstruction,omitempty"`
	FindActionFalse        float32      `json:"FindActionFalse,omitempty"`
	FindActionTrue         float32      `json:"FindActionTrue,omitempty"`
	DependentSkillRatios   []SkillRatio `json:"DependentSkillRatios,omitempty"`
}

type CharismaSkill struct {
	SkillProgressInt float32               `json:"SkillProgressInt,omitempty"`
	SkillProgressAtn float32               `json:"SkillProgressAtn,omitempty"`
	SkillProgressPer float32               `json:"SkillProgressPer,omitempty"`
	BonusSettings    CharismaBonusSettings `json:"BonusSettings,omitempty"`
	Counters         CharismaCounters      `json:"Counters,omitempty"`
}

type MemorySkill struct {
	AnySkillUp    int     `json:"AnySkillUp,omitempty"`
	SkillProgress float32 `json:"SkillProgress,omitempty"`
}

type SurgerySkill struct {
	SurgeryAction float32 `json:"SurgeryAction,omitempty"`
	SkillProgress float32 `json:"SkillProgress,omitempty"`
}

type AimDrillsSkill struct {
	WeaponShotAction float32 `json:"WeaponShotAction,omitempty"`
}

type TroubleShootingSkill struct {
	MalfRepairSpeedBonusPerLevel    float32 `json:"MalfRepairSpeedBonusPerLevel,omitempty"`
	SkillPointsPerMalfFix           float32 `json:"SkillPointsPerMalfFix,omitempty"`
	EliteDurabilityChanceReduceMult float32 `json:"EliteDurabilityChanceReduceMult,omitempty"`
	EliteAmmoChanceReduceMult       float32 `json:"EliteAmmoChanceReduceMult,omitempty"`
	EliteMagChanceReduceMult        float32 `json:"EliteMagChanceReduceMult,omitempty"`
}

type SkillRatio struct {
	SkillId string  `json:"SkillId,omitempty"`
	Ratio   float32 `json:"Ratio,omitempty"`
}

type IntellectCounters struct {
	FirearmsDurability    CounterPoints `json:"firearmsDurability,omitempty"`
	ArmorDurability       CounterPoints `json:"armorDurability,omitempty"`
	MeleeWeaponDurability CounterPoints `json:"meleeWeaponDurability,omitempty"`
}

type CharismaBonusSettings struct {
	LevelBonusSettings CharismaLevelBonus `json:"LevelBonusSettings,omitempty"`
	EliteBonusSettings CharismaEliteBonus `json:"EliteBonusSettings,omitempty"`
}

type CharismaLevelBonus struct {
	RepeatableQuestChangeDiscount float32 `json:"RepeatableQuestChangeDiscount,omitempty"`
	HealthRestoreDiscount         float32 `json:"HealthRestoreDiscount,omitempty"`
	HealthRestoreTraderDiscount   float32 `json:"HealthRestoreTraderDiscount,omitempty"`
	InsuranceDiscount             float32 `json:"InsuranceDiscount,omitempty"`
	InsuranceTraderDiscount       float32 `json:"InsuranceTraderDiscount,omitempty"`
	PaidExitDiscount              float32 `json:"PaidExitDiscount,omitempty"`
}

type CharismaEliteBonus struct {
	ScavCaseDiscount          float32 `json:"ScavCaseDiscount,omitempty"`
	FenceStandingLossDiscount float32 `json:"FenceStandingLossDiscount,omitempty"`
	RepeatableQuestExtraCount int     `json:"RepeatableQuestExtraCount,omitempty"`
}

type CharismaCounters struct {
	RepeatableQuestCompleteCount CounterPoints `json:"repeatableQuestCompleteCount,omitempty"`
	InsuranceCost                CounterPoints `json:"insuranceCost,omitempty"`
	RepairCost                   CounterPoints `json:"repairCost,omitempty"`
	RestoredHealthCost           CounterPoints `json:"restoredHealthCost,omitempty"`
	ScavCaseCost                 CounterPoints `json:"scavCaseCost,omitempty"`
}
