package data

import "go-service-project/business/domain"

type AuthorRepository interface {
	GetAuthor() domain.Author
	BeginTx() AuthorTxRepository
}

type AuthorTxRepository interface {
	AuthorRepository
	Transaction
}
