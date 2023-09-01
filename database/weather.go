package database

import (
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var weather = Weather{}

func GetWeather() *Weather {
	return &weather
}
func setWeather() {
	raw := tools.GetJSONRawMessage(weatherPath)
	err := json.Unmarshal(raw, &weather)
	if err != nil {
		panic(err)
	}
}

// #region Weather structs
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

// #endregion
