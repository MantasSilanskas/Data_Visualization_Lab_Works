package main

import (
	"log"
	"math/rand"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/results"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/server"
	"github.com/go-echarts/go-echarts/v2/opts"
)

const filename = "filesNames.csv"

func main() {

	client, err := db.Connection()
	if err != nil {
		log.Println("Failed to connect to database. Error:", err)
		return
	}

	err = internal.InsertData(client, filename)
	if err != nil {
		log.Println("Failed to insert data to database. Error:", err)
		return
	}

	names, err := db.UniqueDevicesIDs(client)
	if err != nil {
		log.Println("Failed to extract file names from database. Error:", err)
		return
	}

	res, err := results.CalcResults(names, client)
	if err != nil {
		log.Println("Failed to calculate results. Error:", err)
		return
	}

	humidity, temperature, co2 := results.PrepareResults(res)

	_, _, _ = results.GenerateHumidityBarData(humidity)
	_, _, _ = results.GenerateTemperatureBarData(temperature)
	_, _, _ = results.GenerateCo2BarData(co2)

	err = server.Connection()
	if err != nil {
		log.Println("Failed to start HTTP server. Error:", err)
		return
	}
}

func generateLineItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}
