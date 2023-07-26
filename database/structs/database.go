package structs

type CoreStruct struct {
	PlayerTemplate PlayerTemplate
	ClientSettings ClientSettings
	ServerConfig   ServerConfig
	Globals        Globals
	Locations      Locations
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics MatchMetrics
}

type Locales struct {
	Locales   Locale
	Languages map[string]string
}

type Locale struct {
	CH   LocaleData `json:"ch"`
	CZ   LocaleData `json:"cz"`
	EN   LocaleData `json:"en"`
	FR   LocaleData `json:"fr"`
	GE   LocaleData `json:"ge"`
	HU   LocaleData `json:"hu"`
	IT   LocaleData `json:"it"`
	JP   LocaleData `json:"jp"`
	KR   LocaleData `json:"kr"`
	PL   LocaleData `json:"pl"`
	PO   LocaleData `json:"po"`
	SK   LocaleData `json:"sk"`
	ES   LocaleData `json:"es"`
	ESMX LocaleData `json:"es-mx"`
	TU   LocaleData `json:"tu"`
	RU   LocaleData `json:"ru"`
}

type LocaleData struct {
	Locale map[string]string
	Menu   map[string]string
}

type Handbook struct {
	Categories [87]HandbookCategory `json:"Categories"`
	Items      [2819]HandbookItem   `json:"Items"`
}

type HandbookCategory struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Icon     string `json:"Icon"`
	Color    string `json:"Color"`
	Order    string `json:"Order"`
}

type HandbookItem struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Price    int    `json:"Price"`
}
