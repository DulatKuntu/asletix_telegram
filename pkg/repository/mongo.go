package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DBURI  string
	DBName string
}

const (
	collectionUser   = "users"
	collectionRefers = "refers"
	collectionFeeds  = "myWorkoutsHistory"
)

func NewMongoDB(cfg Config) (*mongo.Database, *mongo.Client, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBURI))
	if err != nil {
		return nil, nil, err
	}

	database := client.Database(cfg.DBName)

	return database, client, nil
}
