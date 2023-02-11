package utils

import (
	"TinyTikTok/conf/setup"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// SlowLog 慢执行日志
func SlowLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		latency := time.Since(begin)
		status := c.Writer.Status()
		if latency > 20000000 {
			setup.Logger("slow").Info().Msg(fmt.Sprintf("%s %s %s %d\n", c.Request.RequestURI, begin, latency, status))
		}
	}
}
