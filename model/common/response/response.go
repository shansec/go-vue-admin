package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	RESET   = 101
	ERROR   = 201
	SUCCESS = 200
	NOAUTH  = 4033
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailNoAuthDetailed(data interface{}, message string, c *gin.Context) {
	Result(NOAUTH, data, message, c)
}

func Reset(c *gin.Context) {
	Result(RESET, map[string]interface{}{}, "操作失败", c)
}

func ResetWithMessage(message string, c *gin.Context) {
	Result(RESET, map[string]interface{}{}, message, c)
}

func ResetWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(RESET, data, message, c)
}
