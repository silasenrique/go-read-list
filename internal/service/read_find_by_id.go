package service

import (
	"errors"
	"go-readlist/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ReadService) FindById(id string) (*entities.Read, error) {
	if id == "" {
		return nil, errors.New("precisa ser informado alguma coisa")
	}

	parseId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	data, err := r.rep.FindById(parseId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
