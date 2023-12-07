package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

var weaponMastering = make(map[string]int16)

func setGlobals() {
	raw := tools.GetJSONRawMessage(globalsFilePath)
	if err := json.UnmarshalNoEscape(raw, &db.core.Globals); err != nil {
		log.Fatalln(err)
	}

	db.core.Globals.Config.Handbook.DefaultCategory = ""
}

func IndexWeaponMasteries() {
	for idx, mastery := range db.core.Globals.Config.Mastering {
		for _, template := range mastery.Templates {
			weaponMastering[template] = int16(idx)
		}
	}
}

func SetNewWeaponMastery(name string) {

}

func GetWeaponMasteryByID(uid string) (*ConfigMastering, error) {
	idx, ok := weaponMastering[uid]
	if !ok {
		return nil, fmt.Errorf("uid does not exist in weapon mastery")
	}

	return &db.core.Globals.Config.Mastering[idx], nil
}

type Globals struct {
	Config               Config         `json:"config"`
	BotPresets           [18]any        `json:"bot_presets"`
	BotWeaponScatterings [4]any         `json:"BotWeaponScatterings"`
	ItemPresets          map[string]any `json:"ItemPresets"`
}

type configAiming struct {
	AimProceduralIntensity    int     `json:"AimProceduralIntensity"`
	CameraSnapGlobalMult      float64 `json:"CameraSnapGlobalMult"`
	HeavyWeight               int     `json:"HeavyWeight"`
	LightWeight               float64 `json:"LightWeight"`
	MaxTimeHeavy              float64 `json:"MaxTimeHeavy"`
	MaxTimeLight              int     `json:"MaxTimeLight"`
	MinTimeHeavy              float64 `json:"MinTimeHeavy"`
	MinTimeLight              float64 `json:"MinTimeLight"`
	ProceduralIntensityByPose Vector3 `json:"ProceduralIntensityByPose"`
	RecoilBackBonus           int     `json:"RecoilBackBonus"`
	RecoilConvergenceMult     int     `json:"RecoilConvergenceMult"`
	RecoilCrank               bool    `json:"RecoilCrank"`
	RecoilDamping             float64 `json:"RecoilDamping"`
	RecoilHandDamping         float64 `json:"RecoilHandDamping"`
	RecoilScaling             int     `json:"RecoilScaling"`
	RecoilVertBonus           int     `json:"RecoilVertBonus"`
	RecoilXIntensityByPose    Vector3 `json:"RecoilXIntensityByPose"`
	RecoilYIntensityByPose    Vector3 `json:"RecoilYIntensityByPose"`
	RecoilZIntensityByPose    Vector3 `json:"RecoilZIntensityByPose"`
}
type configArmorMaterials struct {
	Destructibility          float64 `json:"Destructibility"`
	ExplosionDestructibility float64 `json:"ExplosionDestructibility"`
	MaxRepairDegradation     float64 `json:"MaxRepairDegradation"`
	MaxRepairKitDegradation  float64 `json:"MaxRepairKitDegradation"`
	MinRepairDegradation     float64 `json:"MinRepairDegradation"`
	MinRepairKitDegradation  float64 `json:"MinRepairKitDegradation"`
}

type configAudioSettings struct {
	AudioGroupPresets []audioGroupPresets `json:"AudioGroupPresets"`
}
type audioGroupPresets struct {
	AngleToAllowBinaural       int     `json:"AngleToAllowBinaural"`
	DisabledBinauralByDistance bool    `json:"DisabledBinauralByDistance"`
	DistanceToAllowBinaural    int     `json:"DistanceToAllowBinaural"`
	GroupType                  int     `json:"GroupType"`
	HeightToAllowBinaural      int     `json:"HeightToAllowBinaural"`
	Name                       string  `json:"Name"`
	OcclusionEnabled           bool    `json:"OcclusionEnabled"`
	OcclusionIntensity         int     `json:"OcclusionIntensity"`
	OverallVolume              float64 `json:"OverallVolume"`
}

type customizationVoice struct {
	IsNotRandom bool     `json:"isNotRandom"`
	Side        []string `json:"side"`
	Voice       string   `json:"voice"`
}
type customizationSavageBody struct {
	Body        string `json:"body"`
	Hands       string `json:"hands"`
	IsNotRandom bool   `json:"isNotRandom"`
}
type customizationSavageFeet struct {
	NotRandom   bool   `json:"NotRandom"`
	Feet        string `json:"feet"`
	IsNotRandom bool   `json:"isNotRandom"`
}
type customizationSavageHead struct {
	NotRandom   bool   `json:"NotRandom"`
	Head        string `json:"head"`
	IsNotRandom bool   `json:"isNotRandom"`
}
type configCustomization struct {
	BodyParts struct {
		Body  string `json:"Body"`
		Feet  string `json:"Feet"`
		Hands string `json:"Hands"`
		Head  string `json:"Head"`
	} `json:"BodyParts"`
	CustomizationVoice []customizationVoice               `json:"CustomizationVoice"`
	SavageBody         map[string]customizationSavageBody `json:"SavageBody"`
	SavageFeet         map[string]customizationSavageFeet `json:"SavageFeet"`
	SavageHead         map[string]customizationSavageHead `json:"SavageHead"`
}

type configCoopSettings struct {
	AvailableVersions []string `json:"AvailableVersions"`
}

type fenceSettingsLevels struct {
	AvailableExits                   int     `json:"AvailableExits"`
	BotApplySilenceChance            int     `json:"BotApplySilenceChance"`
	BotFollowChance                  int     `json:"BotFollowChance"`
	BotGetInCoverChance              int     `json:"BotGetInCoverChance"`
	BotHelpChance                    int     `json:"BotHelpChance"`
	BotSpreadoutChance               int     `json:"BotSpreadoutChance"`
	BotStopChance                    int     `json:"BotStopChance"`
	ExfiltrationPriceModifier        float64 `json:"ExfiltrationPriceModifier"`
	HostileBosses                    bool    `json:"HostileBosses"`
	HostileScavs                     bool    `json:"HostileScavs"`
	PaidExitCostModifier             float64 `json:"PaidExitCostModifier"`
	PriceModifier                    float64 `json:"PriceModifier"`
	SavageCooldownModifier           float64 `json:"SavageCooldownModifier"`
	ScavAttackSupport                bool    `json:"ScavAttackSupport"`
	ScavCaseTimeModifier             float64 `json:"ScavCaseTimeModifier"`
	ScavEquipmentSpawnChanceModifier int     `json:"ScavEquipmentSpawnChanceModifier"`
}
type configFenceSettings struct {
	FenceId                   string                         `json:"FenceId"`
	Levels                    map[string]fenceSettingsLevels `json:"Levels"`
	PaidExitStandingNumerator float64                        `json:"paidExitStandingNumerator"`
}

type configBufferZone struct {
	CustomerAccessTime        int `json:"CustomerAccessTime"`
	CustomerCriticalTimeStart int `json:"CustomerCriticalTimeStart"`
	CustomerKickNotifTime     int `json:"CustomerKickNotifTime"`
}
type configBallistic struct {
	GlobalDamageDegradationCoefficient float64 `json:"GlobalDamageDegradationCoefficient"`
}
type configGraphicSettings struct {
	ExperimentalFogInCity bool `json:"ExperimentalFogInCity"`
}
type configHealth struct {
	Effects struct {
		Berserk struct {
			DefaultDelay       int `json:"DefaultDelay"`
			DefaultResidueTime int `json:"DefaultResidueTime"`
			WorkingTime        int `json:"WorkingTime"`
		} `json:"Berserk"`
		BodyTemperature struct {
			DefaultBuildUpTime int `json:"DefaultBuildUpTime"`
			DefaultResidueTime int `json:"DefaultResidueTime"`
			LoopTime           int `json:"LoopTime"`
		} `json:"BodyTemperature"`
		BreakPart struct {
			BulletHitProbability struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            float64 `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"BulletHitProbability"`
			DefaultDelay       int `json:"DefaultDelay"`
			DefaultResidueTime int `json:"DefaultResidueTime"`
			FallingProbability struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            int     `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"FallingProbability"`
			HealExperience     int  `json:"HealExperience"`
			OfflineDurationMax int  `json:"OfflineDurationMax"`
			OfflineDurationMin int  `json:"OfflineDurationMin"`
			RemovePrice        int  `json:"RemovePrice"`
			RemovedAfterDeath  bool `json:"RemovedAfterDeath"`
		} `json:"BreakPart"`
		ChronicStaminaFatigue struct {
			EnergyRate         float64 `json:"EnergyRate"`
			EnergyRatePerStack float64 `json:"EnergyRatePerStack"`
			TicksEvery         int     `json:"TicksEvery"`
			WorkingTime        int     `json:"WorkingTime"`
		} `json:"ChronicStaminaFatigue"`
		Contusion struct {
			Dummy int `json:"Dummy"`
		} `json:"Contusion"`
		Dehydration struct {
			BleedingHealth            float64 `json:"BleedingHealth"`
			BleedingLifeTime          int     `json:"BleedingLifeTime"`
			BleedingLoopTime          int     `json:"BleedingLoopTime"`
			DamageOnStrongDehydration int     `json:"DamageOnStrongDehydration"`
			DefaultDelay              int     `json:"DefaultDelay"`
			DefaultResidueTime        int     `json:"DefaultResidueTime"`
			StrongDehydrationLoopTime int     `json:"StrongDehydrationLoopTime"`
		} `json:"Dehydration"`
		Disorientation struct {
			Dummy int `json:"Dummy"`
		} `json:"Disorientation"`
		Exhaustion struct {
			Damage             int `json:"Damage"`
			DamageLoopTime     int `json:"DamageLoopTime"`
			DefaultDelay       int `json:"DefaultDelay"`
			DefaultResidueTime int `json:"DefaultResidueTime"`
		} `json:"Exhaustion"`
		Existence struct {
			DestroyedStomachEnergyTimeFactor    int     `json:"DestroyedStomachEnergyTimeFactor"`
			DestroyedStomachHydrationTimeFactor int     `json:"DestroyedStomachHydrationTimeFactor"`
			EnergyDamage                        float64 `json:"EnergyDamage"`
			EnergyLoopTime                      int     `json:"EnergyLoopTime"`
			HydrationDamage                     float64 `json:"HydrationDamage"`
			HydrationLoopTime                   int     `json:"HydrationLoopTime"`
		} `json:"Existence"`
		Flash struct {
			Dummy int `json:"Dummy"`
		} `json:"Flash"`
		Fracture struct {
			BulletHitProbability struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            float64 `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"BulletHitProbability"`
			DefaultDelay       int `json:"DefaultDelay"`
			DefaultResidueTime int `json:"DefaultResidueTime"`
			FallingProbability struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            int     `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"FallingProbability"`
			HealExperience     int  `json:"HealExperience"`
			OfflineDurationMax int  `json:"OfflineDurationMax"`
			OfflineDurationMin int  `json:"OfflineDurationMin"`
			RemovePrice        int  `json:"RemovePrice"`
			RemovedAfterDeath  bool `json:"RemovedAfterDeath"`
		} `json:"Fracture"`
		HeavyBleeding struct {
			DamageEnergy             float64 `json:"DamageEnergy"`
			DamageHealth             float64 `json:"DamageHealth"`
			DamageHealthDehydrated   float64 `json:"DamageHealthDehydrated"`
			DefaultDelay             float64 `json:"DefaultDelay"`
			DefaultResidueTime       float64 `json:"DefaultResidueTime"`
			EliteVitalityDuration    int     `json:"EliteVitalityDuration"`
			EnergyLoopTime           int     `json:"EnergyLoopTime"`
			HealExperience           int     `json:"HealExperience"`
			HealthLoopTime           int     `json:"HealthLoopTime"`
			HealthLoopTimeDehydrated int     `json:"HealthLoopTimeDehydrated"`
			LifeTimeDehydrated       int     `json:"LifeTimeDehydrated"`
			OfflineDurationMax       int     `json:"OfflineDurationMax"`
			OfflineDurationMin       int     `json:"OfflineDurationMin"`
			Probability              struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            float64 `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"Probability"`
			RemovePrice       int  `json:"RemovePrice"`
			RemovedAfterDeath bool `json:"RemovedAfterDeath"`
		} `json:"HeavyBleeding"`
		Intoxication struct {
			DamageHealth       int  `json:"DamageHealth"`
			DefaultDelay       int  `json:"DefaultDelay"`
			DefaultResidueTime int  `json:"DefaultResidueTime"`
			HealExperience     int  `json:"HealExperience"`
			HealthLoopTime     int  `json:"HealthLoopTime"`
			OfflineDurationMax int  `json:"OfflineDurationMax"`
			OfflineDurationMin int  `json:"OfflineDurationMin"`
			RemovePrice        int  `json:"RemovePrice"`
			RemovedAfterDeath  bool `json:"RemovedAfterDeath"`
		} `json:"Intoxication"`
		LightBleeding struct {
			DamageEnergy             float64 `json:"DamageEnergy"`
			DamageHealth             float64 `json:"DamageHealth"`
			DamageHealthDehydrated   float64 `json:"DamageHealthDehydrated"`
			DefaultDelay             float64 `json:"DefaultDelay"`
			DefaultResidueTime       float64 `json:"DefaultResidueTime"`
			EliteVitalityDuration    int     `json:"EliteVitalityDuration"`
			EnergyLoopTime           int     `json:"EnergyLoopTime"`
			HealExperience           int     `json:"HealExperience"`
			HealthLoopTime           int     `json:"HealthLoopTime"`
			HealthLoopTimeDehydrated int     `json:"HealthLoopTimeDehydrated"`
			LifeTimeDehydrated       int     `json:"LifeTimeDehydrated"`
			OfflineDurationMax       int     `json:"OfflineDurationMax"`
			OfflineDurationMin       int     `json:"OfflineDurationMin"`
			Probability              struct {
				B            float64 `json:"B"`
				FunctionType string  `json:"FunctionType"`
				K            float64 `json:"K"`
				Threshold    float64 `json:"Threshold"`
			} `json:"Probability"`
			RemovePrice       int  `json:"RemovePrice"`
			RemovedAfterDeath bool `json:"RemovedAfterDeath"`
		} `json:"LightBleeding"`
		LowEdgeHealth struct {
			DefaultDelay       float64 `json:"DefaultDelay"`
			DefaultResidueTime int     `json:"DefaultResidueTime"`
			StartCommonHealth  int     `json:"StartCommonHealth"`
		} `json:"LowEdgeHealth"`
		MedEffect struct {
			DrinkStartDelay      int     `json:"DrinkStartDelay"`
			DrugsStartDelay      int     `json:"DrugsStartDelay"`
			FoodStartDelay       int     `json:"FoodStartDelay"`
			LoopTime             float64 `json:"LoopTime"`
			MedKitStartDelay     int     `json:"MedKitStartDelay"`
			MedicalStartDelay    int     `json:"MedicalStartDelay"`
			StartDelay           int     `json:"StartDelay"`
			StimulatorStartDelay int     `json:"StimulatorStartDelay"`
		} `json:"MedEffect"`
		MildMusclePain struct {
			GymEffectivity     float64 `json:"GymEffectivity"`
			OfflineDurationMax int     `json:"OfflineDurationMax"`
			OfflineDurationMin int     `json:"OfflineDurationMin"`
			TraumaChance       int     `json:"TraumaChance"`
		} `json:"MildMusclePain"`
		Pain struct {
			HealExperience int `json:"HealExperience"`
			TremorDelay    int `json:"TremorDelay"`
		} `json:"Pain"`
		PainKiller struct {
			Dummy int `json:"Dummy"`
		} `json:"PainKiller"`
		RadExposure struct {
			Damage         int `json:"Damage"`
			DamageLoopTime int `json:"DamageLoopTime"`
		} `json:"RadExposure"`
		Regeneration struct {
			BodyHealth struct {
				Chest struct {
					Value float64 `json:"Value"`
				} `json:"Chest"`
				Head struct {
					Value float64 `json:"Value"`
				} `json:"Head"`
				LeftArm struct {
					Value float64 `json:"Value"`
				} `json:"LeftArm"`
				LeftLeg struct {
					Value float64 `json:"Value"`
				} `json:"LeftLeg"`
				RightArm struct {
					Value float64 `json:"Value"`
				} `json:"RightArm"`
				RightLeg struct {
					Value float64 `json:"Value"`
				} `json:"RightLeg"`
				Stomach struct {
					Value float64 `json:"Value"`
				} `json:"Stomach"`
			} `json:"BodyHealth"`
			Energy     int `json:"Energy"`
			Hydration  int `json:"Hydration"`
			Influences struct {
				Fracture struct {
					EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage"`
					HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage"`
					HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage"`
				} `json:"Fracture"`
				HeavyBleeding struct {
					EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage"`
					HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage"`
					HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage"`
				} `json:"HeavyBleeding"`
				Intoxication struct {
					EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage"`
					HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage"`
					HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage"`
				} `json:"Intoxication"`
				LightBleeding struct {
					EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage"`
					HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage"`
					HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage"`
				} `json:"LightBleeding"`
				RadExposure struct {
					EnergySlowDownPercentage    int `json:"EnergySlowDownPercentage"`
					HealthSlowDownPercentage    int `json:"HealthSlowDownPercentage"`
					HydrationSlowDownPercentage int `json:"HydrationSlowDownPercentage"`
				} `json:"RadExposure"`
			} `json:"Influences"`
			LoopTime                int `json:"LoopTime"`
			MinimumHealthPercentage int `json:"MinimumHealthPercentage"`
		} `json:"Regeneration"`
		SandingScreen struct {
			Dummy int `json:"Dummy"`
		} `json:"SandingScreen"`
		SevereMusclePain struct {
			GymEffectivity     int `json:"GymEffectivity"`
			OfflineDurationMax int `json:"OfflineDurationMax"`
			OfflineDurationMin int `json:"OfflineDurationMin"`
			TraumaChance       int `json:"TraumaChance"`
		} `json:"SevereMusclePain"`
		Stimulator struct {
			BuffLoopTime float64 `json:"BuffLoopTime"`
			Buffs        struct {
				BuffsAdrenaline []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"BuffsAdrenaline"`
				BuffsGoldenStarBalm []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"BuffsGoldenStarBalm"`
				BuffsPropital []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"BuffsPropital"`
				BuffsSJ1TGLabs []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"BuffsSJ1TGLabs"`
				BuffsSJ6TGLabs []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"BuffsSJ6TGLabs"`
				BuffsZagustin []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"BuffsZagustin"`
				Buffs3BTG []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_3bTG"`
				BuffsAHF1M []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_AHF1M"`
				BuffsAntidote []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_Antidote"`
				BuffsBodyTemperature []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_BodyTemperature"`
				BuffsKultistsToxin []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_KultistsToxin"`
				BuffsL1 []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_L1"`
				BuffsMULE []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_MULE"`
				BuffsMeldonin []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_Meldonin"`
				BuffsObdolbos []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        float64 `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_Obdolbos"`
				BuffsObdolbos2 []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_Obdolbos2"`
				BuffsP22 []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_P22"`
				BuffsPNB []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_PNB"`
				BuffsPerfotoran []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_Perfotoran"`
				BuffsSJ12TGLabs []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_SJ12_TGLabs"`
				BuffsTrimadol []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_Trimadol"`
				BuffsDrinkAquamari []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_aquamari"`
				BuffsDrinkHotrod []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_hotrod"`
				BuffsDrinkJack []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_jack"`
				BuffsDrinkJuiceArmy []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_drink_juice_army"`
				BuffsDrinkMaxenergy []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_maxenergy"`
				BuffsDrinkMilk []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_milk"`
				BuffsDrinkMoonshine []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_moonshine"`
				BuffsDrinkPurewater []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_purewater"`
				BuffsDrinkTarcola []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_drink_tarcola"`
				BuffsDrinkVodka []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_vodka"`
				BuffsDrinkVodkaBAD []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        float64 `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_drink_vodka_BAD"`
				BuffsDrinkWater []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_drink_water"`
				BuffsFoodAlyonka []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_alyonka"`
				BuffsFoodBeer []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_beer"`
				BuffsFoodBorodinskiye []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_food_borodinskiye"`
				BuffsFoodCondensedMilk []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_condensed_milk"`
				BuffsFoodEmelya []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_food_emelya"`
				BuffsFoodMayonez []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_mayonez"`
				BuffsFoodMre []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_food_mre"`
				BuffsFoodSlippers []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_slippers"`
				BuffsFoodSugar []struct {
					AbsoluteValue bool   `json:"AbsoluteValue"`
					BuffType      string `json:"BuffType"`
					Chance        int    `json:"Chance"`
					Delay         int    `json:"Delay"`
					Duration      int    `json:"Duration"`
					SkillName     string `json:"SkillName"`
					Value         int    `json:"Value"`
				} `json:"Buffs_food_sugar"`
				BuffsHultafors []struct {
					AbsoluteValue bool          `json:"AbsoluteValue"`
					AppliesTo     []interface{} `json:"AppliesTo"`
					BuffType      string        `json:"BuffType"`
					Chance        int           `json:"Chance"`
					Delay         int           `json:"Delay"`
					Duration      int           `json:"Duration"`
					SkillName     string        `json:"SkillName"`
					Value         int           `json:"Value"`
				} `json:"Buffs_hultafors"`
				BuffsKnife []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"Buffs_knife"`
				BuffsMeleeBleed []struct {
					AbsoluteValue bool          `json:"AbsoluteValue"`
					AppliesTo     []interface{} `json:"AppliesTo"`
					BuffType      string        `json:"BuffType"`
					Chance        float64       `json:"Chance"`
					Delay         int           `json:"Delay"`
					Duration      int           `json:"Duration"`
					SkillName     string        `json:"SkillName"`
					Value         int           `json:"Value"`
				} `json:"Buffs_melee_bleed"`
				BuffsMeleeBlunt []struct {
					AbsoluteValue bool     `json:"AbsoluteValue"`
					AppliesTo     []string `json:"AppliesTo"`
					BuffType      string   `json:"BuffType"`
					Chance        int      `json:"Chance"`
					Delay         int      `json:"Delay"`
					Duration      int      `json:"Duration"`
					SkillName     string   `json:"SkillName"`
					Value         int      `json:"Value"`
				} `json:"Buffs_melee_blunt"`
				BuffseTGchange []struct {
					AbsoluteValue bool    `json:"AbsoluteValue"`
					BuffType      string  `json:"BuffType"`
					Chance        int     `json:"Chance"`
					Delay         int     `json:"Delay"`
					Duration      int     `json:"Duration"`
					SkillName     string  `json:"SkillName"`
					Value         float64 `json:"Value"`
				} `json:"BuffseTGchange"`
			} `json:"Buffs"`
		} `json:"Stimulator"`
		Stun struct {
			Dummy int `json:"Dummy"`
		} `json:"Stun"`
		Tremor struct {
			DefaultDelay       int     `json:"DefaultDelay"`
			DefaultResidueTime float64 `json:"DefaultResidueTime"`
		} `json:"Tremor"`
		Wound struct {
			ThresholdMax int `json:"ThresholdMax"`
			ThresholdMin int `json:"ThresholdMin"`
			WorkingTime  int `json:"WorkingTime"`
		} `json:"Wound"`
	} `json:"Effects"`
	Falling struct {
		DamagePerMeter int `json:"DamagePerMeter"`
		SafeHeight     int `json:"SafeHeight"`
	} `json:"Falling"`
	HealPrice struct {
		EnergyPointPrice    int `json:"EnergyPointPrice"`
		HealthPointPrice    int `json:"HealthPointPrice"`
		HydrationPointPrice int `json:"HydrationPointPrice"`
		TrialLevels         int `json:"TrialLevels"`
		TrialRaids          int `json:"TrialRaids"`
	} `json:"HealPrice"`
	ProfileHealthSettings struct {
		BodyPartsSettings struct {
			Chest struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"Chest"`
			Head struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"Head"`
			LeftArm struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"LeftArm"`
			LeftLeg struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"LeftLeg"`
			RightArm struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"RightArm"`
			RightLeg struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"RightLeg"`
			Stomach struct {
				Default                      int     `json:"Default"`
				Maximum                      int     `json:"Maximum"`
				Minimum                      int     `json:"Minimum"`
				OverDamageReceivedMultiplier float64 `json:"OverDamageReceivedMultiplier"`
			} `json:"Stomach"`
		} `json:"BodyPartsSettings"`
		DefaultStimulatorBuff string `json:"DefaultStimulatorBuff"`
		HealthFactorsSettings struct {
			Energy struct {
				Default int `json:"Default"`
				Maximum int `json:"Maximum"`
				Minimum int `json:"Minimum"`
			} `json:"Energy"`
			Hydration struct {
				Default int `json:"Default"`
				Maximum int `json:"Maximum"`
				Minimum int `json:"Minimum"`
			} `json:"Hydration"`
			Poisoning struct {
				Default int `json:"Default"`
				Maximum int `json:"Maximum"`
				Minimum int `json:"Minimum"`
			} `json:"Poisoning"`
			Radiation struct {
				Default int `json:"Default"`
				Maximum int `json:"Maximum"`
				Minimum int `json:"Minimum"`
			} `json:"Radiation"`
			Temperature struct {
				Default float64 `json:"Default"`
				Maximum int     `json:"Maximum"`
				Minimum int     `json:"Minimum"`
			} `json:"Temperature"`
		} `json:"HealthFactorsSettings"`
	} `json:"ProfileHealthSettings"`
}
type configInertia struct {
	AverageRotationFrameSpan           int     `json:"AverageRotationFrameSpan"`
	BaseJumpPenalty                    float64 `json:"BaseJumpPenalty"`
	BaseJumpPenaltyDuration            float64 `json:"BaseJumpPenaltyDuration"`
	CrouchSpeedAccelerationRange       Vector3 `json:"CrouchSpeedAccelerationRange"`
	DiagonalTime                       Vector3 `json:"DiagonalTime"`
	DurationPower                      float64 `json:"DurationPower"`
	ExitMovementStateSpeedThreshold    Vector3 `json:"ExitMovementStateSpeedThreshold"`
	FallThreshold                      float64 `json:"FallThreshold"`
	InertiaBackwardCoef                Vector3 `json:"InertiaBackwardCoef"`
	InertiaLimits                      Vector3 `json:"InertiaLimits"`
	InertiaLimitsStep                  float64 `json:"InertiaLimitsStep"`
	InertiaTiltCurveMax                Vector3 `json:"InertiaTiltCurveMax"`
	InertiaTiltCurveMin                Vector3 `json:"InertiaTiltCurveMin"`
	MaxMovementAccelerationRangeRight  Vector3 `json:"MaxMovementAccelerationRangeRight"`
	MaxTimeWithoutInput                Vector3 `json:"MaxTimeWithoutInput"`
	MinDirectionBlendTime              float64 `json:"MinDirectionBlendTime"`
	MinMovementAccelerationRangeRight  Vector3 `json:"MinMovementAccelerationRangeRight"`
	MoveTimeRange                      Vector3 `json:"MoveTimeRange"`
	PenaltyPower                       float64 `json:"PenaltyPower"`
	PreSprintAccelerationLimits        Vector3 `json:"PreSprintAccelerationLimits"`
	ProneDirectionAccelerationRange    Vector3 `json:"ProneDirectionAccelerationRange"`
	ProneSpeedAccelerationRange        Vector3 `json:"ProneSpeedAccelerationRange"`
	SideTime                           Vector3 `json:"SideTime"`
	SpeedInertiaAfterJump              Vector3 `json:"SpeedInertiaAfterJump"`
	SpeedLimitAfterFallMax             Vector3 `json:"SpeedLimitAfterFallMax"`
	SpeedLimitAfterFallMin             Vector3 `json:"SpeedLimitAfterFallMin"`
	SpeedLimitDurationMax              Vector3 `json:"SpeedLimitDurationMax"`
	SpeedLimitDurationMin              Vector3 `json:"SpeedLimitDurationMin"`
	SprintAccelerationLimits           Vector3 `json:"SprintAccelerationLimits"`
	SprintBrakeInertia                 Vector3 `json:"SprintBrakeInertia"`
	SprintSpeedInertiaCurveMax         Vector3 `json:"SprintSpeedInertiaCurveMax"`
	SprintSpeedInertiaCurveMin         Vector3 `json:"SprintSpeedInertiaCurveMin"`
	SprintTransitionMotionPreservation Vector3 `json:"SprintTransitionMotionPreservation"`
	TiltAcceleration                   Vector3 `json:"TiltAcceleration"`
	TiltInertiaMaxSpeed                Vector3 `json:"TiltInertiaMaxSpeed"`
	TiltMaxSideBackSpeed               Vector3 `json:"TiltMaxSideBackSpeed"`
	TiltStartSideBackSpeed             Vector3 `json:"TiltStartSideBackSpeed"`
	WalkInertia                        Vector3 `json:"WalkInertia"`
	WeaponFlipSpeed                    Vector3 `json:"WeaponFlipSpeed"`
}
type configInsurance struct {
	MaxStorageTimeInHour int `json:"MaxStorageTimeInHour"`
}
type configItemsCommonSettings struct {
	ItemRemoveAfterInterruptionTime int `json:"ItemRemoveAfterInterruptionTime"`
}
type configMalfunction struct {
	AllowMalfForBots             bool    `json:"AllowMalfForBots"`
	AmmoFeedWt                   float64 `json:"AmmoFeedWt"`
	AmmoJamWt                    float64 `json:"AmmoJamWt"`
	AmmoMalfChanceMult           int     `json:"AmmoMalfChanceMult"`
	AmmoMisfireWt                float64 `json:"AmmoMisfireWt"`
	DurFeedWt                    float64 `json:"DurFeedWt"`
	DurHardSlideMaxWt            float64 `json:"DurHardSlideMaxWt"`
	DurHardSlideMinWt            float64 `json:"DurHardSlideMinWt"`
	DurJamWt                     float64 `json:"DurJamWt"`
	DurMisfireWt                 float64 `json:"DurMisfireWt"`
	DurRangeToIgnoreMalfs        Vector3 `json:"DurRangeToIgnoreMalfs"`
	DurSoftSlideWt               float64 `json:"DurSoftSlideWt"`
	IdleToOutSpeedMultOnMalf     int     `json:"IdleToOutSpeedMultOnMalf"`
	MagazineMalfChanceMult       int     `json:"MagazineMalfChanceMult"`
	MalfRepairHardSlideMult      int     `json:"MalfRepairHardSlideMult"`
	MalfRepairOneHandBrokenMult  float64 `json:"MalfRepairOneHandBrokenMult"`
	MalfRepairTwoHandsBrokenMult float64 `json:"MalfRepairTwoHandsBrokenMult"`
	OutToIdleSpeedMultForPistol  int     `json:"OutToIdleSpeedMultForPistol"`
	OverheatFeedWt               float64 `json:"OverheatFeedWt"`
	OverheatHardSlideMaxWt       float64 `json:"OverheatHardSlideMaxWt"`
	OverheatHardSlideMinWt       float64 `json:"OverheatHardSlideMinWt"`
	OverheatJamWt                float64 `json:"OverheatJamWt"`
	OverheatSoftSlideWt          float64 `json:"OverheatSoftSlideWt"`
	ShowGlowAttemptsCount        int     `json:"ShowGlowAttemptsCount"`
	TimeToQuickdrawPistol        int     `json:"TimeToQuickdrawPistol"`
}
type ConfigMastering struct {
	Level2    int      `json:"Level2"`
	Level3    int      `json:"Level3"`
	Name      string   `json:"Name"`
	Templates []string `json:"Templates"`
}
type configFractureCausedBy struct {
	B            float64 `json:"B"`
	FunctionType string  `json:"FunctionType"`
	K            float64 `json:"K"`
	Threshold    float64 `json:"Threshold"`
}

type ragFairMaxActiveOffer struct {
	Count int     `json:"count"`
	From  float64 `json:"from"`
	To    float64 `json:"to"`
}
type ragFairMaxSumRarity struct {
	Common    Value `json:"Common"`
	NotExist  Value `json:"Not_exist"`
	Rare      Value `json:"Rare"`
	Superrare Value `json:"Superrare"`
}
type configRagFair struct {
	ChangePriceCoef                         int                     `json:"ChangePriceCoef"`
	BalancerAveragePriceCoefficient         int                     `json:"balancerAveragePriceCoefficient"`
	BalancerMinPriceCount                   int                     `json:"balancerMinPriceCount"`
	BalancerRemovePriceCoefficient          int                     `json:"balancerRemovePriceCoefficient"`
	BalancerUserItemSaleCooldown            int                     `json:"balancerUserItemSaleCooldown"`
	BalancerUserItemSaleCooldownEnabled     bool                    `json:"balancerUserItemSaleCooldownEnabled"`
	CommunityItemTax                        int                     `json:"communityItemTax"`
	CommunityRequirementTax                 int                     `json:"communityRequirementTax"`
	CommunityTax                            int                     `json:"communityTax"`
	DelaySinceOfferAdd                      int                     `json:"delaySinceOfferAdd"`
	Enabled                                 bool                    `json:"enabled"`
	IncludePveTraderSales                   bool                    `json:"includePveTraderSales"`
	IsOnlyFoundInRaidAllowed                bool                    `json:"isOnlyFoundInRaidAllowed"`
	MaxActiveOfferCount                     []ragFairMaxActiveOffer `json:"maxActiveOfferCount"`
	MaxRenewOfferTimeInHour                 int                     `json:"maxRenewOfferTimeInHour"`
	MaxSumForDecreaseRatingPerOneSale       int                     `json:"maxSumForDecreaseRatingPerOneSale"`
	MaxSumForIncreaseRatingPerOneSale       int                     `json:"maxSumForIncreaseRatingPerOneSale"`
	MaxSumForRarity                         ragFairMaxSumRarity     `json:"maxSumForRarity"`
	MinUserLevel                            int                     `json:"minUserLevel"`
	OfferDurationTimeInHour                 int                     `json:"offerDurationTimeInHour"`
	OfferDurationTimeInHourAfterRemove      float64                 `json:"offerDurationTimeInHourAfterRemove"`
	OfferPriorityCost                       int                     `json:"offerPriorityCost"`
	PriceStabilizerEnabled                  bool                    `json:"priceStabilizerEnabled"`
	PriceStabilizerStartIntervalInHours     int                     `json:"priceStabilizerStartIntervalInHours"`
	PriorityTimeModifier                    int                     `json:"priorityTimeModifier"`
	RatingDecreaseCount                     float64                 `json:"ratingDecreaseCount"`
	RatingIncreaseCount                     float64                 `json:"ratingIncreaseCount"`
	RatingSumForDecrease                    int                     `json:"ratingSumForDecrease"`
	RatingSumForIncrease                    int                     `json:"ratingSumForIncrease"`
	RenewPricePerHour                       float64                 `json:"renewPricePerHour"`
	SellInOnePiece                          int                     `json:"sellInOnePiece"`
	UniqueBuyerTimeoutInDays                int                     `json:"uniqueBuyerTimeoutInDays"`
	YouSellOfferMaxStorageTimeInHour        int                     `json:"youSellOfferMaxStorageTimeInHour"`
	YourOfferDidNotSellMaxStorageTimeInHour int                     `json:"yourOfferDidNotSellMaxStorageTimeInHour"`
}

type configOverheat struct {
	AutoshotChance              float64 `json:"AutoshotChance"`
	AutoshotMinOverheat         int     `json:"AutoshotMinOverheat"`
	AutoshotPossibilityDuration int     `json:"AutoshotPossibilityDuration"`
	BarrelMoveMaxMult           int     `json:"BarrelMoveMaxMult"`
	BarrelMoveRndDuration       int     `json:"BarrelMoveRndDuration"`
	DurReduceMaxMult            int     `json:"DurReduceMaxMult"`
	DurReduceMinMult            int     `json:"DurReduceMinMult"`
	EnableSlideOnMaxOverheat    bool    `json:"EnableSlideOnMaxOverheat"`
	FirerateOverheatBorder      int     `json:"FirerateOverheatBorder"`
	FireratePitchMult           float64 `json:"FireratePitchMult"`
	FirerateReduceMaxMult       float64 `json:"FirerateReduceMaxMult"`
	FirerateReduceMinMult       float64 `json:"FirerateReduceMinMult"`
	FixSlideOverheat            int     `json:"FixSlideOverheat"`
	MaxCOIIncreaseMult          float64 `json:"MaxCOIIncreaseMult"`
	MaxMalfChance               int     `json:"MaxMalfChance"`
	MaxOverheat                 int     `json:"MaxOverheat"`
	MaxOverheatCoolCoef         int     `json:"MaxOverheatCoolCoef"`
	MaxWearOnMaxOverheat        float64 `json:"MaxWearOnMaxOverheat"`
	MaxWearOnOverheat           float64 `json:"MaxWearOnOverheat"`
	MinMalfChance               float64 `json:"MinMalfChance"`
	MinOverheat                 int     `json:"MinOverheat"`
	MinWearOnMaxOverheat        float64 `json:"MinWearOnMaxOverheat"`
	MinWearOnOverheat           float64 `json:"MinWearOnOverheat"`
	ModCoolFactor               int     `json:"ModCoolFactor"`
	ModHeatFactor               int     `json:"ModHeatFactor"`
	OverheatProblemsStart       int     `json:"OverheatProblemsStart"`
	OverheatWearLimit           float64 `json:"OverheatWearLimit"`
	StartSlideOverheat          int     `json:"StartSlideOverheat"`
}

type repairSettingsItemEnhancementSettings struct {
	DamageReduction        PriceModifier `json:"DamageReduction"`
	MalfunctionProtections PriceModifier `json:"MalfunctionProtections"`
	WeaponSpread           PriceModifier `json:"WeaponSpread"`
}
type repairSettingsRepairStrategies struct {
	Armor    RepairStrategy `json:"Armor"`
	Firearms RepairStrategy `json:"Firearms"`
}
type configRepairSettings struct {
	ItemEnhancementSettings  repairSettingsItemEnhancementSettings `json:"ItemEnhancementSettings"`
	MinimumLevelToApplyBuff  int                                   `json:"MinimumLevelToApplyBuff"`
	RepairStrategies         repairSettingsRepairStrategies        `json:"RepairStrategies"`
	ArmorClassDivisor        int                                   `json:"armorClassDivisor"`
	DurabilityPointCostArmor float64                               `json:"durabilityPointCostArmor"`
	DurabilityPointCostGuns  float64                               `json:"durabilityPointCostGuns"`
}

type requirementReferencesAlpinist struct {
	Count          int    `json:"Count"`
	Id             string `json:"Id"`
	RequiredSlot   string `json:"RequiredSlot"`
	Requirement    string `json:"Requirement"`
	RequirementTip string `json:"RequirementTip"`
}
type configRequirementReferences struct {
	Alpinist []requirementReferencesAlpinist `json:"Alpinist"`
}

type configRestrictionsInRaid struct {
	MaxInLobby int    `json:"MaxInLobby"`
	MaxInRaid  int    `json:"MaxInRaid"`
	TemplateId string `json:"TemplateId"`
}
type configSkillsSettings struct {
	AdvancedModding []interface{} `json:"AdvancedModding"`
	AimDrills       struct {
		WeaponShotAction float64 `json:"WeaponShotAction"`
	} `json:"AimDrills"`
	Assault struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"Assault"`
	AttachedLauncher []interface{} `json:"AttachedLauncher"`
	Attention        struct {
		DependentSkillRatios []struct {
			Ratio   float64 `json:"Ratio"`
			SkillId string  `json:"SkillId"`
		} `json:"DependentSkillRatios"`
		ExamineWithInstruction int     `json:"ExamineWithInstruction"`
		FindActionFalse        float64 `json:"FindActionFalse"`
		FindActionTrue         float64 `json:"FindActionTrue"`
	} `json:"Attention"`
	Auctions              []interface{} `json:"Auctions"`
	Barter                []interface{} `json:"Barter"`
	BearAksystems         []interface{} `json:"BearAksystems"`
	BearAssaultoperations []interface{} `json:"BearAssaultoperations"`
	BearAuthority         []interface{} `json:"BearAuthority"`
	BearHeavycaliber      []interface{} `json:"BearHeavycaliber"`
	BearRawpower          []interface{} `json:"BearRawpower"`
	BotReload             []interface{} `json:"BotReload"`
	BotSound              []interface{} `json:"BotSound"`
	Charisma              struct {
		BonusSettings struct {
			EliteBonusSettings struct {
				FenceStandingLossDiscount float64 `json:"FenceStandingLossDiscount"`
				RepeatableQuestExtraCount int     `json:"RepeatableQuestExtraCount"`
				ScavCaseDiscount          float64 `json:"ScavCaseDiscount"`
			} `json:"EliteBonusSettings"`
			LevelBonusSettings struct {
				HealthRestoreDiscount         float64 `json:"HealthRestoreDiscount"`
				HealthRestoreTraderDiscount   float64 `json:"HealthRestoreTraderDiscount"`
				InsuranceDiscount             float64 `json:"InsuranceDiscount"`
				InsuranceTraderDiscount       float64 `json:"InsuranceTraderDiscount"`
				PaidExitDiscount              float64 `json:"PaidExitDiscount"`
				RepeatableQuestChangeDiscount float64 `json:"RepeatableQuestChangeDiscount"`
			} `json:"LevelBonusSettings"`
		} `json:"BonusSettings"`
		Counters struct {
			InsuranceCost struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"insuranceCost"`
			RepairCost struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"repairCost"`
			RepeatableQuestCompleteCount struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"repeatableQuestCompleteCount"`
			RestoredHealthCost struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"restoredHealthCost"`
			ScavCaseCost struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"scavCaseCost"`
		} `json:"Counters"`
		SkillProgressAtn float64 `json:"SkillProgressAtn"`
		SkillProgressInt float64 `json:"SkillProgressInt"`
		SkillProgressPer float64 `json:"SkillProgressPer"`
	} `json:"Charisma"`
	Cleanoperations []interface{} `json:"Cleanoperations"`
	CovertMovement  struct {
		MovementAction float64 `json:"MovementAction"`
	} `json:"CovertMovement"`
	Crafting struct {
		CraftTimeReductionPerLevel      float64 `json:"CraftTimeReductionPerLevel"`
		CraftingCycleHours              int     `json:"CraftingCycleHours"`
		CraftingPointsToInteligence     int     `json:"CraftingPointsToInteligence"`
		EliteExtraProductions           int     `json:"EliteExtraProductions"`
		PointsPerCraftingCycle          int     `json:"PointsPerCraftingCycle"`
		PointsPerUniqueCraftCycle       float64 `json:"PointsPerUniqueCraftCycle"`
		ProductionTimeReductionPerLevel float64 `json:"ProductionTimeReductionPerLevel"`
		UniqueCraftsPerCycle            int     `json:"UniqueCraftsPerCycle"`
	} `json:"Crafting"`
	DMR struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"DMR"`
	Endurance struct {
		DependentSkillRatios []struct {
			Ratio   float64 `json:"Ratio"`
			SkillId string  `json:"SkillId"`
		} `json:"DependentSkillRatios"`
		GainPerFatigueStack float64 `json:"GainPerFatigueStack"`
		MovementAction      float64 `json:"MovementAction"`
		QTELevelMultipliers struct {
			Field1 struct {
				Multiplier float64 `json:"Multiplier"`
			} `json:"10"`
			Field2 struct {
				Multiplier float64 `json:"Multiplier"`
			} `json:"25"`
			Field3 struct {
				Multiplier float64 `json:"Multiplier"`
			} `json:"50"`
		} `json:"QTELevelMultipliers"`
		SprintAction float64 `json:"SprintAction"`
	} `json:"Endurance"`
	FieldMedicine []interface{} `json:"FieldMedicine"`
	FirstAid      []interface{} `json:"FirstAid"`
	Freetrading   []interface{} `json:"Freetrading"`
	HMG           []interface{} `json:"HMG"`
	Health        struct {
		SkillProgress float64 `json:"SkillProgress"`
	} `json:"Health"`
	HeavyVests struct {
		BluntThroughputDamageHVestsReducePerLevel float64 `json:"BluntThroughputDamageHVestsReducePerLevel"`
		BuffMaxCount                              int     `json:"BuffMaxCount"`
		BuffSettings                              struct {
			CommonBuffChanceLevelBonus        float64 `json:"CommonBuffChanceLevelBonus"`
			CommonBuffMinChanceValue          float64 `json:"CommonBuffMinChanceValue"`
			CurrentDurabilityLossToRemoveBuff float64 `json:"CurrentDurabilityLossToRemoveBuff"`
			MaxDurabilityLossToRemoveBuff     float64 `json:"MaxDurabilityLossToRemoveBuff"`
			RareBuffChanceCoff                float64 `json:"RareBuffChanceCoff"`
			ReceivedDurabilityMaxPercent      int     `json:"ReceivedDurabilityMaxPercent"`
		} `json:"BuffSettings"`
		Counters struct {
			ArmorDurability struct {
				Divisor float64 `json:"divisor"`
				Points  int     `json:"points"`
			} `json:"armorDurability"`
		} `json:"Counters"`
		MoveSpeedPenaltyReductionHVestsReducePerLevel  float64 `json:"MoveSpeedPenaltyReductionHVestsReducePerLevel"`
		RicochetChanceHVestsCurrentDurabilityThreshold float64 `json:"RicochetChanceHVestsCurrentDurabilityThreshold"`
		RicochetChanceHVestsEliteLevel                 float64 `json:"RicochetChanceHVestsEliteLevel"`
		RicochetChanceHVestsMaxDurabilityThreshold     float64 `json:"RicochetChanceHVestsMaxDurabilityThreshold"`
		WearAmountRepairHVestsReducePerLevel           float64 `json:"WearAmountRepairHVestsReducePerLevel"`
		WearChanceRepairHVestsReduceEliteLevel         float64 `json:"WearChanceRepairHVestsReduceEliteLevel"`
	} `json:"HeavyVests"`
	HideoutManagement struct {
		ConsumptionReductionPerLevel float64 `json:"ConsumptionReductionPerLevel"`
		EliteSlots                   struct {
			AirFilteringUnit struct {
				Container int `json:"Container"`
				Slots     int `json:"Slots"`
			} `json:"AirFilteringUnit"`
			BitcoinFarm struct {
				Container int `json:"Container"`
				Slots     int `json:"Slots"`
			} `json:"BitcoinFarm"`
			Generator struct {
				Container int `json:"Container"`
				Slots     int `json:"Slots"`
			} `json:"Generator"`
			WaterCollector struct {
				Container int `json:"Container"`
				Slots     int `json:"Slots"`
			} `json:"WaterCollector"`
		} `json:"EliteSlots"`
		SkillBoostPercent         int `json:"SkillBoostPercent"`
		SkillPointsPerAreaUpgrade int `json:"SkillPointsPerAreaUpgrade"`
		SkillPointsPerCraft       int `json:"SkillPointsPerCraft"`
		SkillPointsRate           struct {
			AirFilteringUnit struct {
				PointsGained  int `json:"PointsGained"`
				ResourceSpent int `json:"ResourceSpent"`
			} `json:"AirFilteringUnit"`
			Generator struct {
				PointsGained  int `json:"PointsGained"`
				ResourceSpent int `json:"ResourceSpent"`
			} `json:"Generator"`
			SolarPower struct {
				PointsGained  int `json:"PointsGained"`
				ResourceSpent int `json:"ResourceSpent"`
			} `json:"SolarPower"`
			WaterCollector struct {
				PointsGained  int `json:"PointsGained"`
				ResourceSpent int `json:"ResourceSpent"`
			} `json:"WaterCollector"`
		} `json:"SkillPointsRate"`
	} `json:"HideoutManagement"`
	Immunity struct {
		HealthNegativeEffect   float64 `json:"HealthNegativeEffect"`
		ImmunityMiscEffects    int     `json:"ImmunityMiscEffects"`
		ImmunityPainKiller     int     `json:"ImmunityPainKiller"`
		ImmunityPoisonBuff     int     `json:"ImmunityPoisonBuff"`
		StimulatorNegativeBuff float64 `json:"StimulatorNegativeBuff"`
	} `json:"Immunity"`
	Intellect struct {
		Counters struct {
			ArmorDurability struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"armorDurability"`
			FirearmsDurability struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"firearmsDurability"`
			MeleeWeaponDurability struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"meleeWeaponDurability"`
		} `json:"Counters"`
		DependentSkillRatios []struct {
			Ratio   float64 `json:"Ratio"`
			SkillId string  `json:"SkillId"`
		} `json:"DependentSkillRatios"`
		ExamineAction              int     `json:"ExamineAction"`
		RepairAction               float64 `json:"RepairAction"`
		RepairPointsCostReduction  float64 `json:"RepairPointsCostReduction"`
		SkillProgress              float64 `json:"SkillProgress"`
		WearAmountReducePerLevel   float64 `json:"WearAmountReducePerLevel"`
		WearChanceReduceEliteLevel float64 `json:"WearChanceReduceEliteLevel"`
	} `json:"Intellect"`
	LMG        []interface{} `json:"LMG"`
	Launcher   []interface{} `json:"Launcher"`
	LightVests struct {
		BuffMaxCount int `json:"BuffMaxCount"`
		BuffSettings struct {
			CommonBuffChanceLevelBonus        float64 `json:"CommonBuffChanceLevelBonus"`
			CommonBuffMinChanceValue          float64 `json:"CommonBuffMinChanceValue"`
			CurrentDurabilityLossToRemoveBuff float64 `json:"CurrentDurabilityLossToRemoveBuff"`
			MaxDurabilityLossToRemoveBuff     float64 `json:"MaxDurabilityLossToRemoveBuff"`
			RareBuffChanceCoff                float64 `json:"RareBuffChanceCoff"`
			ReceivedDurabilityMaxPercent      int     `json:"ReceivedDurabilityMaxPercent"`
		} `json:"BuffSettings"`
		Counters struct {
			ArmorDurability struct {
				Divisor float64 `json:"divisor"`
				Points  int     `json:"points"`
			} `json:"armorDurability"`
		} `json:"Counters"`
		MeleeDamageLVestsReducePerLevel               float64 `json:"MeleeDamageLVestsReducePerLevel"`
		MoveSpeedPenaltyReductionLVestsReducePerLevel float64 `json:"MoveSpeedPenaltyReductionLVestsReducePerLevel"`
		WearAmountRepairLVestsReducePerLevel          float64 `json:"WearAmountRepairLVestsReducePerLevel"`
		WearChanceRepairLVestsReduceEliteLevel        float64 `json:"WearChanceRepairLVestsReduceEliteLevel"`
	} `json:"LightVests"`
	Lockpicking []interface{} `json:"Lockpicking"`
	MagDrills   struct {
		MagazineCheckAction    float64 `json:"MagazineCheckAction"`
		RaidLoadedAmmoAction   float64 `json:"RaidLoadedAmmoAction"`
		RaidUnloadedAmmoAction float64 `json:"RaidUnloadedAmmoAction"`
	} `json:"MagDrills"`
	Melee struct {
		BuffSettings struct {
			CommonBuffChanceLevelBonus   int     `json:"CommonBuffChanceLevelBonus"`
			CommonBuffMinChanceValue     int     `json:"CommonBuffMinChanceValue"`
			RareBuffChanceCoff           int     `json:"RareBuffChanceCoff"`
			ReceivedDurabilityMaxPercent float64 `json:"ReceivedDurabilityMaxPercent"`
		} `json:"BuffSettings"`
	} `json:"Melee"`
	Memory struct {
		AnySkillUp    int     `json:"AnySkillUp"`
		SkillProgress float64 `json:"SkillProgress"`
	} `json:"Memory"`
	Metabolism struct {
		DecreaseNegativeEffectDurationRate int `json:"DecreaseNegativeEffectDurationRate"`
		DecreasePoisonDurationRate         int `json:"DecreasePoisonDurationRate"`
		EnergyRecoveryRate                 int `json:"EnergyRecoveryRate"`
		HydrationRecoveryRate              int `json:"HydrationRecoveryRate"`
		IncreasePositiveEffectDurationRate int `json:"IncreasePositiveEffectDurationRate"`
	} `json:"Metabolism"`
	NightOps   []interface{} `json:"NightOps"`
	Perception struct {
		DependentSkillRatios []struct {
			Ratio   float64 `json:"Ratio"`
			SkillId string  `json:"SkillId"`
		} `json:"DependentSkillRatios"`
		OnlineAction float64 `json:"OnlineAction"`
		UniqueLoot   float64 `json:"UniqueLoot"`
	} `json:"Perception"`
	Pistol struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"Pistol"`
	ProneMovement []interface{} `json:"ProneMovement"`
	RecoilControl struct {
		RecoilAction        float64 `json:"RecoilAction"`
		RecoilBonusPerLevel float64 `json:"RecoilBonusPerLevel"`
	} `json:"RecoilControl"`
	Revolver struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"Revolver"`
	SMG    []interface{} `json:"SMG"`
	Search struct {
		FindAction   float64 `json:"FindAction"`
		SearchAction float64 `json:"SearchAction"`
	} `json:"Search"`
	Shadowconnections []interface{} `json:"Shadowconnections"`
	Shotgun           struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"Shotgun"`
	SilentOps         []interface{} `json:"SilentOps"`
	SkillProgressRate float64       `json:"SkillProgressRate"`
	Sniper            struct {
		WeaponChamberAction float64 `json:"WeaponChamberAction"`
		WeaponFixAction     float64 `json:"WeaponFixAction"`
		WeaponReloadAction  float64 `json:"WeaponReloadAction"`
		WeaponShotAction    float64 `json:"WeaponShotAction"`
	} `json:"Sniper"`
	Sniping  []interface{} `json:"Sniping"`
	Strength struct {
		DependentSkillRatios []struct {
			Ratio   float64 `json:"Ratio"`
			SkillId string  `json:"SkillId"`
		} `json:"DependentSkillRatios"`
		FistfightAction     float64 `json:"FistfightAction"`
		MovementActionMax   float64 `json:"MovementActionMax"`
		MovementActionMin   float64 `json:"MovementActionMin"`
		PushUpMax           float64 `json:"PushUpMax"`
		PushUpMin           float64 `json:"PushUpMin"`
		QTELevelMultipliers []struct {
			Level      int     `json:"Level"`
			Multiplier float64 `json:"Multiplier"`
		} `json:"QTELevelMultipliers"`
		SprintActionMax float64 `json:"SprintActionMax"`
		SprintActionMin float64 `json:"SprintActionMin"`
		ThrowAction     float64 `json:"ThrowAction"`
	} `json:"Strength"`
	StressResistance struct {
		HealthNegativeEffect float64 `json:"HealthNegativeEffect"`
		LowHPDuration        float64 `json:"LowHPDuration"`
	} `json:"StressResistance"`
	Surgery struct {
		SkillProgress float64 `json:"SkillProgress"`
		SurgeryAction int     `json:"SurgeryAction"`
	} `json:"Surgery"`
	Taskperformance []interface{} `json:"Taskperformance"`
	Throwing        struct {
		ThrowAction float64 `json:"ThrowAction"`
	} `json:"Throwing"`
	TroubleShooting struct {
		EliteAmmoChanceReduceMult       float64 `json:"EliteAmmoChanceReduceMult"`
		EliteDurabilityChanceReduceMult float64 `json:"EliteDurabilityChanceReduceMult"`
		EliteMagChanceReduceMult        float64 `json:"EliteMagChanceReduceMult"`
		MalfRepairSpeedBonusPerLevel    float64 `json:"MalfRepairSpeedBonusPerLevel"`
		SkillPointsPerMalfFix           float64 `json:"SkillPointsPerMalfFix"`
	} `json:"TroubleShooting"`
	UsecArsystems                 []interface{} `json:"UsecArsystems"`
	UsecDeepweaponmoddingSettings []interface{} `json:"UsecDeepweaponmodding_Settings"`
	UsecLongrangeopticsSettings   []interface{} `json:"UsecLongrangeoptics_Settings"`
	UsecNegotiations              []interface{} `json:"UsecNegotiations"`
	UsecTactics                   []interface{} `json:"UsecTactics"`
	Vitality                      struct {
		DamageTakenAction    float64 `json:"DamageTakenAction"`
		HealthNegativeEffect float64 `json:"HealthNegativeEffect"`
	} `json:"Vitality"`
	WeaponModding                  []interface{} `json:"WeaponModding"`
	WeaponSkillProgressRate        int           `json:"WeaponSkillProgressRate"`
	WeaponSkillRecoilBonusPerLevel float64       `json:"WeaponSkillRecoilBonusPerLevel"`
	WeaponTreatment                struct {
		BuffMaxCount int `json:"BuffMaxCount"`
		BuffSettings struct {
			CommonBuffChanceLevelBonus        float64 `json:"CommonBuffChanceLevelBonus"`
			CommonBuffMinChanceValue          float64 `json:"CommonBuffMinChanceValue"`
			CurrentDurabilityLossToRemoveBuff float64 `json:"CurrentDurabilityLossToRemoveBuff"`
			MaxDurabilityLossToRemoveBuff     float64 `json:"MaxDurabilityLossToRemoveBuff"`
			RareBuffChanceCoff                float64 `json:"RareBuffChanceCoff"`
			ReceivedDurabilityMaxPercent      int     `json:"ReceivedDurabilityMaxPercent"`
		} `json:"BuffSettings"`
		Counters struct {
			FirearmsDurability struct {
				Divisor int `json:"divisor"`
				Points  int `json:"points"`
			} `json:"firearmsDurability"`
		} `json:"Counters"`
		DurLossReducePerLevel                float64       `json:"DurLossReducePerLevel"`
		Filter                               []interface{} `json:"Filter"`
		SkillPointsPerRepair                 int           `json:"SkillPointsPerRepair"`
		WearAmountRepairGunsReducePerLevel   float64       `json:"WearAmountRepairGunsReducePerLevel"`
		WearChanceRepairGunsReduceEliteLevel float64       `json:"WearChanceRepairGunsReduceEliteLevel"`
	} `json:"WeaponTreatment"`
} //this needs to be broken down
type configSquadSettings struct {
	CountOfRequestsToOnePlayer int `json:"CountOfRequestsToOnePlayer"`
	SecondsForExpiredRequest   int `json:"SecondsForExpiredRequest"`
	SendRequestDelaySeconds    int `json:"SendRequestDelaySeconds"`
}
type configStamina struct {
	AimConsumptionByPose               Vector3 `json:"AimConsumptionByPose"`
	AimDrainRate                       float64 `json:"AimDrainRate"`
	AimRangeFinderDrainRate            float64 `json:"AimRangeFinderDrainRate"`
	AimingSpeedMultiplier              float64 `json:"AimingSpeedMultiplier"`
	BaseHoldBreathConsumption          int     `json:"BaseHoldBreathConsumption"`
	BaseOverweightLimits               Vector3 `json:"BaseOverweightLimits"`
	BaseRestorationRate                float64 `json:"BaseRestorationRate"`
	Capacity                           int     `json:"Capacity"`
	CrouchConsumption                  Vector3 `json:"CrouchConsumption"`
	ExhaustedMeleeDamageMultiplier     float64 `json:"ExhaustedMeleeDamageMultiplier"`
	ExhaustedMeleeSpeed                float64 `json:"ExhaustedMeleeSpeed"`
	FallDamageMultiplier               float64 `json:"FallDamageMultiplier"`
	FatigueAmountToCreateEffect        int     `json:"FatigueAmountToCreateEffect"`
	FatigueRestorationRate             float64 `json:"FatigueRestorationRate"`
	GrenadeHighThrow                   int     `json:"GrenadeHighThrow"`
	GrenadeLowThrow                    int     `json:"GrenadeLowThrow"`
	HandsCapacity                      int     `json:"HandsCapacity"`
	HandsRestoration                   float64 `json:"HandsRestoration"`
	HoldBreathStaminaMultiplier        Vector3 `json:"HoldBreathStaminaMultiplier"`
	JumpConsumption                    int     `json:"JumpConsumption"`
	OverweightConsumptionByPose        Vector3 `json:"OverweightConsumptionByPose"`
	OxygenCapacity                     int     `json:"OxygenCapacity"`
	OxygenRestoration                  int     `json:"OxygenRestoration"`
	PoseLevelConsumptionPerNotch       Vector3 `json:"PoseLevelConsumptionPerNotch"`
	PoseLevelDecreaseSpeed             Vector3 `json:"PoseLevelDecreaseSpeed"`
	PoseLevelIncreaseSpeed             Vector3 `json:"PoseLevelIncreaseSpeed"`
	ProneConsumption                   int     `json:"ProneConsumption"`
	RestorationMultiplierByPose        Vector3 `json:"RestorationMultiplierByPose"`
	SafeHeightOverweight               float64 `json:"SafeHeightOverweight"`
	SitToStandConsumption              int     `json:"SitToStandConsumption"`
	SoundRadius                        Vector3 `json:"SoundRadius"`
	SprintAccelerationLowerLimit       float64 `json:"SprintAccelerationLowerLimit"`
	SprintDrainRate                    int     `json:"SprintDrainRate"`
	SprintOverweightLimits             Vector3 `json:"SprintOverweightLimits"`
	SprintSensitivityLowerLimit        float64 `json:"SprintSensitivityLowerLimit"`
	SprintSpeedLowerLimit              float64 `json:"SprintSpeedLowerLimit"`
	StaminaExhaustionCausesJiggle      bool    `json:"StaminaExhaustionCausesJiggle"`
	StaminaExhaustionRocksCamera       bool    `json:"StaminaExhaustionRocksCamera"`
	StaminaExhaustionStartsBreathSound bool    `json:"StaminaExhaustionStartsBreathSound"`
	StandupConsumption                 Vector3 `json:"StandupConsumption"`
	TransitionSpeed                    Vector3 `json:"TransitionSpeed"`
	WalkConsumption                    Vector3 `json:"WalkConsumption"`
	WalkOverweightLimits               Vector3 `json:"WalkOverweightLimits"`
	WalkSpeedOverweightLimits          Vector3 `json:"WalkSpeedOverweightLimits"`
	WalkVisualEffectMultiplier         float64 `json:"WalkVisualEffectMultiplier"`
	WeaponFastSwitchConsumption        int     `json:"WeaponFastSwitchConsumption"`
}
type configStaminaDrain struct {
	LeftPlatoPoint  float64 `json:"LeftPlatoPoint"`
	LowerLeftPoint  float64 `json:"LowerLeftPoint"`
	LowerRightPoint float64 `json:"LowerRightPoint"`
	RightLimit      float64 `json:"RightLimit"`
	RightPlatoPoint int     `json:"RightPlatoPoint"`
	ZeroValue       int     `json:"ZeroValue"`
}
type configStaminaRestoration struct {
	LeftPlatoPoint  float64 `json:"LeftPlatoPoint"`
	LowerLeftPoint  float64 `json:"LowerLeftPoint"`
	LowerRightPoint float64 `json:"LowerRightPoint"`
	RightLimit      float64 `json:"RightLimit"`
	RightPlatoPoint float64 `json:"RightPlatoPoint"`
	ZeroValue       float64 `json:"ZeroValue"`
}

type armorClass struct {
	Resistance int `json:"resistance"`
}
type configArmor struct {
	Class []armorClass `json:"class"`
}

type tradingSettingsBuyoutRestrictions struct {
	MinDurability        float64 `json:"MinDurability"`
	MinFoodDrinkResource float64 `json:"MinFoodDrinkResource"`
	MinMedsResource      float64 `json:"MinMedsResource"`
}
type configTradingSettings struct {
	BuyoutRestrictions tradingSettingsBuyoutRestrictions `json:"BuyoutRestrictions"`
}
type configWeaponFastDrawSettings struct {
	HandShakeCurveFrequency            int     `json:"HandShakeCurveFrequency"`
	HandShakeCurveIntensity            int     `json:"HandShakeCurveIntensity"`
	HandShakeMaxDuration               int     `json:"HandShakeMaxDuration"`
	HandShakeTremorIntensity           int     `json:"HandShakeTremorIntensity"`
	WeaponFastSwitchMaxSpeedMult       int     `json:"WeaponFastSwitchMaxSpeedMult"`
	WeaponFastSwitchMinSpeedMult       float64 `json:"WeaponFastSwitchMinSpeedMult"`
	WeaponPistolFastSwitchMaxSpeedMult int     `json:"WeaponPistolFastSwitchMaxSpeedMult"`
	WeaponPistolFastSwitchMinSpeedMult float64 `json:"WeaponPistolFastSwitchMinSpeedMult"`
}
type configContent struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Root string `json:"root"`
}

type expHeal struct {
	ExpForEnergy    int `json:"expForEnergy"`
	ExpForHeal      int `json:"expForHeal"`
	ExpForHydration int `json:"expForHydration"`
}
type killCombo struct {
	Percent int `json:"percent"`
}
type expKill struct {
	BloodLossToLitre        float64     `json:"bloodLossToLitre"`
	BotExpOnDamageAllHealth int         `json:"botExpOnDamageAllHealth"`
	BotHeadShotMult         float64     `json:"botHeadShotMult"`
	Combo                   []killCombo `json:"combo"`
	LongShotDistance        int         `json:"longShotDistance"`
	PmcExpOnDamageAllHealth int         `json:"pmcExpOnDamageAllHealth"`
	PmcHeadShotMult         float64     `json:"pmcHeadShotMult"`
	VictimBotLevelExp       int         `json:"victimBotLevelExp"`
	VictimLevelExp          int         `json:"victimLevelExp"`
}
type levelTable struct {
	Exp int `json:"exp"`
}
type expLevel struct {
	ClanLevel   int          `json:"clan_level"`
	ExpTable    []levelTable `json:"exp_table"`
	Mastering1  int          `json:"mastering1"`
	Mastering2  int          `json:"mastering2"`
	SavageLevel int          `json:"savage_level"`
	TradeLevel  int          `json:"trade_level"`
}
type expMatchEnd struct {
	README                     string  `json:"README"`
	KilledMult                 int     `json:"killedMult"`
	LeftMult                   int     `json:"leftMult"`
	MiaMult                    int     `json:"miaMult"`
	MiaExpReward               int     `json:"mia_exp_reward"`
	RunnerMult                 float64 `json:"runnerMult"`
	RunnerExpReward            int     `json:"runner_exp_reward"`
	SurvivedMult               float64 `json:"survivedMult"`
	SurvivedExpRequirement     int     `json:"survived_exp_requirement"`
	SurvivedExpReward          int     `json:"survived_exp_reward"`
	SurvivedSecondsRequirement int     `json:"survived_seconds_requirement"`
}
type expLootAttempts struct {
	KExp float64 `json:"k_exp"`
}
type configExp struct {
	ExpForLockedDoorBreach int               `json:"expForLockedDoorBreach"`
	ExpForLockedDoorOpen   int               `json:"expForLockedDoorOpen"`
	Heal                   expHeal           `json:"heal"`
	Kill                   expKill           `json:"kill"`
	Level                  expLevel          `json:"level"`
	LootAttempts           []expLootAttempts `json:"loot_attempts"`
	MatchEnd               expMatchEnd       `json:"match_end"`
	TriggerMult            int               `json:"triggerMult"`
}

type configHandbook struct {
	DefaultCategory string `json:"defaultCategory"`
}

type ratingCategories struct {
	AvgEarnings       bool `json:"avgEarnings"`
	Experience        bool `json:"experience"`
	InventoryFullCost bool `json:"inventoryFullCost"`
	Kd                bool `json:"kd"`
	LongestShot       bool `json:"longestShot"`
	PmcKills          bool `json:"pmcKills"`
	RagFairStanding   bool `json:"ragFairStanding"`
	RaidCount         bool `json:"raidCount"`
	SurviveRatio      bool `json:"surviveRatio"`
	TimeOnline        bool `json:"timeOnline"`
}
type configRating struct {
	Categories    ratingCategories `json:"categories"`
	LevelRequired int              `json:"levelRequired"`
	Limit         int              `json:"limit"`
}

type tournamentCategories struct {
	Dogtags bool `json:"dogtags"`
}
type configTournament struct {
	Categories    tournamentCategories `json:"categories"`
	LevelRequired int                  `json:"levelRequired"`
	Limit         int                  `json:"limit"`
}

type Config struct {
	AimPunchMagnitude                  int                             `json:"AimPunchMagnitude"`
	Aiming                             configAiming                    `json:"Aiming"`
	ArmorMaterials                     map[string]configArmorMaterials `json:"ArmorMaterials"`
	AudioSettings                      configAudioSettings             `json:"AudioSettings"`
	AzimuthPanelShowsPlayerOrientation bool                            `json:"AzimuthPanelShowsPlayerOrientation"`
	Ballistic                          configBallistic                 `json:"Ballistic"`
	BaseCheckTime                      int                             `json:"BaseCheckTime"`
	BaseLoadTime                       float64                         `json:"BaseLoadTime"`
	BaseUnloadTime                     float64                         `json:"BaseUnloadTime"`
	BotsEnabled                        bool                            `json:"BotsEnabled"`
	BufferZone                         configBufferZone                `json:"BufferZone"`
	CoopSettings                       configCoopSettings              `json:"CoopSettings"`
	Customization                      configCustomization             `json:"Customization"`
	DiscardLimitsEnabled               bool                            `json:"DiscardLimitsEnabled"`
	EventType                          []string                        `json:"EventType"`
	FenceSettings                      configFenceSettings             `json:"FenceSettings"`
	FractureCausedByBulletHit          configFractureCausedBy          `json:"FractureCausedByBulletHit"`
	FractureCausedByFalling            configFractureCausedBy          `json:"FractureCausedByFalling"`
	GameSearchingTimeout               int                             `json:"GameSearchingTimeout"`
	GlobalItemPriceModifier            int                             `json:"GlobalItemPriceModifier"`
	GlobalLootChanceModifier           float64                         `json:"GlobalLootChanceModifier"`
	GraphicSettings                    configGraphicSettings           `json:"GraphicSettings"`
	HandsOverdamage                    float64                         `json:"HandsOverdamage"`
	Health                             configHealth                    `json:"Health"`
	Inertia                            configInertia                   `json:"Inertia"`
	Insurance                          configInsurance                 `json:"Insurance"`
	ItemsCommonSettings                configItemsCommonSettings       `json:"ItemsCommonSettings"`
	LegsOverdamage                     int                             `json:"LegsOverdamage"`
	LoadTimeSpeedProgress              int                             `json:"LoadTimeSpeedProgress"`
	Malfunction                        configMalfunction               `json:"Malfunction"`
	MarksmanAccuracy                   float64                         `json:"MarksmanAccuracy"`
	Mastering                          []ConfigMastering               `json:"Mastering"`
	MaxBotsAliveOnMap                  int                             `json:"MaxBotsAliveOnMap"`
	MaxLoyaltyLevelForAll              bool                            `json:"MaxLoyaltyLevelForAll"`
	Overheat                           configOverheat                  `json:"Overheat"`
	RagFair                            configRagFair                   `json:"RagFair"`
	RepairSettings                     configRepairSettings            `json:"RepairSettings"`
	RequirementReferences              configRequirementReferences     `json:"RequirementReferences"`
	RestrictionsInRaid                 []configRestrictionsInRaid      `json:"RestrictionsInRaid"`
	SavagePlayCooldown                 int                             `json:"SavagePlayCooldown"`
	SavagePlayCooldownDevelop          int                             `json:"SavagePlayCooldownDevelop"`
	SavagePlayCooldownNdaFree          int                             `json:"SavagePlayCooldownNdaFree"`
	SessionsToShowHotKeys              int                             `json:"SessionsToShowHotKeys"`
	SkillAtrophy                       bool                            `json:"SkillAtrophy"`
	SkillEnduranceWeightThreshold      float64                         `json:"SkillEnduranceWeightThreshold"`
	SkillExpPerLevel                   int                             `json:"SkillExpPerLevel"`
	SkillFatiguePerPoint               float64                         `json:"SkillFatiguePerPoint"`
	SkillFatigueReset                  int                             `json:"SkillFatigueReset"`
	SkillFreshEffectiveness            float64                         `json:"SkillFreshEffectiveness"`
	SkillFreshPoints                   int                             `json:"SkillFreshPoints"`
	SkillMinEffectiveness              float64                         `json:"SkillMinEffectiveness"`
	SkillPointsBeforeFatigue           int                             `json:"SkillPointsBeforeFatigue"`
	SkillsSettings                     configSkillsSettings            `json:"SkillsSettings"`
	SprintSpeed                        Vector3                         `json:"SprintSpeed"`
	SquadSettings                      configSquadSettings             `json:"SquadSettings"`
	Stamina                            configStamina                   `json:"Stamina"`
	StaminaDrain                       configStaminaDrain              `json:"StaminaDrain"`
	StaminaRestoration                 configStaminaRestoration        `json:"StaminaRestoration"`
	StomachOverdamage                  float64                         `json:"StomachOverdamage"`
	TODSkyDate                         string                          `json:"TODSkyDate"`
	TeamSearchingTimeout               int                             `json:"TeamSearchingTimeout"`
	TestValue                          int                             `json:"TestValue"`
	TimeBeforeDeploy                   int                             `json:"TimeBeforeDeploy"`
	TimeBeforeDeployLocal              int                             `json:"TimeBeforeDeployLocal"`
	TradingSetting                     int                             `json:"TradingSetting"`
	TradingSettings                    configTradingSettings           `json:"TradingSettings"`
	TradingUnlimitedItems              bool                            `json:"TradingUnlimitedItems"`
	UncheckOnShot                      bool                            `json:"UncheckOnShot"`
	WAVECOEFHIGH                       float64                         `json:"WAVE_COEF_HIGH"`
	WAVECOEFHORDE                      float64                         `json:"WAVE_COEF_HORDE"`
	WAVECOEFLOW                        float64                         `json:"WAVE_COEF_LOW"`
	WAVECOEFMID                        float64                         `json:"WAVE_COEF_MID"`
	WalkSpeed                          Vector3                         `json:"WalkSpeed"`
	WallContusionAbsorption            Vector3                         `json:"WallContusionAbsorption"`
	WeaponFastDrawSettings             configWeaponFastDrawSettings    `json:"WeaponFastDrawSettings"`
	WeaponSkillProgressRate            float64                         `json:"WeaponSkillProgressRate"`
	Armor                              configArmor                     `json:"armor"`
	Content                            configContent                   `json:"content"`
	Exp                                configExp                       `json:"exp"`
	Handbook                           configHandbook                  `json:"handbook"`
	Rating                             configRating                    `json:"rating"`
	TBaseLockpicking                   int                             `json:"t_base_lockpicking"`
	TBaseLooting                       int                             `json:"t_base_looting"`
	Tournament                         configTournament                `json:"tournament"`
}
