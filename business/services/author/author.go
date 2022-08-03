package author_service

import (
	"go-service-project/business/interface/data"

	"go.uber.org/zap"
)

type Wrapper struct {
	authorRepo data.AuthorRepository
	logger     *zap.Logger
}

func New(authorRepo data.AuthorRepository, logger *zap.Logger) *Wrapper {
	return &Wrapper{
		authorRepo: authorRepo,
		logger:     logger,
	}
}

func (w *Wrapper) GetAuthor() {
	w.logger.Info("Getting Author")
	w.authorRepo.GetAuthor()
}
