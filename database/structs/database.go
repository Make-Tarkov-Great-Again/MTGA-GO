package structs

type CoreStruct struct {
	PlayerTemplate PlayerTemplate
	ClientSettings ClientSettings
	ServerConfig   ServerConfig
	Globals        Globals
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

type Weather struct {
	WeatherInfo struct {
		Timestamp     int64   `json:"timestamp"`
		Cloud         float64 `json:"cloud"`
		WindSpeed     int     `json:"wind_speed"`
		WindDirection int     `json:"wind_direction"`
		WindGustiness float64 `json:"wind_gustiness"`
		Rain          int     `json:"rain"`
		RainIntensity int     `json:"rain_intensity"`
		Fog           float64 `json:"fog"`
		Temperature   int     `json:"temp"`
		Pressure      int     `json:"pressure"`
		Date          string  `json:"date"`
		Time          string  `json:"time"`
	} `json:"weather"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Acceleration int    `json:"acceleration"`
}

type Customization struct {
	ID     string `json:"_id,omitempty"`
	Name   string `json:"_name,omitempty"`
	Parent string `json:"_parent,omitempty"`
	Type   string `json:"_type,omitempty"`
	Proto  string `json:"_proto,omitempty"`
	Props  struct {
		Name                string      `json:"Name,omitempty"`
		ShortName           string      `json:"ShortName,omitempty"`
		Description         string      `json:"Description,omitempty"`
		Side                []string    `json:"Side,omitempty"`
		BodyPart            string      `json:"BodyPart,omitempty"`
		Prefab              interface{} `json:"Prefab,omitempty"`
		WatchPrefab         interface{} `json:"WatchPrefab,omitempty"`
		IntegratedArmorVest bool        `json:"IntegratedArmorVest,omitempty"`
		WatchPosition       XYZ         `json:"WatchPosition,omitempty"`
		WatchRotation       XYZ         `json:"WatchRotation,omitempty"`
		AvailableAsDefault  bool        `json:"AvailableAsDefault,omitempty"`
	} `json:"_props,omitempty"`
}
