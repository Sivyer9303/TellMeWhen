package remind

import (
	"tellMeWhen/model"
	"time"
)

type ReminderType interface {
	// 开始提醒
	start()
}

type MsgFormatter interface {
	FormatMsg(reminder *model.Reminder) string
}

// 固定间隔提醒器
type ReminderPer struct {
	sendChan chan<- SenderMsg
	reminder model.Reminder
	duration time.Duration
}

// 创建一个固定间隔提醒器
func NewReminderPer(reminder model.Reminder, ch chan<- SenderMsg) *ReminderPer {
	return &ReminderPer{
		sendChan: ch,
		reminder: reminder,
		duration: time.Second,
	}
}

func (rp *ReminderPer) start() {
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
	for {
		select {
		case <-endChan:
			return
		case <-tick:
			// 到了该触发的时候了，组装数据
			msg := SenderMsg{
				id:  rp.reminder.ID,
				way: rp.reminder.ReminderWay,
			}
			rp.sendChan <- msg
		}
	}
}
