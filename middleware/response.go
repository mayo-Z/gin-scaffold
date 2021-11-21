package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const SuccessCode ResponseCode = iota

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
}

func ResponseError(ctx *gin.Context, code ResponseCode, err error) {

	res := &Response{
		ErrorCode: code,
		ErrorMsg:  err.Error(),
		Data:      "",
	}
	ctx.JSON(200, res)
	response, _ := json.Marshal(res)
	ctx.Set("response", string(response))
	ctx.AbortWithError(200, err)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {

	res := &Response{
		ErrorCode: SuccessCode,
		ErrorMsg:  "",
		Data:      data,
	}
	ctx.JSON(200, res)
	response, _ := json.Marshal(res)
	ctx.Set("response", string(response))
}
