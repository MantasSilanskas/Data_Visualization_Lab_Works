package main

import (
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	//var filename = "filesNames.csv"
	filter := bson.M{}

	client, err := database.Connection()
	if err != nil {
		log.Println("failed to connect to database.", err)
		return
	}

	// err = pkg.InsertData(client, filename)
	// if err != nil {
	// 	log.Println("failed to insert data to database. ", err)
	// 	return
	// }

	files, err := database.FilterAll(client, filter)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range files {
		log.Println(v.FileName)
	}
	log.Println(len(files))
}
