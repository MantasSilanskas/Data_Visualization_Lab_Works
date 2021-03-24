package server

import (
	"log"
	"net/http"
	"os"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/results"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func HumidityMeanHandler(w http.ResponseWriter, r *http.Request) {
	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.Co2Mean)
	_, err := os.Create("humidityMean.html")
	if err != nil {
		log.Println("Failed to create humidyti means bar")
	}
	bar.Render(w)
	log.Println("Humidity mean chart has been rendered at localhost:8080/humidityMean")
}

func HumidityCountHandler(w http.ResponseWriter, r *http.Request) {

	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.Co2Count)
	_, err := os.Create("humidityCounts.html")
	if err != nil {
		log.Println("Failed to create humidyti data counts bar")
	}
	bar.Render(w)
	log.Println("Humidity count chart has been rendered at localhost:8080/humidityCount")
}

func Co2MeanHandler(w http.ResponseWriter, r *http.Request) {

	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.Co2Mean)
	_, err := os.Create("co2Mean.html")
	if err != nil {
		log.Println("Failed to create co2 mean bar")
	}
	bar.Render(w)
	log.Println("CO2 mean chart has been rendered at localhost:8080/co2Mean")
}

func Co2CountHandler(w http.ResponseWriter, r *http.Request) {

	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.Co2Count)
	_, err := os.Create("co2Counts.html")
	if err != nil {
		log.Println("Failed to create CO2 data counts bar")
	}
	bar.Render(w)
	log.Println("CO2 mean chart has been rendered at localhost:8080/co2Count")
}

func TemperatureMeanHandler(w http.ResponseWriter, r *http.Request) {

	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.TemperatureMean)
	_, err := os.Create("temperatureMean.html")
	if err != nil {
		log.Println("Failed to create temperature means bar")
	}
	bar.Render(w)
	log.Println("Temperature mean chart has been rendered at localhost:8080/temperatureMean")
}

func TemperatureCountHandler(w http.ResponseWriter, r *http.Request) {

	nameItems := results.DevicesID
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity data counts",
		}),
		charts.WithLegendOpts(opts.Legend{Show: true}))
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.TemperatureCount)
	_, err := os.Create("temperatureCounts.html")
	if err != nil {
		log.Println("Failed to create temperature data counts bar")
	}
	bar.Render(w)
	log.Println("Temperature count chart has been rendered at localhost:8080/temperatureCount")
}
