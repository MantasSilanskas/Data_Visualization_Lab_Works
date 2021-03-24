package server

import (
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func Connection() error {

	// Handlers
	http.HandleFunc("/", HumidityHandler)

	// HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Failed to start HTTP server. Error:", err)
		return err
	}
	return nil
}

func HumidityHandler(w http.ResponseWriter, _ *http.Request) {

	nameItems := []string{"1", "2", "3", "4", "5", "6", "7"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Devices humidity means and total counts of humidity data",
		}))
	bar.SetXAxis(nameItems).
		AddSeries("Mean", generateLineItems()).
		AddSeries("Count", generateLineItems())
	_, err := os.Create("humidiry-bar.html")
	if err != nil {
		log.Println("Failed to create humidyti bar")
	}
	bar.Render(w)
}

// generate random data for line chart
func generateLineItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}
