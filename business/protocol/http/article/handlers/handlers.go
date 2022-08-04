package article_handler

import (
	"encoding/json"
	article_service "go-service-project/business/services/article"
	"net/http"

	"go.uber.org/zap"
)

type Container struct {
	logger         *zap.Logger
	articleService *article_service.Container
}

func InitHandlers(logger *zap.Logger, articleService *article_service.Container) *Container {
	return &Container{
		logger:         logger,
		articleService: articleService,
	}
}

func (c *Container) GetArticleList(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Visiting Path: ")
	c.logger.Info(r.RequestURI)

	articleList := c.articleService.GetArticleList()

	json, err := json.Marshal(articleList)
	if err != nil {
		c.logger.Warn(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
