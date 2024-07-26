package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func JSON(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func Succeed(c *gin.Context, data any) {
	Failed(c, Code.Success, "", data)
}

func Success(c *gin.Context) {
	Succeed(c, nil)
}

func Failed(c *gin.Context, code *ErrorCode, msg string, data any) {
	if len(msg) == 0 {
		msg = code.Msg()
	}
	response := newResponse(code.Code(), msg, data)
	JSON(c, http.StatusOK, response)
}

func newResponse(code int, msg string, data any) *Result {
	response := &Result{
		Code: code,
		Msg:  msg,
	}

	if data == nil {
		response.Data = make([]interface{}, 0)
	} else {
		response.Data = data
	}

	return response
}
