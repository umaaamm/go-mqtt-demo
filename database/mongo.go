package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var SensorCollection *mongo.Collection

func InitMongoDB(uri, dbName, collection string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	SensorCollection = MongoClient.Database(dbName).Collection(collection)
	fmt.Println("Connected to MongoDB")
	return nil
}
