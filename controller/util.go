package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func paramInt64(ctx *gin.Context, key string) (int64, error) {
	s := ctx.Param(key)
	return strconv.ParseInt(s, 10, 64)
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
