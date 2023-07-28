package receive

import (
	"github.com/spf13/cast"
	"my-QQbot/global"
	"my-QQbot/server/send"
	"my-QQbot/tools"
)

func GroupMessage(form map[string]interface{}) {
	userId := cast.ToInt64(form["user_id"])
	if userId != global.MasterQQId {
		return
	}
	groupId := cast.ToInt64(form["group_id"])
	//  messageId := cast.ToInt32(form["message_id"])
	msg := cast.ToString(form["message"])
	if tools.JudgeAtRobot(msg) {
		go send.MsgSend(groupId, msg)
	}
	//	messageD := md5.Sum([]byte(message))
	//	messageHash := hex.EncodeToString(messageD[:])
}
