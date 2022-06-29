package remind

import (
	"fmt"
	"tellMeWhen/common"
	"tellMeWhen/model"
	"tellMeWhen/query"
	"time"

	"github.com/gorhill/cronexpr"
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
	fmt.Println("start cron")
	for {
		now := time.Now()
		reminder := rc.reminder
		endTime := reminder.CircleEndTime
		startTime := reminder.CircleStartTime
		fmt.Println("start to deal cron reminder ", startTime.Format("2006-01-02 15:04:05"), "        ", endTime.Format("2006-01-02 15:04:05"))
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
		fmt.Println("next send time : ", next.Format("2006-01-02 15:04:05"))
		duration := next.Sub(now)
		after := time.After(duration)
		<-after
		way := query.GetRemindQuery().GetReminderWayById(reminder.ReminderWayId)
		msg := SenderMsg{way: *way}
		sendChan <- msg
	}
}

func (rc *ReminderCron) GetReminderType() string {
	return common.ReminderCron
}
