package database

import (
	"MT-GO/tools"
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

// #region Profile getters

const profilesPath string = "user/profiles/"

var profiles = map[string]*Profile{}

func GetProfiles() map[string]*Profile {
	return profiles
}

func GetProfileByUID(uid string) *Profile {
	if profile, ok := profiles[uid]; ok {
		return profile
	}

	fmt.Println("No profile with UID ", uid, ", you stupid motherfucker")
	return nil
}

func GetAccountByUID(uid string) *Account {
	profile := GetProfileByUID(uid)
	if profile.Account != nil {
		return profile.Account
	}

	fmt.Println("Profile with UID ", uid, " does not have an account, how the fuck did you get here????!?!?!?!?!?")
	return nil
}

func GetStorageByUID(uid string) *Storage {
	if profile, ok := profiles[uid]; ok {
		return profile.Storage
	}

	fmt.Println("Profile with UID ", uid, " does not have a storage")
	return nil
}

func GetDialogueByUID(uid string) *map[string]interface{} {
	if profile, ok := profiles[uid]; ok {
		return &profile.Dialogue
	}

	fmt.Println("Profile with UID ", uid, " does not have dialogue")
	return nil
}

func GetCache(sid string) *Cache {
	profile, ok := profiles[sid]
	if ok {
		return &profile.Cache
	}
	return nil
}

// #endregion

// #region Profile setters

func setProfiles() map[string]*Profile {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return profiles
	}
	for _, user := range users {
		profile := &Profile{}
		userPath := filepath.Join(profilesPath, user)

		path := filepath.Join(userPath, "account.json")
		if tools.FileExist(path) {
			profile.Account = setAccount(path)
		}

		path = filepath.Join(userPath, "character.json")
		if tools.FileExist(path) {
			profile.Character = setCharacter(path)
			if profile.Character.Info.Nickname != "" {
				Nicknames[profile.Character.Info.Nickname] = struct{}{}
			}
		}

		path = filepath.Join(userPath, "storage.json")
		if tools.FileExist(path) {
			profile.Storage = setStorage(path)
		}

		path = filepath.Join(userPath, "dialogue.json")
		if tools.FileExist(path) {
			profile.Dialogue = setDialogue(path)
		}

		profiles[user] = profile
	}

	for _, profile := range profiles {
		profile.Cache = Cache{
			Quests: QuestCache{
				Index:  map[string]int8{},
				Quests: map[string]CharacterQuest{},
			},
			Traders: TraderCache{
				Index:         map[string]*AssortIndex{},
				Assorts:       map[string]*Assort{},
				LoyaltyLevels: map[string]int8{},
			},
		}

		if profile.Character.ID != "" {
			traders := GetTraders()
			for _, trader := range traders {
				if trader.Assort == nil {
					continue
				}
				trader.GetStrippedAssort(profile.Character)
			}
		}
	}

	return profiles
}

func setAccount(path string) *Account {
	output := new(Account)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setCharacter(path string) *Character {
	output := new(Character)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setStorage(path string) *Storage {
	output := new(Storage)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setDialogue(path string) map[string]interface{} {
	output := make(map[string]interface{})

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, &output)
	if err != nil {
		panic(err)
	}

	return output
}

// #endregion

// #region Profile save

func (profile Profile) SaveProfile() {
	sessionID := profile.Account.UID
	profileDirPath := filepath.Join(profilesPath, sessionID)
	if !tools.FileExist(profileDirPath) {
		err := os.Mkdir(profileDirPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	profile.Account.SaveAccount()
	profile.Character.Save(sessionID)
	SaveDialogue(sessionID, profile.Dialogue)
	SaveStorage(sessionID, *profile.Storage)
	fmt.Println()
	fmt.Println("Profile saved")
}

func (a *Account) SaveAccount() {
	accountFilePath := filepath.Join(profilesPath, a.UID, "account.json")

	err := tools.WriteToFile(accountFilePath, a)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account saved")
}

func SaveStorage(sessionID string, storage Storage) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage saved")
}

func SaveDialogue(sessionID string, dialogue map[string]interface{}) {
	dialogueFilePath := filepath.Join(profilesPath, sessionID, "dialogue.json")

	err := tools.WriteToFile(dialogueFilePath, dialogue)
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialogue saved")
}

// #endregion

// #region Profile structs

type Profile struct {
	Account   *Account
	Character *Character
	Storage   *Storage
	Dialogue  map[string]interface{}
	Cache     Cache
}

type Cache struct {
	Quests  QuestCache
	Traders TraderCache
}

type TraderCache struct {
	Index         map[string]*AssortIndex
	Assorts       map[string]*Assort
	LoyaltyLevels map[string]int8
}

type QuestCache struct {
	Index  map[string]int8
	Quests map[string]CharacterQuest
}

type CharacterQuest struct {
	QID            string
	StartTime      int
	Status         string
	StatusTimers   map[string]int
	AvailableAfter int `json:",omitempty"`
}

type Usernames map[string]string

var Nicknames = make(map[string]struct{})

type Account struct {
	AID                 int           `json:"aid"`
	UID                 string        `json:"uid"`
	Username            string        `json:"username"`
	Password            string        `json:"password"`
	Wipe                bool          `json:"wipe"`
	Edition             string        `json:"edition"`
	Friends             Friends       `json:"friends"`
	Matching            Matching      `json:"Matching"`
	FriendRequestInbox  []interface{} `json:"friendRequestInbox"`
	FriendRequestOutbox []interface{} `json:"friendRequestOutbox"`
	TarkovPath          string        `json:"tarkovPath"`
	Lang                string        `json:"lang"`
}
type Friends struct {
	Friends      []FriendRequest `json:"Friends"`
	Ignore       []string        `json:"Ignore"`
	InIgnoreList []string        `json:"InIgnoreList"`
}
type Matching struct {
	LookingForGroup bool `json:"LookingForGroup"`
}

type FriendRequest struct {
	ID      string               `json:"_id"`
	From    string               `json:"from"`
	To      string               `json:"to"`
	Date    int32                `json:"date"`
	Profile FriendRequestProfile `json:"profile"`
}

type FriendRequestProfile struct {
	ID   int32
	Info struct {
		Nickname       string         `json:"Nickname"`
		Side           string         `json:"Side"`
		Level          int8           `json:"Level"`
		MemberCategory MemberCategory `json:"MemberCategory"`
	}
}

type Character struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         InventoryInfo          `json:"Inventory"`
	Skills            PlayerSkills           `json:"Skills"`
	Stats             PlayerStats            `json:"Stats"`
	Encyclopedia      map[string]bool        `json:"Encyclopedia"`
	ConditionCounters ConditionCounters      `json:"ConditionCounters"`
	BackendCounters   map[string]interface{} `json:"BackendCounters"`
	InsuredItems      []InsuredItem          `json:"InsuredItems"`
	Hideout           PlayerHideout          `json:"Hideout"`
	Bonuses           []Bonus                `json:"Bonuses"`
	Notes             struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests       []CharacterQuest             `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo            `json:"RagfairInfo"`
	WishList     []string                     `json:"WishList"`
	TradersInfo  map[string]PlayerTradersInfo `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []interface{} `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type PlayerTradersInfo struct {
	Unlocked bool    `json:"unlocked"`
	Disabled bool    `json:"disabled"`
	SalesSum float32 `json:"salesSum"`
	Standing float32 `json:"standing"`
}

type PlayerRagfairInfo struct {
	Rating          float64       `json:"rating"`
	IsRatingGrowing bool          `json:"isRatingGrowing"`
	Offers          []interface{} `json:"offers"`
}

type PlayerHideoutArea struct {
	Type                  int           `json:"type"`
	Level                 int           `json:"level"`
	Active                bool          `json:"active"`
	PassiveBonusesEnabled bool          `json:"passiveBonusesEnabled"`
	CompleteTime          int           `json:"completeTime"`
	Constructing          bool          `json:"constructing"`
	Slots                 []interface{} `json:"slots"`
	LastRecipe            string        `json:"lastRecipe"`
}
type PlayerHideout struct {
	Production  map[string]interface{} `json:"Production"`
	Areas       []PlayerHideoutArea    `json:"Areas"`
	Improvement map[string]interface{} `json:"Improvement"`
	//Seed        int                    `json:"Seed"`
}

type ConditionCounters struct {
	Counters []interface{} `json:"Counters"`
}

type PlayerStats struct {
	Eft EftStats `json:"Eft"`
}

type StatCounters struct {
	Items []interface{} `json:"Items"`
}

type EftStats struct {
	SessionCounters        map[string]interface{} `json:"SessionCounters"`
	OverallCounters        map[string]interface{} `json:"OverallCounters"`
	SessionExperienceMult  int                    `json:"SessionExperienceMult"`
	ExperienceBonusMult    int                    `json:"ExperienceBonusMult"`
	TotalSessionExperience int                    `json:"TotalSessionExperience"`
	LastSessionDate        int                    `json:"LastSessionDate"`
	Aggressor              map[string]interface{} `json:"Aggressor"`
	DroppedItems           []interface{}          `json:"DroppedItems"`
	FoundInRaidItems       []interface{}          `json:"FoundInRaidItems"`
	Victims                []interface{}          `json:"Victims"`
	CarriedQuestItems      []interface{}          `json:"CarriedQuestItems"`
	DamageHistory          map[string]interface{} `json:"DamageHistory"`
	LastPlayerState        *float32               `json:"LastPlayerState"`
	TotalInGameTime        int                    `json:"TotalInGameTime"`
	SurvivorClass          string                 `json:"SurvivorClass"`
}

type SkillsCommon struct {
	ID                        string `json:"Id"`
	Progress                  int    `json:"Progress"`
	PointsEarnedDuringSession int    `json:"PointsEarnedDuringSession"`
	LastAccess                int64  `json:"LastAccess"`
}
type SkillsMastering struct {
	ID       string `json:"Id"`
	Progress int    `json:"Progress"`
}

type PlayerSkills struct {
	Common    []SkillsCommon    `json:"Common"`
	Mastering []SkillsMastering `json:"Mastering"`
	Points    int               `json:"Points"`
}

type PlayerInfo struct {
	Nickname               string                 `json:"Nickname"`
	LowerNickname          string                 `json:"LowerNickname"`
	Side                   string                 `json:"Side"`
	Voice                  string                 `json:"Voice"`
	Level                  int                    `json:"Level"`
	Experience             int                    `json:"Experience"`
	RegistrationDate       int                    `json:"RegistrationDate"`
	GameVersion            string                 `json:"GameVersion"`
	AccountType            int                    `json:"AccountType"`
	MemberCategory         int                    `json:"MemberCategory"`
	LockedMoveCommands     bool                   `json:"lockedMoveCommands"`
	SavageLockTime         int                    `json:"SavageLockTime"`
	LastTimePlayedAsSavage int                    `json:"LastTimePlayedAsSavage"`
	Settings               map[string]interface{} `json:"Settings"`
	NicknameChangeDate     int                    `json:"NicknameChangeDate"`
	NeedWipeOptions        []interface{}          `json:"NeedWipeOptions"`
	LastCompletedWipe      struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedWipe"`
	LastCompletedEvent struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedEvent"`
	BannedState             bool          `json:"BannedState"`
	BannedUntil             int           `json:"BannedUntil"`
	IsStreamerModeAvailable bool          `json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool          `json:"SquadInviteRestriction"`
	Bans                    []interface{} `json:"Bans"`
}

type InfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int     `json:"Experience"`
	StandingForKill float32 `json:"StandingForKill"`
	AggressorBonus  float32 `json:"AggressorBonus"`
}

type PlayerCustomization struct {
	Head  string `json:"Head"`
	Body  string `json:"Body"`
	Feet  string `json:"Feet"`
	Hands string `json:"Hands"`
}

type InsuredItem struct {
	Tid    string `json:"tid"`
	ItemID string `json:"itemId"`
}

type Bonus struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	TemplateID string `json:"templateId"`
}

type InventoryInfo struct {
	Items              []interface{} `json:"items"`
	Equipment          string        `json:"equipment"`
	Stash              string        `json:"stash"`
	SortingTable       string        `json:"sortingTable"`
	QuestRaidItems     string        `json:"questRaidItems"`
	QuestStashItems    string        `json:"questStashItems"`
	FastPanel          interface{}   `json:"fastPanel"`
	HideoutAreaStashes interface{}   `json:"hideoutAreaStashes"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth   `json:"Hydration"`
	Energy      CurrMaxHealth   `json:"Energy"`
	Temperature CurrMaxHealth   `json:"Temperature"`
	BodyParts   BodyPartsHealth `json:"BodyParts"`
	UpdateTime  int             `json:"UpdateTime"`
}

type HealthOf struct {
	Health CurrMaxHealth `json:"Health"`
}

type CurrMaxHealth struct {
	Current float32 `json:"Current"`
	Maximum float32 `json:"Maximum"`
}

type BodyPartsHealth struct {
	Head     HealthOf `json:"Head"`
	Chest    HealthOf `json:"Chest"`
	Stomach  HealthOf `json:"Stomach"`
	LeftArm  HealthOf `json:"LeftArm"`
	RightArm HealthOf `json:"RightArm"`
	LeftLeg  HealthOf `json:"LeftLeg"`
	RightLeg HealthOf `json:"RightLeg"`
}

type Storage struct {
	//ID        string                 `json:"_id"`
	Suites    []string      `json:"suites"`
	Builds    Builds        `json:"builds"`
	Insurance []interface{} `json:"insurance"`
	Mailbox   []interface{} `json:"mailbox"`
}

type Builds struct {
	EquipmentBuilds []EquipmentBuild `json:"equipmentBuilds"`
	WeaponBuilds    []interface{}    `json:"weaponBuilds"`
}

type EquipmentBuild struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Root      string        `json:"root"`
	Items     []interface{} `json:"items"`
	Type      string        `json:"type"`
	FastPanel []interface{} `json:"fastPanel"`
	BuildType string        `json:"buildType"`
}

// #endregion
