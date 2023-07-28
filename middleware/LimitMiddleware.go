package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

var (
	RequestCount int
	latestSet    time.Time
)

func LimitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		RequestCount++
		if time.Since(latestSet) > time.Minute {
			RequestCount = 0
			latestSet = time.Now()
		}
		if RequestCount >= 10 {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
