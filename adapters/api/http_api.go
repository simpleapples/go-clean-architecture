package api

import (
	"clean_architecutre_demo/usecases"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type httpAPIAdapter struct {
	app            *gin.Engine
	articleUsecase usecases.IArticleUsecase
}

func newHTTPAPIAdapter() IAPIAdapter {
	articleUsecase := usecases.NewArticleUsecase()
	app := gin.Default()
	return &httpAPIAdapter{app: app, articleUsecase: articleUsecase}
}

func (a *httpAPIAdapter) RunService(port int) error {
	addr := fmt.Sprintf(":%d", port)
	return a.app.Run(addr)
}

func (a *httpAPIAdapter) InitHandlers() {
	a.app.GET("/article/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		article := a.articleUsecase.Show(id)
		c.JSON(200, gin.H{
			"title":   article.Title,
			"content": article.Content,
		})
	})
}
