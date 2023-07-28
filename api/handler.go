package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"my-QQbot/server/send"
)

func HandleMsg(ctx *gin.Context) {
	var form map[string]interface{}
	if ctx.ShouldBind(&form) != nil {
		return
	}
	postType := cast.ToString(form["post_type"])
	switch postType {
	case "message":
		userId := cast.ToInt64(form["user_id"])
		log.Println(userId)
		messageId := cast.ToInt32(form["message_id"])
		log.Println(messageId)
		message := cast.ToString(form["message"])
		log.Println(message)
		//	messageD := md5.Sum([]byte(message))
		//	messageHash := hex.EncodeToString(messageD[:])
		groupId := cast.ToInt64(form["group_id"])
		send.MsgSend(groupId, message)
	}
}
