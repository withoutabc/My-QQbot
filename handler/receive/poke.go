package receive

import (
	"github.com/spf13/cast"
	"my-QQbot/global"
	"my-QQbot/server/send"
)

func Poke(form map[string]interface{}) {
	groupId := cast.ToInt64(form["group_id"])
	userId := cast.ToInt64(form["user_id"])
	targetId := cast.ToInt64(form["target_id"])
	if targetId == global.RobotQQId {
		go send.Poke(groupId, userId)
	}
}
