package remind

import (
	"tellMeWhen/model"
	"testing"
	"time"
)

func TestSender_Send(t *testing.T) {
	sender := &Sender{
		sendChan:  make(chan SenderMsg, 10),
		reminders: make([]ReminderInterface, 0),
	}
	now := time.Now()
	startTime := now
	endTime := now.Add(time.Second * 60)
	way := model.ReminderWay{
		Name: "测试",
		Desc: "描述",
	}
	reminder := model.Reminder{
		ReminderWay:     way,
		CircleStartTime: startTime,
		CircleEndTime:   endTime,
	}
	per := NewReminderPer(reminder, time.Second)
	sender.AddReminder(per)
	time.Sleep(time.Second * 70)
}
