package results

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/filter"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
func GenerateHumidityBarData(data []HumidityResults) ([]string, []float32, []int) {

	IDs := []string{}
	Means := []float32{}
	Counts := []int{}

	for _, v := range data {
		IDs = append(IDs, v.DevicesID)
		Means = append(Means, v.Mean)
		Counts = append(Counts, v.Count)
	}

	return IDs, Means, Counts
}

// Think of better way to return data then 3 arrays maybe struct?
func GenerateTemperatureBarData(data []TemperatureResults) ([]string, []float32, []int) {

	IDs := []string{}
	Means := []float32{}
	Counts := []int{}

	for _, v := range data {
		IDs = append(IDs, v.DevicesID)
		Means = append(Means, v.Mean)
		Counts = append(Counts, v.Count)
	}

	return IDs, Means, Counts
}

// Think of better way to return data then 3 arrays maybe struct?
func GenerateCo2BarData(data []Co2Results) ([]string, []float32, []int) {

	IDs := []string{}
	Means := []float32{}
	Counts := []int{}

	for _, v := range data {
		IDs = append(IDs, v.DevicesID)
		Means = append(Means, v.Mean)
		Counts = append(Counts, v.Count)
	}

	return IDs, Means, Counts
}
