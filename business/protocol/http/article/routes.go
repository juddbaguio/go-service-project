package article_http

import (
	article_handler "go-service-project/business/protocol/http/article/handlers"

	"github.com/go-chi/chi/v5"
)

func (c *Container) InitRoutes(r *chi.Mux) {
	handler := article_handler.InitHandlers(c.logger, c.articleService)
	r.Get("/", handler.GetArticleList)
}
