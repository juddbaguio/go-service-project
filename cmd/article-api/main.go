package main

import (
	"go-service-project/business/data/repositories/article"
	article_service "go-service-project/business/services/article"
	"go-service-project/foundation/logger"

	"go-service-project/infrastructure/mongo"
	"log"
)

func main() {
	zap := logger.New()
	defer zap.Sync()

	db, err := mongo.NewDB("localhost:8080")
	if err != nil {
		log.Println(err)
	}

	articleRepo := article.NewRepository(db)

	articleSrv := article_service.New(articleRepo, zap)

	articleSrv.GetArticleList()
}
