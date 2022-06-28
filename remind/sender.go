package remind

import (
	"fmt"
	"tellMeWhen/common"
	"tellMeWhen/model"
)

type Sender struct {
	sendChan  chan SenderMsg
	reminders []ReminderInterface
}

type SenderMsg struct {
	id     uint
	way    model.ReminderWay
	params map[string]interface{}
}

func (s *Sender) AddReminder(re ReminderInterface) {
	s.reminders = append(s.reminders, re)
	go re.start(s.sendChan)
	go s.Send()
}

func (s *Sender) Send() {
	for {
		select {
		case msg := <-s.sendChan:
			go sendMsg(msg)
		}
	}
}

// 钉钉 https://oapi.dingtalk.com/robot/send?access_token=a3e76af9d2e3e79423822fa2c742b89405db57481123582105407000e63af7e6
func sendMsg(msg SenderMsg) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("oh shit....sendMsg panic...")
		}
	}()
	way := msg.way
	name := way.Name
	switch name {
	case common.DingTalkText:
		fmt.Println("start dingtalk message")

		break
	default:
		break
	}
	fmt.Println("start sendMsg,msg:", msg)
}
