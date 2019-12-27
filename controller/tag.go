package controller

import (
	"blog-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(ctx *gin.Context) {
	tags, err := model.GetTags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(tags))
}

func GetTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	tag, err := model.GetTag(id)
	if err == model.ErrNoResults {
		ctx.JSON(http.StatusNotFound, rspError(ErrNotFound, err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(tag))
}

func GetTagArticles(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	articles, err := model.GetTagArticles(id)
	if err == model.ErrNoResults {
		ctx.JSON(http.StatusNotFound, rspError(ErrNotFound, err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(articles))
}

func PutTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	var tag model.Tag
	err = ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	err = model.UpdateTag(id, tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(id))
}

func PostTag(ctx *gin.Context) {
	var tag model.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	tag.Id, err = model.AddTag(tag.Tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusCreated, rspSuccess(tag.Id))
}

func DeleteTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspError(ErrBadRequest, err))
		return
	}
	err = model.RemoveTag(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspError(ErrInternal, err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(id))
}
