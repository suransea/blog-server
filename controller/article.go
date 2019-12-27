package controller

import (
	"blog-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticles(ctx *gin.Context) {
	articles, err := model.GetArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(articles))
}

func GetArticle(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	article, err := model.GetArticle(id)
	if err == model.ErrNoResults {
		ctx.JSON(http.StatusNotFound, rspError(ErrNotFound, err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(article))
}

func GetArticleContent(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	content, err := model.GetArticleContent(id)
	if err == model.ErrNoResults {
		ctx.JSON(http.StatusNotFound, rspError(ErrNotFound, err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(content))
}
