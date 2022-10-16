package data

type Transaction interface {
	Do(func(articleRepo ArticleRepository, authorRepo AuthorRepository) error) error
}
