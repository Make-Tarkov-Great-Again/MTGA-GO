package pkg

import (
	"MT-GO/data"
	"MT-GO/tools"
	"fmt"
	"log"
	"math"
	"strings"
)

func GetBrandName() map[string]string {
	config := data.GetServerConfig()
	brand := make(map[string]string)
	if config.BrandName != "" {
		brand["name"] = config.BrandName
	} else {
		brand["name"] = "MTGA"
	}

	return brand
}

var gameConfig *GameConfig

func SetGameConfig() {
	gameConfig = &GameConfig{
		Languages: data.GetLanguages(),
		NdaFree:   false,
		Taxonomy:  6,
		Backend: Backend{
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
}

func GetGameConfig(sessionID string) (*GameConfig, error) {
	lang := "en"
	profile, err := data.GetProfileByUID(sessionID)
	if err != nil {
		return nil, err
	}
	if profile.Account.Lang != "" {
		lang = profile.Account.Lang
	}

	config := *gameConfig
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
	character := data.GetCharacterByID(sessionID)
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
	output := make([]string, 0, len(customization))
	for id, c := range customization {
		if c.Props.Side == nil || len(c.Props.Side) == 0 {
			continue
		}
		output = append(output, id)
	}
	return output
}

func ValidateNickname(nickname string) *ResponseBody {
	if len(nickname) == 0 {
		body := ApplyResponseBody(nil)
		body.Err = 226
		body.Errmsg = "226 - "

		return body
	}

	_, ok := data.Nicknames[nickname]
	if ok {
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

	edition := data.GetEditionByName("Edge Of Darkness")
	if edition == nil {
		log.Fatalln("[MainProfileCreate] Edition is nil, this ain't good fella!")
	}
	var pmc data.Character

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

	if customization, err := data.GetCustomizationByID(voiceId); err != nil {
		log.Fatalln(err)
	} else {
		pmc.Info.Voice = customization.Name
	}

	time := int32(tools.GetCurrentTimeInSeconds())
	pmc.Info.RegistrationDate = time
	pmc.Health.UpdateTime = time

	pmc.Customization.Head = headId

	stats := &pmc.Stats.Eft
	stats.SessionCounters = nil
	stats.OverallCounters = map[string]any{"Items": []any{}}
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

	hideout := &pmc.Hideout

	resizedAreas := make([]data.PlayerHideoutArea, 0, len(hideout.Areas))
	resizedAreas = append(resizedAreas, hideout.Areas...)
	hideout.Areas = resizedAreas

	hideout.Improvement = make(map[string]any)

	profile.Character = &pmc

	data.SetProfileCache(sessionId)
	profile.SaveProfile()
}

var channels = map[string]Channel{}

var templateChannel = &Channel{
	Status: "ok",
	Notifier: &Notifier{
		Server:         "",
		ChannelID:      "",
		URL:            "",
		NotifierServer: "",
		WS:             "",
	},
	NotifierServer: "",
}

func SetChannelTemplate() {
	templateChannel.Notifier.Server = data.GetLobbyIPandPort()
}

func GetChannel(sessionID string) Channel {
	channel := *templateChannel
	channel.Notifier.ChannelID = sessionID
	channel.Notifier.NotifierServer = fmt.Sprintf(notiFormat, data.GetLobbyIPandPort(), sessionID)
	channel.Notifier.WS = fmt.Sprintf(wssFormat, data.GetWebSocketAddress(), sessionID)
	channels[sessionID] = channel

	return channel
}

func GetChannelNotifier(sessionID string) (*Notifier, error) {
	channel, ok := channels[sessionID]
	if !ok {
		return nil, fmt.Errorf(channelNotExist, sessionID)
	}

	if channel.Notifier == nil {
		return nil, fmt.Errorf(channelNotifierNotExist, sessionID)
	}

	return channel.Notifier, nil
}

func GetProfileStatuses(sessionID string) *ProfileStatuses {
	character := data.GetCharacterByID(sessionID)
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
	quests, err := data.GetCharacterByID(sessionID).GetQuestsAvailableToPlayer()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return quests, nil
}

var supplyData *SupplyData

func GetMainPrices() *SupplyData {
	if supplyData != nil {
		supplyData.SupplyNextTime = data.SetResupplyTimer()
		return supplyData
	}

	prices := data.GetPrices()
	nextResupply := data.SetResupplyTimer()

	supplyData = &SupplyData{
		SupplyNextTime: nextResupply,
		Prices:         prices,
		CurrencyCourses: CurrencyCourses{
			RUB: prices[*GetCurrencyByName("RUB")],
			EUR: prices[*GetCurrencyByName("EUR")],
			DOL: prices[*GetCurrencyByName("USD")],
		},
	}
	return supplyData
}

func GetInsuranceCosts(sessionID string, traders []string, items []string) (map[string]map[string]int32, error) {
	output := make(map[string]map[string]int32)
	character := data.GetCharacterByID(sessionID)

	invCache, err := data.GetInventoryCacheByID(sessionID)
	if err != nil {
		return nil, err
	}

	traderCache, err := data.GetTraderCacheByID(character.ID)
	if err != nil {
		return nil, err
	}

	for _, tid := range traders {
		trader, err := data.GetTraderByUID(tid)
		if err != nil {
			return nil, err
		}

		changed := int8(0)

		trader.SetTraderLoyaltyLevel(character)
		if traderCache.Insurances[tid] != nil && traderCache.Insurances[tid].LoyaltyLevel != character.TradersInfo[tid].LoyaltyLevel {
			traderCache.Insurances[tid].LoyaltyLevel = character.TradersInfo[tid].LoyaltyLevel
			traderCache.Insurances[tid].PriceCoef = trader.Base.LoyaltyLevels[character.TradersInfo[tid].LoyaltyLevel].InsurancePriceCoef
			changed = 1
		} else {
			traderCache.Insurances[tid] = &data.Insurances{
				LoyaltyLevel: character.TradersInfo[tid].LoyaltyLevel,
				PriceCoef:    trader.Base.LoyaltyLevels[character.TradersInfo[tid].LoyaltyLevel].InsurancePriceCoef,
				Items:        make(map[string]int32),
			}
		}

		output[tid] = make(map[string]int32)

		for _, itemID := range items {
			itemTPL := character.Inventory.Items[*invCache.GetIndexOfItemByID(itemID)].TPL
			if _, ok := traderCache.Insurances[tid].Items[itemTPL]; ok && changed == 0 {
				output[tid][itemTPL] = traderCache.Insurances[tid].Items[itemTPL]
				continue
			}

			itemPrice, err := data.GetPriceByID(itemTPL)
			if err != nil {
				return nil, err
			}
			insuranceCost := int32(math.Round(float64(itemPrice) * 0.3))
			if traderCache.Insurances[tid].PriceCoef > 0 {
				insuranceCost *= int32(1 - traderCache.Insurances[tid].PriceCoef/100)
			}

			if traderCache.Insurances[tid].Items[itemTPL] != insuranceCost {
				traderCache.Insurances[tid].Items[itemTPL] = insuranceCost
			}

			//TODO: continue with cache
			output[tid][itemTPL] = traderCache.Insurances[tid].Items[itemTPL]
		}
	}

	return output, nil
}

const (
	storageBuildNotExist    = "Storage builds for %s does not exist"
	channelNotifierNotExist = "Channel.Notifier for %s does not exist"
	channelNotExist         = "Channel for %s does not exist"
	notiFormat              = "%s/push/notifier/get/%s"
	wssFormat               = "%s/push/notifier/getwebsocket/%s"
	localeRoute             = "/client/locale/"
	itemsRoute              = "/client/items"
	customizationRoute      = "/client/customization"
	globalsRoute            = "/client/globals"
	mainSettingsRoute       = "/client/settings"
)

type traderInsuranceInfo struct {
	LoyaltyLevel int8
	PriceCoef    int16
}

type SupplyData struct {
	SupplyNextTime  int              `json:"supplyNextTime"`
	Prices          map[string]int32 `json:"prices"`
	CurrencyCourses CurrencyCourses  `json:"currencyCourses"`
}

type CurrencyCourses struct {
	RUB int32 `json:"5449016a4bdc2d6f028b456f"`
	EUR int32 `json:"569668774bdc2da2298b4568"`
	DOL int32 `json:"5696686a4bdc2da3298b456a"`
}

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

type Backend struct {
	Lobby     string `json:"Lobby"`
	Trading   string `json:"Trading"`
	Messaging string `json:"Messaging"`
	Main      string `json:"Main"`
	RagFair   string `json:"RagFair"`
}

type GameConfig struct {
	Aid               string            `json:"aid"`
	Lang              string            `json:"lang"`
	Languages         map[string]string `json:"languages"`
	NdaFree           bool              `json:"ndaFree"`
	Taxonomy          int               `json:"taxonomy"`
	ActiveProfileID   string            `json:"activeProfileId"`
	Backend           Backend           `json:"backend"`
	UseProtobuf       bool              `json:"useProtobuf"`
	UtcTime           int64             `json:"utc_time"`
	TotalInGame       int               `json:"totalInGame"`
	ReportAvailable   bool              `json:"reportAvailable"`
	TwitchEventMember bool              `json:"twitchEventMember"`
}

type Notifier struct {
	Server         string `json:"server"`
	ChannelID      string `json:"channel_id"`
	URL            string `json:"url"`
	NotifierServer string `json:"notifierServer"`
	WS             string `json:"ws"`
}

type Channel struct {
	Status         string    `json:"status"`
	Notifier       *Notifier `json:"notifier"`
	NotifierServer string    `json:"notifierServer"`
}
