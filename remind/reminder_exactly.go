package remind

import (
	"fmt"
	"tellMeWhen/common"
	"tellMeWhen/model"
	"tellMeWhen/query"
	"time"
)

type ReminderExactly struct {
	time     time.Time
	reminder model.Reminder
}

func NewReminderExactly(time time.Time, reminder model.Reminder) *ReminderExactly {
	return &ReminderExactly{
		time:     time,
		reminder: reminder,
	}
}

func (re *ReminderExactly) start(sendChan chan<- SenderMsg) {
	t := re.time
	now := time.Now()
	after := t.After(now)
	if after {
		timeChan := time.After(t.Sub(now))
		<-timeChan
		way := query.GetRemindQuery().GetReminderWayById(re.reminder.ReminderWayId)
		msg := SenderMsg{id: 1111, way: *way}
		sendChan <- msg
	} else {
		fmt.Println("the time is before now,can not start")
	}
}

func (re *ReminderExactly) GetReminderType() string {
	return common.ReminderExactly
}
