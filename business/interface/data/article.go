package data

import "go-service-project/business/domain"

type ArticleRepository interface {
	GetArticleList() []domain.Article
}
