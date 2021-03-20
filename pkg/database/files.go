package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Files collection name
var filesCollection = "files"

// AddFile adds file to files collection after they data has
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

// FilterAll returns all documents in files collection
func FilterAll(client *mongo.Client, filter bson.M) ([]BSONFile, error) {

	list := []BSONFile{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database(database).Collection(filesCollection)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Failed to extract data from database. Error:", err)
		return list, err
	}

	for cursor.Next(ctx) {
		file := BSONFile{}
		err := cursor.Decode(&file)
		if err != nil {
			log.Println("Failed to decode cursor. Error:", err)
			return list, err
		}
		list = append(list, file)
	}

	return list, nil
}
