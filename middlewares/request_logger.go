package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

// RequestLogger Request logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetString("x-request-id")
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL.Path

		t := time.Now()

		c.Next()

		latency := float32(time.Since(t).Seconds())

		status := c.Writer.Status()
		log.Info().Str("request_id", requestId).Str("client_ip", clientIp).
			Str("user_agent", userAgent).Str("method", method).Str("path", path).
			Float32("latency", latency).Int("status", status).Msg("")

	}
}
