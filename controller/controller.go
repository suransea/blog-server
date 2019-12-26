package controller

import (
	"github.com/gin-gonic/gin"
)

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Init(engine *gin.Engine) {
	tag := engine.Group("/tags")
	tag.GET("", GetTags)
	tag.GET("/:id", GetTag)
	tag.POST("", PostTag)
	tag.PUT("/:id", PutTag)
	tag.DELETE("/:id", DeleteTag)
}
