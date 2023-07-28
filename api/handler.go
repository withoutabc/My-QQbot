package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"my-QQbot/handler/receive"
)

// HandleMsg 处理请求信息
func HandleMsg(ctx *gin.Context) {
	var form map[string]interface{}
	if ctx.ShouldBind(&form) != nil {
		return
	}
	postType := cast.ToString(form["post_type"])
	switch postType {
	case "message":
		receive.GroupMessage(form)
	case "notice":
		subType := cast.ToString(form["sub_type"])
		switch subType {
		case "poke":
			receive.Poke(form)
		}
	}
}
