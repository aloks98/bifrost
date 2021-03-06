package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.New()
		hostname, err := os.Hostname()
		if err != nil {
			logger.Error("Cannot determine Hostname: ", err)
			hostname = "Unknown"
		}
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})
		msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(time.RFC1123), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
		if statusCode >= http.StatusInternalServerError {
			entry.Error(msg)
		} else if statusCode >= http.StatusBadRequest {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}
	}
}
