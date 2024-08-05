package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1000, 		// 程序中的错误码
	"msg": xx,				// 提示信息
	"data": {},				// 数据

}


*/

// 定义返回参数的数据结构
type ResponseDate struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResopnseError(c *gin.Context, code ResCode) {
	rd := &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResopnseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
func ResopnseSystemDataSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

// 自定义响应
func ResponseErrorwithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: msg,
	}
	c.JSON(http.StatusOK, &rd)
}
