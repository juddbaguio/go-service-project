package author_service

import (
	"go-service-project/business/domain"
	"go-service-project/business/interface/data"

	"go.uber.org/zap"
)

type Container struct {
	authorRepo data.AuthorRepository
	logger     *zap.Logger
}

func New(authorRepo data.AuthorRepository, logger *zap.Logger) *Container {
	return &Container{
		authorRepo: authorRepo,
		logger:     logger,
	}
}

func (w *Container) GetAuthor() domain.Author {
	w.logger.Info("Getting Author")
	return w.authorRepo.GetAuthor()
}
