package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		os.Getenv("MONGODB_USERNAME"),
		os.Getenv("MONGODB_PASSWORD"),
		os.Getenv("MONGODB_URI"),
	)
	opt := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	conn, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return conn.Database(os.Getenv("MONGODB_DATABASE")), nil
}
