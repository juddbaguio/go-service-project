package author_http

import (
	author_handler "go-service-project/business/protocol/http/author/handlers"

	"github.com/go-chi/chi/v5"
)

func (c *Container) InitRoutes(r *chi.Mux) {
	handler := author_handler.InitHandlers(c.logger, c.authorService)
	r.Get("/", handler.GetAuthorList)
}
