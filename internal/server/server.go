package server

import (
	"log"
	"net/http"
)

func Connection() error {

	// Handlers
	http.HandleFunc("/humidity/mean", HumidityMeanHandler)
	http.HandleFunc("/humidity/count", HumidityCountHandler)
	http.HandleFunc("/temperature/mean", TemperatureMeanHandler)
	http.HandleFunc("/temperature/count", TemperatureCountHandler)
	http.HandleFunc("/co2/mean", Co2MeanHandler)
	http.HandleFunc("/co2/count", Co2CountHandler)

	// HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Failed to start HTTP server. Error:", err)
		return err
	}

	return nil
}
