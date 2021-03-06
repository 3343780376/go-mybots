package test

import (
	"fmt"
	"github.com/3343780376/go-mybots"
	"log"
)

var Bot = go_mybots.Bots{Address: "127.0.0.1", Port: 5700, Admin: 1743224847}

func init() {
	go_mybots.Info, _ = Bot.GetLoginInfo()
	go_mybots.ViewOnCoCommand = append(go_mybots.ViewOnCoCommand, go_mybots.ViewOnC0CommandApi{
		CoCommand: DefaultOnCoCommand, Command: "weather", Allies: "天气", RuleChecked: []go_mybots.Rule{}})
	go_mybots.ViewMessage = append(go_mybots.ViewMessage, go_mybots.ViewMessageApi{OnMessage: DefaultMessageHandle,
		MessageType: go_mybots.MessageTypeApi.Private, SubType: ""})
	go_mybots.ViewMessage = append(go_mybots.ViewMessage, go_mybots.ViewMessageApi{OnMessage: MessageTest,
		MessageType: go_mybots.MessageTypeApi.Group, SubType: ""})
	go_mybots.ViewNotice = append(go_mybots.ViewNotice, go_mybots.ViewOnNoticeApi{OnNotice: DefaultNoticeHandle,
		NoticeType: go_mybots.NoticeTypeApi.GroupIncrease,
		SubType:    "approve"})
	go_mybots.ViewRequest = append(go_mybots.ViewRequest, DefaultRequestHandle)
	go_mybots.ViewMeta = append(go_mybots.ViewMeta, DefaultMetaHandle)
}

func DefaultMessageHandle(event go_mybots.Event) {
	log.Println("收到了私聊信息")
	message := go_mybots.MessageAt(event.UserId).Message
	go Bot.SendPrivateMsg(event.UserId, "hello   world"+message, false)

}

func MessageTest(event go_mybots.Event) {
	log.Println("收到了")
	if event.GroupId == 972264701 {
		go Bot.SendGroupMsg(event.GroupId, event.Message, false)
	}
}

func DefaultNoticeHandle(event go_mybots.Event) {

}

func DefaultRequestHandle(event go_mybots.Event) {

}

func DefaultMetaHandle(event go_mybots.Event) {

}

func DefaultOnCoCommand(event go_mybots.Event, args []string) {
	log.Println("触发了天气命令")
	_, _ = Bot.SendMsg(event.MessageType, event.GroupId, "请输入你要查询的城市", false)
	nextEvent := Bot.GetNextEvent(10, event.UserId)
	fmt.Println(nextEvent.Message)
}
