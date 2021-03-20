package database

import (
	"context"
	"log"
	"time"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Database and collections names
var (
	databaseName = "Local"
	dataFiles    = "files"
	data         = "data"
)

// AddFileNames adds file name to files collection after they data has
// been insert to data collection
func AddFileNames(client mongo.Client, input reader.File) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(databaseName).Collection(dataFiles)

	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "id", Value: primitive.NewObjectID()},
		{Key: "fileId", Value: input.Number},
		{Key: "fileName", Value: input.Name},
	})
	if err != nil {
		log.Println("Failed to insert data to database", err)
		return err
	}
	return nil
}

func AddDeviceData(client mongo.Client, input reader.DeviceData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(databaseName).Collection(data)

	return nil
}
