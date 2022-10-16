package author_service

import (
	"errors"
	"go-service-project/business/domain"
	"go-service-project/business/interface/data"

	"go.uber.org/zap"
)

type Container struct {
	authorRepo data.AuthorRepository
	tx         data.Transaction
	logger     *zap.Logger
}

func New(authorRepo data.AuthorRepository, tx data.Transaction, logger *zap.Logger) *Container {
	return &Container{
		authorRepo: authorRepo,
		logger:     logger,
	}
}

func (w *Container) GetAuthor() domain.Author {
	w.logger.Info("Getting Author")
	return w.authorRepo.GetAuthor()
}

func (w *Container) TxGetAuthor() error {
	err := w.tx.Do(func(articleRepo data.ArticleRepository, authorRepo data.AuthorRepository) error {
		articleList := articleRepo.GetArticleList()
		if len(articleList) == 0 {
			return errors.New("no articles")
		}

		author := authorRepo.GetAuthor()
		if author.ID == 0 {
			return errors.New("no author")
		}

		return nil
	})

	return err
}
