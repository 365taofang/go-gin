package response

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestResp(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		Failed(c, Code.InvalidParams, "", nil)
	})
	err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		return
	}
}
