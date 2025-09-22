package middleware

import (
	"sagala/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		requestID, _ := c.Get("X-Request-ID")

		logFields := logrus.Fields{
			"status_code": c.Writer.Status(),
			"latency_ms":  latency.Milliseconds(),
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        path,
			"request_id":  requestID,
			"user_agent":  c.Request.UserAgent(),
		}

		if len(c.Errors) > 0 {
			logFields["errors"] = c.Errors
			logger.Logger.WithFields(logFields).Error("request completed with errors")
		} else {
			logger.Logger.WithFields(logFields).Info("request completed successfully")
		}
	}
}
