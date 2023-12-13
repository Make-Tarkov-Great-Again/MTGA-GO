package pkg

import (
	"MT-GO/data"
	"MT-GO/tools"
	"fmt"
	"log"
	"math"
	"strings"
)

func GetBrandName() *data.BrandName {
	brand := data.GetBrandName()
	if brand == nil {
		brand := data.BrandName{}
		brandName := data.GetServerConfig().BrandName
		if brandName != "" {
			brand["name"] = brandName
		}
		brand["name"] = "MTGA"
		brand.SetBrandName()
	}
	return brand
}

func SetGameConfig() {
	output := &data.GameConfig{
		Languages: data.GetLanguages(),
		NdaFree:   false,
		Taxonomy:  6,
		Backend: data.Backend{
			Lobby:     data.GetLobbyAddress(),
			Trading:   data.GetTradingAddress(),
			Messaging: data.GetMessageAddress(),
			Main:      data.GetMainAddress(),
			RagFair:   data.GetRagFairAddress(),
		},
		UseProtobuf:       false,
		ReportAvailable:   true,
		TwitchEventMember: false,
	}
	output.Set()
}

func GetGameConfig(sessionID string) (*data.GameConfig, error) {
	lang := "en"
	profile, err := data.GetProfileByUID(sessionID)
	if err != nil {
		return nil, err
	}
	if profile.Account.Lang != "" {
		lang = profile.Account.Lang
	}

	config := data.GetGameConfig()
	config.Aid = sessionID
	config.Lang = lang
	config.ActiveProfileID = sessionID
	config.UtcTime = tools.GetCurrentTimeInSeconds()
	config.TotalInGame = 0
	if profile.Character != nil {
		config.TotalInGame = profile.Character.Stats.Eft.TotalInGameTime
	}

	return &config, nil
}

func GetMainProfileList(sessionID string) []any {
	character, err := data.GetCharacterByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}
	profiles := make([]any, 0, 2)
	if character == nil || character.Info.Nickname == "" {
		log.Println("Character doesn't exist, begin creation")
		return profiles
	}

	playerScav := data.GetPlayerScav()
	playerScav.Info.RegistrationDate = int32(tools.GetCurrentTimeInSeconds())
	playerScav.AID = character.AID
	playerScav.ID = *character.Savage

	profiles = append(profiles, *playerScav, *character)
	return profiles
}

func GetMainAccountCustomization() []string {
	customization := data.GetCustomizations()
	output := make([]string, 0, customization.Len())
	customization.ForEach(func(id string, c *data.Customization) bool {
		if c.Props.Side == nil || len(c.Props.Side) == 0 {
			return true
		}
		output = append(output, id)
		return true
	})
	return output
}

func ValidateNickname(nickname string) *ResponseBody {
	if len(nickname) < 3 {
		body := ApplyResponseBody(nil)
		body.Err = 226
		body.Errmsg = "226 - "

		return body
	}

	if data.IsNicknameUnavailable(nickname) {
		body := ApplyResponseBody(nil)
		body.Err = 225
		body.Errmsg = "225 - "

		return body
	}

	return ApplyResponseBody(&NicknameValidate{
		Status: "ok",
	})
}

func CreateProfile(sessionId string, side string, nickname string, voiceId string, headId string) {
	profile, err := data.GetProfileByUID(sessionId)
	if err != nil {
		log.Fatalln(err)
	}

	edition, err := data.GetEditionByName(strings.ToLower(profile.Account.Edition))
	if err != nil {
		log.Fatalln(err)
	}

	var pmc data.Character[map[string]data.PlayerTradersInfo]
	if side == "Bear" {
		pmc = *edition.Bear
		profile.Storage.Suites = edition.Storage.Bear
	} else {
		pmc = *edition.Usec
		profile.Storage.Suites = edition.Storage.Usec
	}

	pmc.ID = sessionId
	pmc.AID = profile.Account.AID

	sid := tools.GenerateMongoID()
	pmc.Savage = &sid

	pmc.Info.Side = side
	pmc.Info.Nickname = nickname

	pmc.Info.LowerNickname = strings.ToLower(nickname)

	customization, err := data.GetCustomizationByID(voiceId)
	if err != nil {
		log.Fatalln(err)
	}
	pmc.Info.Voice = customization.Name

	time := int32(tools.GetCurrentTimeInSeconds())
	pmc.Info.RegistrationDate = time
	pmc.Health.UpdateTime = time

	pmc.Customization.Head = headId

	stats := &pmc.Stats.Eft
	stats.SessionCounters = nil
	stats.OverallCounters = data.Counter{Items: make([]any, 0)}
	stats.Aggressor = nil
	stats.DroppedItems = make([]any, 0)
	stats.FoundInRaidItems = make([]any, 0)
	stats.Victims = make([]any, 0)
	stats.CarriedQuestItems = make([]any, 0)
	stats.DamageHistory = map[string]any{
		"BodyParts":        []any{},
		"LethalDamage":     nil,
		"LethalDamagePart": "Head",
	}
	stats.SurvivorClass = "Unknown"

	commonSkills := make([]data.SkillsCommon, 0, len(pmc.Skills.Common))
	commonSkills = append(commonSkills, pmc.Skills.Common...)
	pmc.Skills.Common = commonSkills

	hideout := pmc.Hideout
	resizedAreas := make([]data.PlayerHideoutArea, 0, len(hideout.Areas))
	resizedAreas = append(resizedAreas, hideout.Areas...)
	hideout.Areas = resizedAreas

	hideout.Improvement = make(map[string]any)

	profile.Character = &pmc

	data.SetProfileCache(sessionId)
	profile.SaveProfile()
}

func SetChannelTemplate() {
	channel := data.GetChannels()
	channel.Template.Notifier.Server = data.GetLobbyIPandPort()
}

func GetChannel(sessionID string) data.Channel {
	channel := data.GetChannelsTemplate()
	channel.Notifier.ChannelID = sessionID
	channel.Notifier.NotifierServer = fmt.Sprintf(notiFormat, data.GetLobbyIPandPort(), sessionID)
	channel.Notifier.WS = fmt.Sprintf(wssFormat, data.GetWebSocketAddress(), sessionID)
	channel.SetChannel(sessionID)

	return channel
}

func GetChannelNotifier(sessionID string) (*data.Notifier, error) {
	channel, err := data.GetChannel(sessionID)
	if err != nil {
		return nil, err
	}

	if channel.Notifier == nil {
		return nil, fmt.Errorf(channelNotifierNotExist, sessionID)
	}

	return channel.Notifier, nil
}

func GetProfileStatuses(sessionID string) *ProfileStatuses {
	character, err := data.GetCharacterByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}
	return &ProfileStatuses{
		Profiles: []ProfileStatus{
			{
				ProfileID: *character.Savage,
				Status:    "Free",
			},
			{
				ProfileID: character.ID,
				Status:    "Free",
			},
		},
	}
}

func GetBuildsList(sessionID string) (*data.Builds, error) {
	storage, err := data.GetStorageByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	if storage.Builds != nil {
		return storage.Builds, nil
	}
	return nil, fmt.Errorf(storageBuildNotExist, sessionID)
}

func GetQuestList(sessionID string) ([]any, error) {
	character, err := data.GetCharacterByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}
	quests, err := data.GetQuestsAvailableToPlayer(*character)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return quests, nil
}

func GetMainPrices() *data.SupplyData {
	supplyData := data.GetSupplyData()
	supplyData.SupplyNextTime = data.SetResupplyTimer()
	return supplyData
}

func GetInsuranceCosts(sessionID string, traders []string, items []string) (map[string]map[string]int32, error) {
	output := make(map[string]map[string]int32)
	character, err := data.GetCharacterByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	invCache, err := data.GetInventoryCacheByID(sessionID)
	if err != nil {
		return nil, err
	}

	traderCache, err := data.GetTraderCacheByID(character.ID)
	if err != nil {
		return nil, err
	} else if traderCache.Insurances == nil {
		traderCache.Insurances = make(map[string]*data.Insurances)
	}

	for _, tid := range traders {
		trader, err := data.GetTraderByUID(tid)
		if err != nil {
			return nil, err
		}

		changed := int8(0)

		trader.SetTraderLoyaltyLevel(character)
		traderInsurance := traderCache.Insurances[tid]

		if traderInsurance != nil && traderInsurance.LoyaltyLevel != character.TradersInfo[tid].LoyaltyLevel {
			traderInsurance.LoyaltyLevel = character.TradersInfo[tid].LoyaltyLevel
			traderInsurance.PriceCoef = trader.Base.LoyaltyLevels[character.TradersInfo[tid].LoyaltyLevel].InsurancePriceCoef
			changed = 1
		} else {
			traderInsurance = &data.Insurances{
				LoyaltyLevel: character.TradersInfo[tid].LoyaltyLevel,
				PriceCoef:    trader.Base.LoyaltyLevels[character.TradersInfo[tid].LoyaltyLevel].InsurancePriceCoef,
				Items:        make(map[string]int32),
			}
		}

		output[tid] = make(map[string]int32)

		for _, itemID := range items {
			itemTPL := character.Inventory.Items[*invCache.GetIndexOfItemByID(itemID)].TPL
			if value, ok := data.IsItemBlacklist(itemTPL); ok && value == "node" {
				continue
			}

			item, ok := traderInsurance.Items[itemTPL]
			if ok && changed == 0 {
				output[tid][itemTPL] = item
				continue
			}

			itemPrice, _ := data.GetPriceByID(itemTPL)
			insuranceCost := int32(math.Round(float64(itemPrice) * 0.3))
			if traderInsurance.PriceCoef > 0 {
				insuranceCost *= int32(1 - traderInsurance.PriceCoef/100)
			}

			if item != insuranceCost {
				item = insuranceCost
			}

			//TODO: continue with cache
			output[tid][itemTPL] = item
		}
	}

	return output, nil
}

const (
	storageBuildNotExist    = "Storage builds for %s does not exist"
	channelNotifierNotExist = "Channel.Notifier for %s does not exist"
	notiFormat              = "%s/push/notifier/get/%s"
	wssFormat               = "%s/push/notifier/getwebsocket/%s"
)

type ProfileStatuses struct {
	MaxPVECountExceeded bool            `json:"maxPveCountExceeded"`
	Profiles            []ProfileStatus `json:"profiles"`
}

type ProfileStatus struct {
	ProfileID    string `json:"profileid"`
	ProfileToken any    `json:"profileToken"`
	Status       string `json:"status"`
	SID          string `json:"sid"`
	IP           string `json:"ip"`
	Port         int    `json:"port"`
}

type NicknameValidate struct {
	Status string `json:"status"`
}
