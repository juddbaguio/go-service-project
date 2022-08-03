package author

import (
	"go-service-project/infrastructure/postgresql"
)

type Repository struct {
	psql postgresql.DB
}

func NewRepository(psql postgresql.DB) *Repository {
	return &Repository{
		psql: psql,
	}
}
