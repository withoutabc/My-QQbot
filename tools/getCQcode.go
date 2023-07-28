package tools

import "fmt"

func CodePoke(userId int64) string {
	return fmt.Sprintf("[CQ:%s,qq=%v]", "poke", userId)
}
