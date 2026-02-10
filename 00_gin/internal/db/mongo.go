package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Connect(cfg config.config) (*mongo.Client, *mongo.Database, error) {
	// prevent your app from freezing man
	cxt, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

}
