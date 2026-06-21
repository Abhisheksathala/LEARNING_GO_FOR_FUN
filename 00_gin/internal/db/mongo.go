package db

// import (
// 	"context"
// 	"fmt"
// 	"notes-api/internal/config"
// 	"time"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
// 	// prevent your app from freezing man
// 	cxt, cancle := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancle()

// 	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

// 	client, err := mongo.Connect(cxt, clientOptions)

// 	if err != nil {
// 		return nil, nil, fmt.Errorf("mongoconnect Failed")
// 	}
// 	if err := client.Ping(cxt, nil); err != nil {
// 		return nil, nil, fmt.Errorf("mongo ping failed")
// 	}

// 	database := client.Database(cfg.MongoDB)

// 	return client, database, nil

// }

// func Disconnect(client *mongo.Client) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	return client.Disconnect(ctx)
// }

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect =exported function
func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	// prevent app from freezing in startup

	ctx, cancle := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

	defer cancle()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI) // return the client object

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, nil, fmt.Errorf("Mongo connect failed")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Mongo ping failed")
	}

	database := client.Database(cfg.MongoDB)

	return client, database, nil

}

func Disconnect(client *mongo.Client) error {
	// stop over for app from running even after work is done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
