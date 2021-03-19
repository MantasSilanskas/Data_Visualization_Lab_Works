package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("failed to connect to database", err)
	}
	defer client.Disconnect(ctx)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	reader.ReadFileData()

}
