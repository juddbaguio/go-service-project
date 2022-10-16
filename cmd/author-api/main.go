package main

import (
	"go-service-project/business/data/repositories/author"
	"go-service-project/business/data/transactions"
	"go-service-project/business/protocol/http"
	author_http "go-service-project/business/protocol/http/author"
	"os"

	author_service "go-service-project/business/services/author"
	"go-service-project/foundation/logger"
	"go-service-project/infrastructure/mongo"
	"go-service-project/infrastructure/postgresql"
	"log"
)

func main() {
	zap := logger.New()
	defer zap.Sync()

	db, err := postgresql.NewDB("localhost:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	mongo, err := mongo.NewDB("localhost:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	articleRepo := author.NewRepository(db)
	transactionContaienr := transactions.NewContainer(db, mongo)
	authorSrv := author_service.New(articleRepo, transactionContaienr, zap)

	authorHttp := author_http.NewHTTP(zap, authorSrv)
	http.Start(authorHttp)
}
