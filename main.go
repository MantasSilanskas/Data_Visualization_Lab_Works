package main

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/db"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/internal/results"
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

	_, err = results.AllResults(names, client)
	if err != nil {
		log.Println("Failed to calculate results. Error:", err)
		return
	}
}
