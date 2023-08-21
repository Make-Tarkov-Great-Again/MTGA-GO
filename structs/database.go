package structs

type Database struct {
	Core *Core
	//Connections *ConnectionStruct
	Items     map[string]interface{}
	Locales   *Locale
	Languages map[string]interface{}
	Handbook  *Handbook
	Traders   map[string]*Trader
	Flea      *Flea
	Quests    map[string]interface{}
	Hideout   *Hideout

	Locations     *Locations
	Weather       *Weather
	Customization map[string]interface{}
	Editions      map[string]interface{}
	Bot           *Bots
	Profiles      map[string]*Profile
	//bundles  []map[string]interface{}
}

type Edition struct {
	Bear    *PlayerTemplate `json:"bear"`
	Usec    *PlayerTemplate `json:"usec"`
	Storage *EditionStorage `json:"storage"`
}

type EditionStorage struct {
	Bear []string `json:"bear"`
	Usec []string `json:"usec"`
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
	Locale map[string]interface{}
	Menu   *LocaleMenu
}

type LocaleMenu struct {
	Menu map[string]string `json:"menu"`
}

type Handbook struct {
	Categories []HandbookCategory `json:"Categories"`
	Items      []HandbookItem     `json:"Items"`
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
		Timestamp     int     `json:"timestamp"`
		Cloud         float32 `json:"cloud"`
		WindSpeed     int     `json:"wind_speed"`
		WindDirection int     `json:"wind_direction"`
		WindGustiness float32 `json:"wind_gustiness"`
		Rain          int     `json:"rain"`
		RainIntensity int     `json:"rain_intensity"`
		Fog           float32 `json:"fog"`
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
	ID     string                 `json:"_id"`
	Name   string                 `json:"_name"`
	Parent string                 `json:"_parent"`
	Type   string                 `json:"_type"`
	Proto  string                 `json:"_proto"`
	Props  map[string]interface{} `json:"_props"`
}
