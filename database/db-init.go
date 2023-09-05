package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DatabaseInit() {
	options := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		log.Fatalln(err)
	}

	DB = client.Database(os.Getenv("DATABASE"))
}
