package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoStr = os.Getenv("MONGOSTRCONNECT")

func DBMongo(dbname string) *mongo.Database {
	if MongoStr == "" {
		panic(fmt.Errorf("MONGOSTRCONNECT ENV NOT FOUND"))
	}
	clay, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoStr))
	if err != nil {
		panic(fmt.Errorf("MongoConnect: %+v \n", err))
	}
	return clay.Database(dbname)
}
