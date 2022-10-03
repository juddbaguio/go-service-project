package author

import (
	"go-service-project/business/domain"
	"go-service-project/business/interface/data"
)

func (a *Repository) GetAuthor() domain.Author {
	return domain.Author{}
}

func (a *Repository) BeginTx() data.AuthorTxRepository {
	tx := a.psql.Begin()

	container := &authTxContainer{
		txConn:     tx,
		Repository: NewRepository(tx),
	}

	return container
}

func (atx *authTxContainer) Commit() error {
	return atx.txConn.Commit().Error
}

func (atx *authTxContainer) Rollback() error {
	return atx.txConn.Rollback().Error
}
