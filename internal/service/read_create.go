package service

import (
	"go-readlist/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ReadService) Create(read *entities.Read) (*entities.Read, error) {
	read.Id = primitive.NewObjectID()

	err := r.rep.Create(read)

	return read, err
}
