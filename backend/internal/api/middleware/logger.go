package middleware

import (
	"master-guide-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.Info("HTTP Request",
			logger.String("method", param.Method),
			logger.String("path", param.Path),
			logger.Int("status", param.StatusCode),
			logger.String("latency", param.Latency.String()),
			logger.String("client_ip", param.ClientIP),
			logger.String("user_agent", param.Request.UserAgent()),
		)
		return ""
	})
}

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Error("Panic recovered",
			logger.String("path", c.Request.URL.Path),
			logger.String("method", c.Request.Method),
			logger.Any("error", recovered),
		)
		c.JSON(500, gin.H{
			"code":    500,
			"message": "Internal server error",
		})
	})
}
