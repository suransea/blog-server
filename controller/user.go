package controller

import (
	"blog-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, model.User{
		Username: name,
	})
}
