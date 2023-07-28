package send

import (
	"my-QQbot/global"
	"my-QQbot/model"
	"my-QQbot/server"
	"my-QQbot/tools"
)

func Poke(groupId, userId int64) {
	reply := &model.Reply{
		GroupId:    groupId,
		AutoEscape: "false",
	}
	reply.Message = tools.CodePoke(userId)
	server.Send(reply, global.GroupMsg)
}
