package hndlr

// #region Items moving

type QuestAccept struct{}

// #endregion

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

type ProfileCreateRequest struct {
	Side     string `json:"side"`
	Nickname string `json:"nickname"`
	HeadID   string `json:"headId"`
	VoiceID  string `json:"voiceId"`
}

type KeepAlive struct {
	Msg     string `json:"msg"`
	UtcTime int64  `json:"utc_time"`
}
