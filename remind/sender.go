package remind

import "tellMeWhen/model"

type Sender struct {
	sendChan chan SenderMsg
}

type SenderMsg struct {
	id  uint
	way model.ReminderWay
}
