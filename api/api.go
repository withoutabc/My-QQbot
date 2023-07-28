package api

import (
	"github.com/gin-gonic/gin"
	"my-QQbot/middleware"
)

func SeverSetUp() {
	r := gin.Default()
	r.POST("/info", middleware.LimitMiddleware(), HandleMsg)
	r.Run(":5701")
}
