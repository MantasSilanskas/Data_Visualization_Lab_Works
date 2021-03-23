package server

import (
	"log"
	"net/http"
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

func HumidityHandler(w http.ResponseWriter, r *http.Request) {

}
