package middleware

import (
	"github.com/gin-gonic/gin"
	"my-QQbot/global"
	"time"
)

var (
	RequestCount int
	latestSet    time.Time
)

// LimitMiddleware 限流中间件，一分钟至多n条
func LimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		RequestCount++
		if time.Since(latestSet) > time.Minute {
			RequestCount = 0
			latestSet = time.Now()
		}
		if RequestCount >= global.MaxHttpRequest {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
