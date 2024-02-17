package service

import (
	"go-readlist/internal/database"
	"go-readlist/internal/repository"
)

type ReadService struct {
	rep repository.IReadRepository
}

func NewReadCreateService(db database.ICollection) *ReadService {
	return &ReadService{
		rep: repository.NewReadRepository(db),
	}
}
