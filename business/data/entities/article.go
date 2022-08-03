package entities

import "go-service-project/business/domain"

type Article struct {
	Title  string `bson:"title"`
	Author string `bson:"author"`
	Body   string `body:"body"`
}

func (a *Article) ToDomain() domain.Article {
	return domain.Article{
		Title:  a.Title,
		Author: a.Author,
		Body:   a.Body,
	}
}
