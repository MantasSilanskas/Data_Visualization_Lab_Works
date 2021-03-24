package results

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/filter"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/utils"
	"github.com/go-echarts/go-echarts/v2/opts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DevicesID        []string
	HumidityMean     []opts.BarData
	HumidityCount    []opts.BarData
	TemperatureMean  []opts.BarData
	TemperatureCount []opts.BarData
	Co2Mean          []opts.BarData
	Co2Count         []opts.BarData
)

type HumidityResults struct {
	DevicesID string
	Mean      float32
	Count     int
}

type TemperatureResults struct {
	DevicesID string
	Mean      float32
	Count     int
}

type Co2Results struct {
	DevicesID string
	Mean      float32
	Count     int
}

func CalcResults(input []string, client *mongo.Client) (results []utils.CalculatedData, err error) {
	list := []utils.CalculatedData{}

	for _, v := range input {
		data, err := db.FilterDevicesByID(client, bson.M{"deviceId": v})
		if err != nil {
			log.Println("Failed to extract", v, "data from database. Error:", err)
			return list, err
		}
		filteredData, err := filter.FilterDeviceData(data, v)
		if err != nil {
			log.Println("Failed to filter out device data by type. Error:", err)
			return list, err
		}

		results := utils.CalcDeviceData(filteredData)

		list = append(list, results)
	}

	return list, nil
}

func PrepareResults(results []utils.CalculatedData) (humidityResults []HumidityResults, temperatureResults []TemperatureResults, co2Results []Co2Results) {

	humidityResults = prepareHumidityResults(results)
	temperatureResults = prepareTemperatureResults(results)
	co2Results = prepareCo2Results(results)

	return humidityResults, temperatureResults, co2Results
}

func prepareHumidityResults(data []utils.CalculatedData) []HumidityResults {

	list := []HumidityResults{}

	for _, v := range data {
		list = append(list, HumidityResults{
			DevicesID: v.DeviceID,
			Mean:      v.HumidityMean,
			Count:     v.HumidityCount,
		})
	}

	return list
}

func prepareTemperatureResults(data []utils.CalculatedData) []TemperatureResults {

	list := []TemperatureResults{}

	for _, v := range data {
		list = append(list, TemperatureResults{
			DevicesID: v.DeviceID,
			Mean:      v.TemperatureMean,
			Count:     v.TemperatureCount,
		})
	}

	return list
}

func prepareCo2Results(data []utils.CalculatedData) []Co2Results {

	list := []Co2Results{}

	for _, v := range data {
		list = append(list, Co2Results{
			DevicesID: v.DeviceID,
			Mean:      v.Co2Mean,
			Count:     v.Co2Count,
		})
	}
	return list
}

// Think of better way to return data then 3 arrays maybe struct?
func GenerateHumidityBarData(data []HumidityResults) ([]string, []opts.BarData, []opts.BarData) {

	DevicesID = []string{}
	HumidityMean = make([]opts.BarData, 0)
	HumidityCount = make([]opts.BarData, 0)

	for _, v := range data {
		DevicesID = append(DevicesID, v.DevicesID)
		HumidityMean = append(HumidityMean, opts.BarData{Value: int(v.Mean)})
		HumidityCount = append(HumidityCount, opts.BarData{Value: v.Count})
	}

	return DevicesID, HumidityMean, HumidityCount
}

// Think of better way to return data then 3 arrays maybe struct?
func GenerateTemperatureBarData(data []TemperatureResults) ([]string, []opts.BarData, []opts.BarData) {

	DevicesID = []string{}
	TemperatureMean = make([]opts.BarData, 0)
	TemperatureCount = make([]opts.BarData, 0)

	for _, v := range data {
		DevicesID = append(DevicesID, v.DevicesID)
		TemperatureMean = append(TemperatureMean, opts.BarData{Value: v.Mean})
		TemperatureCount = append(TemperatureCount, opts.BarData{Value: v.Count})
	}

	return DevicesID, TemperatureMean, TemperatureCount
}

// Think of better way to return data then 3 arrays maybe struct?
func GenerateCo2BarData(data []Co2Results) ([]string, []opts.BarData, []opts.BarData) {

	DevicesID = []string{}
	Co2Mean = make([]opts.BarData, 0)
	Co2Count = make([]opts.BarData, 0)

	for _, v := range data {
		DevicesID = append(DevicesID, v.DevicesID)
		Co2Mean = append(Co2Mean, opts.BarData{Value: v.Mean})
		Co2Count = append(Co2Mean, opts.BarData{Value: v.Count})
	}

	return DevicesID, Co2Mean, Co2Count
}
