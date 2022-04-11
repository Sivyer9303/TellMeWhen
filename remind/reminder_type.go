package remind

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"tellMeWhen/common"
	"tellMeWhen/model"
	"time"
)

type ReminderInterface interface {
	// 开始提醒
	start(sendChan chan<- SenderMsg)
	GetReminderType() string
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

// cron表达式提醒器
type ReminderCron struct {
	reminder   model.Reminder
	cronExpr   string
	expression *cronexpr.Expression
}

func NewReminderCron(reminder model.Reminder, cron string) *ReminderCron {
	expression, err := cronexpr.Parse(cron)
	if err != nil {
		return nil
	}
	return &ReminderCron{
		reminder:   reminder,
		cronExpr:   cron,
		expression: expression,
	}
}

func (rc *ReminderCron) start(sendChan chan<- SenderMsg) {
	for {
		fmt.Println("start cron")
		now := time.Now()
		endTime := rc.reminder.CircleEndTime
		startTime := rc.reminder.CircleStartTime
		fmt.Println(startTime.Format("2006-01-02 15:03:04"), "        ", endTime.Format("2006-01-02 15:03:04"))
		if startTime.After(now) {
			a := time.After(startTime.Sub(now))
			<-a
			continue
		}
		if endTime.Before(now) {
			return
		}
		next := rc.expression.Next(now)
		for next.Before(now) {
			next = rc.expression.Next(next)
		}
		duration := next.Sub(now)
		after := time.After(duration)
		<-after
		msg := SenderMsg{way: rc.reminder.ReminderWay}
		sendChan <- msg
	}
}

func (rc *ReminderCron) GetReminderType() string {
	return common.ReminderCron
}
