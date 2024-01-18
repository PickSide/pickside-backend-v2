package common

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitDB() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	db = client.Database("Cluster0")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}
