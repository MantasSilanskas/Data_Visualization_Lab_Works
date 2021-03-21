package filter

import "time"

type FilteredData struct {
	DeviceID    string            `json:"deviceId"`
	Humidity    []HumidityData    `json:"humidityData"`
	Temperature []TemperatureData `json:"temperatureData"`
	Co2         []Co2Data         `json:"co2Data"`
}

type HumidityData struct {
	Type      string    `json:"type"`
	Value     float32   `json:"value"`
	TimeStamp time.Time `json:"timeStamp"`
}

type TemperatureData struct {
	Type      string    `json:"type"`
	Value     float32   `json:"value"`
	TimeStamp time.Time `json:"timeStamp"`
}

type Co2Data struct {
	Type      string    `json:"type"`
	Value     float32   `json:"value"`
	TimeStamp time.Time `json:"timeStamp"`
}
