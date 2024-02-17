package repository

import (
	"go-readlist/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReadRepository interface {
	Create(read *entities.Read) error
	FindById(id primitive.ObjectID) (*entities.Read, error)
}
