package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var _mdb *mongo.Client

func MongoDB() *mongo.Client {
	return _mdb
}

func MongoConn(uri string) (bool, error) {
	mongoOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		log.Println("[MongoDB] Connect Error:", err)
		return false, err
	}
	_mdb = client
	return true, nil
}
