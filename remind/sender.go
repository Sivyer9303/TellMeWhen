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
	id  uint
	way model.ReminderWay
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

// 钉钉 https://oapi.dingtalk.com/robot/send?access_token=ec6c5321ae775c71d48c93a33796e3683f62b5282cddc01d96cd8b9e6e02ea21
func sendMsg(msg SenderMsg) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("oh shit....sendMsg panic")
		}
	}()
	way := msg.way
	name := way.Name
	switch name {
	case common.DingTalkText:

	}
	fmt.Println("start sendMsg,msg:", msg)
}
