package utils

import (
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/filter"
)

type CalculatedData struct {
	DeviceID         string  `json:"deviceId"`
	HumidityCount    int     `json:"humidityCount"`
	HumidityMean     float32 `json:"humidityMean"`
	TemperatureCount int     `json:"temperatureCount"`
	TemperatureMean  float32 `json:"temperatureMean"`
	Co2Count         int     `json:"co2Count"`
	Co2Mean          float32 `json:"co2Mean"`
}

func CalcDeviceData(data filter.FilteredData) CalculatedData {

	Hmean, Hcount := calcHumidyti(data)
	Tmean, Tcount := calcTemperature(data)
	Cmean, Ccount := calcCo2(data)

	list := CalculatedData{
		DeviceID:         data.DeviceID,
		HumidityCount:    Hcount,
		HumidityMean:     Hmean,
		TemperatureCount: Tcount,
		TemperatureMean:  Tmean,
		Co2Count:         Ccount,
		Co2Mean:          Cmean,
	}

	return list
}

func calcHumidyti(data filter.FilteredData) (mean float32, count int) {

	var total float32

	if len(data.Humidity) == 0 {
		return 0, 0
	}

	for _, v := range data.Humidity {
		total += v.Value
	}

	mean = total / float32(len(data.Humidity))
	count = len(data.Humidity)

	return mean, count
}

func calcTemperature(data filter.FilteredData) (mean float32, count int) {

	var total float32

	if len(data.Temperature) == 0 {
		return 0, 0
	}

	for _, v := range data.Temperature {
		total += v.Value
	}

	mean = total / float32(len(data.Temperature))
	count = len(data.Temperature)

	return mean, count
}

func calcCo2(data filter.FilteredData) (mean float32, count int) {

	var total float32

	if len(data.Co2) == 0 {
		return 0, 0
	}

	for _, v := range data.Co2 {
		total += v.Value
	}

	mean = total / float32(len(data.Co2))
	count = len(data.Co2)

	return mean, count
}
