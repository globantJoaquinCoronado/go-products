package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGODB_HOST = os.Getenv("MONGODB_HOST")
var uri = "mongodb://" + MONGODB_HOST + ":27017"
var db = "academy-golang"

func GetCollection(collection string) *mongo.Collection {
	fmt.Printf("mongo URI; %s \n", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	cancelFunction()

	if err != nil {
		panic(err.Error)
	}

	return client.Database(db).Collection(collection)
}
