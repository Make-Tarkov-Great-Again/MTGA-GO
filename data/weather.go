package data

import (
	"MT-GO/tools"
	"fmt"
	"math/rand"
	"time"
)

func GetWeather() *Weather {
	return db.weather
}

// #endregion

// #region Weather setters

func setWeather() {
	db.weather = &Weather{
		Acceleration: 7,
	}

	t := time.Now()
	db.weather.Date = t.Format("2006-01-02")
	db.weather.Time = t.Format("15:04:05")
	db.weather.WeatherInfo = generateWeather()
	db.weather.WeatherInfo.Date = db.weather.Date
	db.weather.WeatherInfo.Time = fmt.Sprintf("%s %s", db.weather.Date, db.weather.Time)
}

func generateWeather() WeatherReport {
	cloud := tools.RoundToThousandths(rand.Float32()*2 - 1)
	var rain int8
	var rainIntensity float32
	var fog float32

	if cloud > 0.5 {
		rain = int8(rand.Intn(5))
		rainIntensity = 1

		switch rain {
		case 1:
			fog = 0.008
		case 2:
			fog = 0.012
		case 3:
			fog = 0.02
		case 4:
			fog = 0.03
		default:
			fog = 0.004
		}
	} else {
		rainIntensity = tools.RoundToThousandths(rand.Float32())
		fog = tools.RoundToThousandths(rand.Float32()*0.003 + 0.003)
	}

	return WeatherReport{
		Timestamp:     tools.GetCurrentTimeInSeconds(),
		Cloud:         cloud,
		WindDirection: int8(rand.Intn(7) + 1),
		WindSpeed:     tools.RoundToThousandths(rand.Float32() * 4),
		Rain:          rain,
		RainIntensity: rainIntensity,
		Temperature:   22,
		Pressure:      int16(tools.GetRandomInt(750, 780)),
		Fog:           fog,
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
	WindDirection int8    `json:"wind_direction"`
	WindGustiness float32 `json:"wind_gustiness"`
	Rain          int8    `json:"rain"`
	RainIntensity float32 `json:"rain_intensity"`
	Fog           float32 `json:"fog"`
	Temperature   int8    `json:"temp"`
	Pressure      int16   `json:"pressure"`
	Date          string  `json:"date"`
	Time          string  `json:"time"`
}

// #endregion
