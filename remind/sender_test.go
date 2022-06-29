package remind

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
	//sender := &Sender{
	//	sendChan:  make(chan SenderMsg, 10),
	//	reminders: make([]ReminderInterface, 0),
	//}
	//now := time.Now()
	//startTime := now
	//endTime := now.Add(time.Second * 60)
	//way := model.ReminderWay{
	//	Name: common.DingTalkText,
	//	Desc: "描述cron",
	//}
	//reminder := model.Reminder{
	//	//ReminderWay:     way,
	//	CircleStartTime: startTime,
	//	CircleEndTime:   endTime,
	//}
	//per := NewReminderCron(reminder, "0 30 14 * * 2,3,4,5,6 ?")
	//sender.AddReminder(per)
	//time.Sleep(time.Second * 70)
}

func TestSign(t *testing.T) {
	// https://oapi.dingtalk.com/robot/send?access_token=d9770adfe5f2e9a4fba80a989c71bee821703d83e80776854054ed91d4876e71
	secret := "SEC8ecb6c95842ae90335557f4cb974a7fbc836617f4123b50b6be82c42afa8dff1"
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
