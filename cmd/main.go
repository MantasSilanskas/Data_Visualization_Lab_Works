package main

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg"
	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/db"
)

func main() {

	var filename = "filesNames.csv"

	client, err := db.Connection()
	if err != nil {
		log.Println("failed to connect to database.", err)
		return
	}

	err = pkg.InsertData(client, filename)
	if err != nil {
		log.Println("failed to insert data to database. ", err)
		return
	}

}
