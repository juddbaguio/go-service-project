package data

type Transaction interface {
	Commit() error
	Rollback() error
}
