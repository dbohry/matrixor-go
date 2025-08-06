package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"matrixor/config"
	"time"
)

func Init(cfg config.DatabaseConfig) (*mongo.Client, error) {
	if cfg.URI == "" {
		return nil, fmt.Errorf("MongoDB URI is empty")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	fmt.Printf("%s Connected to MongoDB\n", time.Now().Format(time.RFC3339))
	return client, nil
}
