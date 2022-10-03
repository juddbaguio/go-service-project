package author

import (
	"go-service-project/infrastructure/postgresql"

	"gorm.io/gorm"
)

type Repository struct {
	psql *gorm.DB
}

type authTxContainer struct {
	txConn *gorm.DB
	*Repository
}

func NewRepository(psql postgresql.DB) *Repository {
	return &Repository{
		psql: psql,
	}
}
