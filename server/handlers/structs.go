package handlers

// #region Items moving

type QuestAccept struct{}

// #endregion

type SupplyData struct {
	SupplyNextTime  int               `json:"supplyNextTime"`
	Prices          map[string]*int32 `json:"prices"`
	CurrencyCourses CurrencyCourses   `json:"currencyCourses"`
}

type CurrencyCourses struct {
	RUB int32 `json:"5449016a4bdc2d6f028b456f"`
	EUR int32 `json:"569668774bdc2da2298b4568"`
	DOL int32 `json:"5696686a4bdc2da3298b456a"`
}

type Version struct {
	IsValid       bool   `json:"isValid"`
	LatestVersion string `json:"latestVersion"`
}

type ServerListing struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type CurrentGroup struct {
	Squad []any `json:"squad"`
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

type Notifier struct {
	Server         string `json:"server"`
	ChannelID      string `json:"channel_id"`
	URL            string `json:"url"`
	NotifierServer string `json:"notifierServer"`
	WS             string `json:"ws"`
}

type Channel struct {
	Status         string   `json:"status"`
	Notifier       Notifier `json:"notifier"`
	NotifierServer string   `json:"notifierServer"`
}

type ProfileCreateRequest struct {
	Side     string `json:"side"`
	Nickname string `json:"nickname"`
	HeadID   string `json:"headId"`
	VoiceID  string `json:"voiceId"`
}

type NicknameValidate struct {
	Status string `json:"status"`
}

type KeepAlive struct {
	Msg     string `json:"msg"`
	UtcTime int64  `json:"utc_time"`
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

type DialogView struct {
	Type     int8   `json:"type"`
	DialogID string `json:"dialogId"`
	Limit    int8   `json:"limit"`
	Time     int64  `json:"time"`
}

type FriendRequestMailbox struct {
	Err  int   `json:"err"`
	Data []any `json:"data"`
}
