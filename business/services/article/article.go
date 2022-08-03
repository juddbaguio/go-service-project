package article_service

import (
	"go-service-project/business/interface/data"

	"go.uber.org/zap"
)

type Wrapper struct {
	articleRepo data.ArticleRepository
	logger      *zap.Logger
}

func New(articleRepo data.ArticleRepository, logger *zap.Logger) *Wrapper {
	return &Wrapper{
		articleRepo: articleRepo,
		logger:      logger,
	}
}

func (w *Wrapper) GetArticleList() {
	w.logger.Info("Fetching Articles")
	w.articleRepo.GetArticleList()
}
