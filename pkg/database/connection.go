package database

import (
	"context"
	"log"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	AddFileNames(input reader.File) error
	AddDeviceData(input []reader.DeviceData) error
	connection() (*mongo.Client, error)
}

// Connection connect to local mongo database
func Connection() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("failed to connect to database. ", err)
		return client, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to ping database. ", err)
		return client, err
	}

	log.Println("Connected to MongoDB successfully!")

	return client, nil
}
