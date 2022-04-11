package remind

import (
	"tellMeWhen/model"
	"testing"
	"time"
)

func TestSender_Send_Reminder_Per(t *testing.T) {
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

func TestSender_Send_Reminder_Cron(t *testing.T) {
	sender := &Sender{
		sendChan:  make(chan SenderMsg, 10),
		reminders: make([]ReminderInterface, 0),
	}
	now := time.Now()
	startTime := now
	endTime := now.Add(time.Second * 60)
	way := model.ReminderWay{
		Name: "测试cron",
		Desc: "描述cron",
	}
	reminder := model.Reminder{
		ReminderWay:     way,
		CircleStartTime: startTime,
		CircleEndTime:   endTime,
	}
	per := NewReminderCron(reminder, "* * 10 * * * *")
	sender.AddReminder(per)
	time.Sleep(time.Second * 70)
}
