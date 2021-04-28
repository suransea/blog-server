package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Common response struct.
type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Error codes in a response.
const (
	ErrNotFound   = 1
	ErrInternal   = 2
	ErrBadRequest = 3
)

// Initialize routes for the engine.
func Init(engine *gin.Engine) {
	article := engine.Group("/articles")
	article.GET("", GetArticles)
	article.GET("/:id", GetArticle)
	article.GET("/:id/content", GetArticleContent)
	article.GET("/:id/tags", GetArticleTags)

	tag := engine.Group("/tags")
	tag.GET("", GetTags)
	tag.GET("/:id", GetTag)
	tag.GET("/:id/articles", GetTagArticles)
	tag.POST("", PostTag)
	tag.PUT("/:id", PutTag)
	tag.DELETE("/:id", DeleteTag)
}

func paramInt64(ctx *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(ctx.Param(key), 10, 64)
}

func rspSuccess(data interface{}) Rsp {
	return Rsp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func rspError(code int, err error) Rsp {
	return Rsp{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	}
}
