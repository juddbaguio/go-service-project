package main

import (
	"go-service-project/business/data/repositories/article"
	"go-service-project/business/protocol/http"
	article_http "go-service-project/business/protocol/http/article"
	article_service "go-service-project/business/services/article"
	"go-service-project/foundation/logger"
	"os"

	"go-service-project/infrastructure/mongo"
	"log"
)

func main() {
	zap := logger.New()
	defer zap.Sync()

	db, err := mongo.NewDB("localhost:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	articleRepo := article.NewRepository(db)
	articleSrv := article_service.New(articleRepo, zap)
	articleHttp := article_http.NewHTTP(zap, articleSrv)

	http.Start(articleHttp)
}
