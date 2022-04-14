package remind

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"tellMeWhen/common"
	"tellMeWhen/model"
	"time"
)

var _ ReminderInterface = (*ReminderCron)(nil)

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
