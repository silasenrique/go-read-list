package repository

import (
	"context"
	"go-readlist/internal/database"
	"go-readlist/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadRepository struct {
	db database.ICollection
}

func NewReadRepository(db database.ICollection) IReadRepository {
	return &ReadRepository{db}
}

func (r *ReadRepository) Create(read *entities.Read) error {
	return r.db.Create(context.Background(), read)
}

func (r *ReadRepository) FindById(id primitive.ObjectID) (*entities.Read, error) {
	result := r.db.GetById(context.Background(), id)
	if result.Err() != nil {
		return nil, result.Err()
	}

	content := &entities.Read{}

	err := result.Decode(content)
	if err != nil {
		return nil, err
	}

	return content, err
}
