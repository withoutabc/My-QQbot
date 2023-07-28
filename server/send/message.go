package send

import (
	"my-QQbot/model"
	"my-QQbot/server"
)

func MsgSend(groupId int64, msg string) {
	menu := &model.Reply{
		GroupId:    groupId,
		AutoEscape: "false",
	}
	menu.Message = msg
	server.Send(menu)
}
