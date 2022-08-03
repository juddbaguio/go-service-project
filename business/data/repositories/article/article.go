package article

import (
	"go-service-project/infrastructure/mongo"
)

type Repository struct {
	mongo mongo.DB
}

func NewRepository(mongo mongo.DB) *Repository {
	return &Repository{
		mongo: mongo,
	}
}
