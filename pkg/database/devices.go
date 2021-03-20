package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//Devices data collection name
var devicesCollection = "devicesData"

// AddDeviceData adds device data to database
func AddDeviceData(client mongo.Client, input []BSONDeviceData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(database).Collection(devicesCollection)

	for _, v := range input {
		_, err := collection.InsertOne(ctx, v)
		if err != nil {
			log.Println("Failed to insert data to database. Error:", err)
			return err
		}
	}

	return nil
}
