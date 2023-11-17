package data

import (
	"MT-GO/tools"
	"fmt"
	"math/rand"
	"time"
)

var weather = Weather{
	Acceleration: 7,
}

func GetWeather() *Weather {
	return &weather
}

// #endregion

// #region Weather setters

func setWeather() {
	t := time.Now()
	weather.Date = t.Format("2006-01-02")
	weather.Time = t.Format("15:04:05")
	weather.WeatherInfo = generateWeather()
	weather.WeatherInfo.Date = weather.Date
	weather.WeatherInfo.Time = fmt.Sprintf("%s %s", weather.Date, weather.Time)
}

func generateWeather() WeatherReport {
	num := tools.RoundToThousandths(rand.Float32()*2 - 1)
	num2 := rand.Intn(7) + 1
	num3 := tools.RoundToThousandths(rand.Float32() * 4)
	num4 := tools.RoundToThousandths(rand.Float32())
	var num5 int
	var num6 float32

	if num > 0.5 {
		num5 = rand.Intn(5)
		num4 = 1
		num6 = 0.004
		switch num5 {
		case 1:
			num6 = 0.008
		case 2:
			num6 = 0.012
		case 3:
			num6 = 0.02
		case 4:
			num6 = 0.03
		}
	} else {
		num5 = 0
		num6 = tools.RoundToThousandths(rand.Float32()*0.003 + 0.003)
	}

	return WeatherReport{
		Timestamp:     tools.GetCurrentTimeInSeconds(),
		Cloud:         num,
		WindDirection: num2,
		WindSpeed:     num3,
		Rain:          num5,
		RainIntensity: num4,
		Temperature:   22,
		Pressure:      780,
		Fog:           num6,
	}
}

// #endregion

// #region Weather structs

type Weather struct {
	WeatherInfo  WeatherReport `json:"weather"`
	Date         string        `json:"date"`
	Time         string        `json:"time"`
	Acceleration int           `json:"acceleration"`
}

type WeatherReport struct {
	Timestamp     int64   `json:"timestamp"`
	Cloud         float32 `json:"cloud"`
	WindSpeed     float32 `json:"wind_speed"`
	WindDirection int     `json:"wind_direction"`
	WindGustiness float32 `json:"wind_gustiness"`
	Rain          int     `json:"rain"`
	RainIntensity float32 `json:"rain_intensity"`
	Fog           float32 `json:"fog"`
	Temperature   int     `json:"temp"`
	Pressure      int     `json:"pressure"`
	Date          string  `json:"date"`
	Time          string  `json:"time"`
}

// #endregion
