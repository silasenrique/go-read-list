package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	collection *mongo.Collection
}

func (c *Collection) Create(ctx context.Context, document interface{}) error {
	_, err := c.collection.InsertOne(ctx, &document)

	return err
}

func (c *Collection) GetById(ctx context.Context, id primitive.ObjectID) *mongo.SingleResult {
	return c.collection.FindOne(ctx, bson.M{"_id": id})
}
