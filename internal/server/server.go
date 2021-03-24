package server

import (
	"log"
	"net/http"
)

func Connection() error {

	// Handlers
	http.HandleFunc("/humidityMean", HumidityMeanHandler)
	http.HandleFunc("/humidityCount", HumidityCountHandler)
	http.HandleFunc("/temperatureMean", TemperatureMeanHandler)
	http.HandleFunc("/temperatureCount", TemperatureCountHandler)
	http.HandleFunc("/co2Mean", Co2MeanHandler)
	http.HandleFunc("/co2Count", Co2CountHandler)

	// HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Failed to start HTTP server. Error:", err)
		return err
	}
	return nil
}
