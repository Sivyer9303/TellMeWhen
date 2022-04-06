package remind

import (
	"fmt"
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

func sendMsg(msg SenderMsg) {
	fmt.Println("start sendMsg,msg:", msg)
}
