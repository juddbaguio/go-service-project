package author

import (
	"go-service-project/infrastructure/postgresql"

	"gorm.io/gorm"
)

type Repository struct {
	psql *gorm.DB
}

func NewRepository(psql postgresql.DB) *Repository {
	return &Repository{
		psql: psql,
	}
}
