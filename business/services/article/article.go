package article_service

import (
	"go-service-project/business/domain"
	"go-service-project/business/interface/data"

	"go.uber.org/zap"
)

type Container struct {
	articleRepo data.ArticleRepository
	logger      *zap.Logger
}

func New(articleRepo data.ArticleRepository, logger *zap.Logger) *Container {
	return &Container{
		articleRepo: articleRepo,
		logger:      logger,
	}
}

func (w *Container) GetArticleList() []domain.Article {
	w.logger.Info("Fetching Articles")
	articleList := w.articleRepo.GetArticleList()

	return articleList
}
