package controller

import (
	"blog-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, model.Article{
		Id: id,
	})
}
