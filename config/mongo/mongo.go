package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var db *mongo.Client

func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("Error creating mongodb: %s", err)
	}

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}

	fmt.Println("connected to nosql database:")

	db = client
}

func Get() *mongo.Client {
	return db
}
