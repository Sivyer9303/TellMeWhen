package remind

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"tellMeWhen/common"
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
		Name: common.DingTalkText,
		Desc: "描述cron",
	}
	reminder := model.Reminder{
		ReminderWay:     way,
		CircleStartTime: startTime,
		CircleEndTime:   endTime,
	}
	per := NewReminderCron(reminder, "0 30 14 * * * ?")
	sender.AddReminder(per)
	time.Sleep(time.Second * 70)
}

func TestSign(t *testing.T) {
	secret := "SECf896a430fa46b7fb4c8539255fa218283167dca01179b4be8079d07ada900bf7"
	// 1656406786128
	// 1656406830
	// 1656406849850203500
	// 1656406890000
	ts := time.Now().UnixNano() / 1e6
	fmt.Println(ts)
	strToHash := fmt.Sprintf("%d\n%s", ts, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	toString := base64.StdEncoding.EncodeToString(data)
	fmt.Println(toString)
}
