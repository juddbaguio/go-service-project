package transactions

import (
	"context"
	"go-service-project/business/data/repositories/article"
	"go-service-project/business/data/repositories/author"
	"go-service-project/business/interface/data"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type TransactionContainer struct {
	dbConn      *gorm.DB
	mongoClient *mongo.Client
}

func NewContainer(db *gorm.DB, mongo *mongo.Client) *TransactionContainer {
	return &TransactionContainer{
		dbConn:      db,
		mongoClient: mongo,
	}
}

func (tc *TransactionContainer) Do(fn func(articleRepo data.ArticleRepository, authorRepo data.AuthorRepository) error) error {
	tx := tc.dbConn.Begin()
	mongoSession, _ := tc.mongoClient.StartSession()

	defer func() {
		tx.Rollback()
		mongoSession.AbortTransaction(context.Background())
	}()

	if err := fn(article.NewRepository(mongoSession.Client()), author.NewRepository(tx)); err != nil {
		return err
	}

	tx.Commit()
	mongoSession.CommitTransaction(context.Background())
	return nil
}
