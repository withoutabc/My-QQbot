package server

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"my-QQbot/global"
)

func Send(inter interface{}, ty string) {
	client := req.C().EnableDumpAllWithoutResponse()
	_, err := client.R().SetBody(inter).
		Post(fmt.Sprintf("%s/%s", global.CqHttpUrl, ty))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("发送了一条消息")
}
