package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient(dbName, uri string) (*mongo.Database, error) {
	opt := options.Client().ApplyURI(uri)
	conn, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return nil, err
	}
	return conn.Database(dbName), nil
}
