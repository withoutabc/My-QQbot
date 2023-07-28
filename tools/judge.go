package tools

import (
	"fmt"
	"my-QQbot/global"
	"strings"
)

func JudgeAtRobot(msg string) bool {
	if strings.Contains(msg, fmt.Sprintf("[CQ:at,qq=%v]", global.RobotQQId)) ||
		strings.Contains(msg, "@绝") {
		return true
	}
	return false
}
