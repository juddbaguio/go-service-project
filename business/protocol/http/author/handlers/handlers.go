package author_handler

import (
	"encoding/json"
	author_service "go-service-project/business/services/author"
	"net/http"

	"go.uber.org/zap"
)

type Container struct {
	logger        *zap.Logger
	authorService *author_service.Container
}

func InitHandlers(logger *zap.Logger, authorService *author_service.Container) *Container {
	return &Container{
		logger:        logger,
		authorService: authorService,
	}
}

func (c *Container) GetAuthorList(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Visiting Path: ")
	c.logger.Info(r.RequestURI)

	author := c.authorService.GetAuthor()

	json, err := json.Marshal(author)
	if err != nil {
		c.logger.Warn(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
