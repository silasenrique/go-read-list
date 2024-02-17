package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	*mongo.Client
}

func NewConnection(connectionString string) *DB {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Send a ping to confirm a successful connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return &DB{client}
}

func (d *DB) GetCollection(database, collection string) ICollection {
	return &Collection{collection: d.Database(database).Collection(collection)}
}
