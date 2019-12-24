package controller

import (
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
	user := engine.Group("/users")
	user.GET("/:name", GetUser)

	article := engine.Group("/articles")
	article.GET("/:id", GetArticle)
}
