package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

	fmt.Println("hello")
	nameItems := []string{"testas"}
	bar := charts.NewBar()
	bar.SetGlobalOptions()
	bar.SetXAxis(nameItems)
	bar.ExtendYAxis(opts.YAxis{Name: "Mean", Data: []float32{5.5}})
	bar.ExtendYAxis(opts.YAxis{Name: "Count", Data: []int{5}})
	_, err := os.Create("humidiry-bar.html")
	if err != nil {
		log.Println("Failed to create humidyti bar")
	}
	bar.Render(w)
}
