package middleware

import (
	"github.com/dingdinglz/test-blog/config"
	"github.com/gin-gonic/gin"
)

// CORS CORS中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取配置的允许来源
		allowOrigins := config.AppConfig.CORS.AllowOrigins
		origin := "*"
		if len(allowOrigins) > 0 {
			origin = allowOrigins[0]
		}

		// 设置CORS响应头
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// 处理OPTIONS预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
