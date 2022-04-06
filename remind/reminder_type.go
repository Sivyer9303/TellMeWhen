package remind

import (
	"fmt"
	"tellMeWhen/common"
	"tellMeWhen/model"
	"time"
)

type ReminderInterface interface {
	// 开始提醒
	start(sendChan chan<- SenderMsg)
	GetReminderType() string
}

type MsgFormatter interface {
	FormatMsg(reminder *model.Reminder) string
}

var _ ReminderInterface = (*ReminderPer)(nil)

// 固定间隔提醒器
type ReminderPer struct {
	reminder model.Reminder
	duration time.Duration
}

// 创建一个固定间隔提醒器
func NewReminderPer(reminder model.Reminder, duration time.Duration) *ReminderPer {
	return &ReminderPer{
		reminder: reminder,
		duration: duration,
	}
}

func (rp *ReminderPer) start(sendChan chan<- SenderMsg) {
	fmt.Println("start")
	endTime := rp.reminder.CircleEndTime
	startTime := rp.reminder.CircleStartTime
	now := time.Now()
	if now.After(endTime) {
		return
	} else if now.Before(startTime) {
		// 等待直到该启动
		startChan := time.After(now.Sub(startTime))
		<-startChan
	}
	endChan := time.After(endTime.Sub(now))
	tick := time.Tick(rp.duration)
	fmt.Println("start for select")
	for {
		select {
		case <-endChan:
			return
		case <-tick:
			// 到了该触发的时候了，组装数据
			fmt.Println("开始发送消息")
			msg := SenderMsg{
				id:  rp.reminder.ID,
				way: rp.reminder.ReminderWay,
			}
			sendChan <- msg
		}
	}
}

func (rp *ReminderPer) GetReminderType() string {
	return common.ReminderPer
}
