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

func rspErrInternal(err error) Rsp {
	return Rsp{
		Code: 1,
		Msg:  err.Error(),
		Data: nil,
	}
}

func rspErrBadRequest(err error) Rsp {
	return Rsp{
		Code: 2,
		Msg:  err.Error(),
		Data: nil,
	}
}
