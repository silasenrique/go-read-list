package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	Create(ctx context.Context, document interface{}) error
	GetById(ctx context.Context, id primitive.ObjectID) *mongo.SingleResult
}
