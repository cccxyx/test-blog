package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 计算处理时间
		latency := time.Since(startTime)

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// 输出日志
		log.Printf("[%s] %s | %s | %d | %v | %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			"INFO",
			method,
			statusCode,
			latency,
			clientIP+" "+path,
		)
	}
}
