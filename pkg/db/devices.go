package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// FilterAllDevices returns all documents in devicesData collection
func FilterAllDevices(client *mongo.Client) ([]BSONDeviceData, error) {

	list := []BSONDeviceData{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(database).Collection(devicesCollection)

	c, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to extract data from database. Error:", err)
		return list, err
	}

	for c.Next(ctx) {
		file := BSONDeviceData{}

		err := c.Decode(&file)
		if err != nil {
			log.Println("Failed to decode cursor. Error:", err)
			return list, err
		}
		list = append(list, file)
	}

	return list, nil
}

func FilterDevicesByID(client *mongo.Client, filter bson.M) ([]BSONDeviceData, error) {
	list := []BSONDeviceData{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(database).Collection(devicesCollection)

	c, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Failed to extract data from database. Error:", err)
		return list, err
	}

	for c.Next(ctx) {
		file := BSONDeviceData{}

		err := c.Decode(&file)
		if err != nil {
			log.Println("Failed to decode cursor. Error:", err)
			return list, err
		}
		list = append(list, file)
	}

	return list, nil
}
