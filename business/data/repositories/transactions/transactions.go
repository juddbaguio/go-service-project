package transactions

import "go-service-project/business/interface/data"

type TransactionContainer struct {
	AuthorRepo data.AuthorRepository
}
