package database

import (
	"context"
	"log"
	"time"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
)

//Database and collections names
var (
	databaseName = "local"
	dataFiles    = "files"
	data         = "devicesData"
)

// AddFileNames adds file name to files collection after they data has
// been insert to data collection
func AddFile(client mongo.Client, input reader.File) error {

	input.InputTime = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(databaseName).Collection(dataFiles)

	_, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Println("Failed to insert data to database. Error:", err)
		return err
	}
	return nil
}

// AddDeviceData adds device data to database
func AddDeviceData(client mongo.Client, input []reader.DeviceData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(databaseName).Collection(data)

	for _, v := range input {
		_, err := collection.InsertOne(ctx, v)
		if err != nil {
			log.Println("Failed to insert data to database. Error:", err)
			return err
		}
	}

	return nil
}
