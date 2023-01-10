package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI())) //create a new client

	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second) //Trying to connect to mongodb in 20 seconds
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil) //requesting
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB() //to access our db from anywhere

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golangAPI").Collection(collectionName)
} //create collection in mongo
