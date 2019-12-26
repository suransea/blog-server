package controller

import (
	"blog-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(ctx *gin.Context) {
	tags, err := model.SelectTags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspErrInternal(err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(tags))
}

func GetTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspErrBadRequest(err))
		return
	}
	tag, err := model.SelectTagById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspErrInternal(err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(tag))
}

func PutTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspErrBadRequest(err))
		return
	}
	var tag model.Tag
	err = ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspErrBadRequest(err))
		return
	}
	err = model.UpdateTagById(id, tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspErrInternal(err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(id))
}

func PostTag(ctx *gin.Context) {
	var tag model.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspErrBadRequest(err))
		return
	}
	tag.Id, err = model.InsertTag(tag.Tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspErrInternal(err))
		return
	}
	ctx.JSON(http.StatusCreated, rspSuccess(tag.Id))
}

func DeleteTag(ctx *gin.Context) {
	id, err := paramInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rspErrBadRequest(err))
		return
	}
	err = model.DeleteTagById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, rspErrInternal(err))
		return
	}
	ctx.JSON(http.StatusOK, rspSuccess(id))
}
