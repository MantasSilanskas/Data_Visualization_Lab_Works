package utils

import "github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/filter"

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

	for _, v := range data.Humidity {
		total += v.Value
	}

	mean = total / float32(len(data.Humidity))

	return mean, len(data.Humidity)
}

func calcTemperature(data filter.FilteredData) (mean float32, count int) {

	var total float32

	for _, v := range data.Temperature {
		total += v.Value
	}

	mean = total / float32(len(data.Temperature))

	return mean, len(data.Temperature)
}

func calcCo2(data filter.FilteredData) (mean float32, count int) {

	var total float32

	for _, v := range data.Co2 {
		total += v.Value
	}

	mean = total / float32(len(data.Co2))

	return mean, len(data.Co2)
}
