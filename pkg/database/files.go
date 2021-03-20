package database

import (
	"context"
	"log"
	"time"

	"github.com/MantasSilanskas/Data_Visualization_Lab_Works/pkg/reader"
	"go.mongodb.org/mongo-driver/mongo"
)

//Files collection name
var filesCollection = "files"

// AddFileNames adds file name to files collection after they data has
// been insert to data collection
func AddFile(client mongo.Client, input BSONFile) error {

	input.InputTime = time.Now().Local()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(database).Collection(filesCollection)

	_, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Println("Failed to insert data to database. Error:", err)
		return err
	}
	return nil
}

func ASDAS(client mongo.Client, input reader.File) error {

	return nil
}
