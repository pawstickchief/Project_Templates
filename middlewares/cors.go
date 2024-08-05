package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求,支持options访问
func Cors(ClientUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", ClientUrl) //支持全域名访问，不安全，部署后需要固定限制为客户端网址
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
