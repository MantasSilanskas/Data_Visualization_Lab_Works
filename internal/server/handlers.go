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

func HumidityMeanHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Humidity records data mean bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices humidity records data mean"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Humidity mean (%)", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)

	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.HumidityMean)

	f, err := os.Create("charts/humidityMean.html")
	if err != nil {
		log.Println("Failed to create humidyti means bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("Humidity mean chart has been rendered at localhost:8080/humidity/mean")
}

func HumidityCountHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Humidity record data count bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices humidity record count data"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Humidity records count", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.HumidityCount)

	f, err := os.Create("charts/humidityCounts.html")
	if err != nil {
		log.Println("Failed to create humidyti data counts bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("Humidity count chart has been rendered at localhost:8080/humidity/count")
}

func Co2MeanHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "CO2 records data mean bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices CO2 records data mean"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "CO2 mean (ppm)", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.Co2Mean)

	f, err := os.Create("charts/co2Mean.html")
	if err != nil {
		log.Println("Failed to create co2 mean bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("CO2 mean chart has been rendered at localhost:8080/co2/mean")
}

func Co2CountHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "CO2 records data count bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices CO2 records data count"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "CO2 records count", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.Co2Count)

	f, err := os.Create("charts/co2Counts.html")
	if err != nil {
		log.Println("Failed to create CO2 data counts bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("CO2 mean chart has been rendered at localhost:8080/co2/count")
}

func TemperatureMeanHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Temperature records data mean bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices temperature records data mean"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Tempereture mean (C)", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Mean", results.TemperatureMean)

	f, err := os.Create("charts/temperatureMean.html")
	if err != nil {
		log.Println("Failed to create temperature means bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("Temperature mean chart has been rendered at localhost:8080/temperature/mean")
}

func TemperatureCountHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Temperature records data mean bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices temperature records count data"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Device ID", AxisLabel: &opts.AxisLabel{Color: "red"}}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Temperature records count", AxisLabel: &opts.AxisLabel{Color: "red"}}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Count", results.TemperatureCount)

	f, err := os.Create("charts/temperatureCounts.html")
	if err != nil {
		log.Println("Failed to create temperature data counts bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("Temperature count chart has been rendered at localhost:8080/temperature/count")
}

func DevicesDataMeansHandlerBad(w http.ResponseWriter, _ *http.Request) {
	nameItems := results.DevicesID

	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, PageTitle: "Devices records data mean bar chart"}),
		charts.WithTitleOpts(opts.Title{Title: "Devices records data means results"}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Temperature mean (C)", results.TemperatureMean).
		AddSeries("Co2 mean (ppm)", results.Co2Mean).
		AddSeries("Humidity mean (%)", results.HumidityMean)

	f, err := os.Create("charts/devicesDataMeanBad.html")
	if err != nil {
		log.Println("Failed to create temperature data counts bar")
	}

	bar.Render(f)
	bar.Render(w)

	log.Println("Devices records data mean chart has been rendered at localhost:8080/devices/data/mean/bad")
}
