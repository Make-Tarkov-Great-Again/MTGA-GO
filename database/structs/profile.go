package structs

type Profile struct {
	Account   *Account
	Character *PlayerTemplate
	Storage   *Storage
	Dialogue  map[string]*Dialogue
}

type Storage struct {
	ID        string        `json:"_id"`
	Suites    []string      `json:"suites"`
	Builds    Builds        `json:"builds"`
	Insurance []interface{} `json:"insurance"`
	Mailbox   []Mail        `json:"mailbox"`
}

type Builds struct{}

type Insurance struct {
	ScheduledTime  int           `json:"scheduledTime"`
	TraderId       string        `json:"traderId"`
	MessageContent Message       `json:"messageContent"`
	Items          []interface{} `json:"items"`
}

type Mail struct {
	Type     string     `json:"type,omitempty"`
	EventId  string     `json:"eventId,omitempty"`
	DialogId string     `json:"dialogId,omitempty"`
	Message  []Dialogue `json:"message,omitempty"`
}

type Dialogue struct {
	AttachmentsNew int       `json:"attachmentsNew,omitempty"`
	New            int       `json:"new,omitempty"`
	Type           int       `json:"type,omitempty"`
	Users          []string  `json:"users,omitempty"`
	Pinned         bool      `json:"pinned,omitempty"`
	Messages       []Message `json:"messages,omitempty"`
	ID             string    `json:"_id,omitempty"`
}

type Message struct {
	ID              string        `json:"_id,omitempty"`
	UID             string        `json:"uid,omitempty"`
	Type            string        `json:"type,omitempty"`
	DT              int           `json:"dt,omitempty"`
	TemplateId      string        `json:"templateId,omitempty"`
	Text            string        `json:"text,omitempty"`
	RewardCollected bool          `json:"rewardCollected,omitempty"`
	HasRewards      bool          `json:"hasRewards,omitempty"`
	Items           []interface{} `json:"items,omitempty"`
	MaxStorageTime  int           `json:"maxStorageTime,omitempty"`
	SystemData      struct {
		Date     int    `json:"date,omitempty"`
		Time     int    `json:"time,omitempty"`
		Location string `json:"location,omitempty"`
	} `json:"systemData,omitempty"`
	ProfileChangeEvents []interface{} `json:"profileChangeEvents,omitempty"`
}

type MessageContent struct {
	TemplateId     string `json:"templateId"`
	Type           int    `json:"type"`
	MaxStorageTime int    `json:"maxStorageTime"`
}
