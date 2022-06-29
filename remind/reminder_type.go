package remind

import (
	"tellMeWhen/model"
	"tellMeWhen/query"
	"time"
)

type ReminderInterface interface {
	// 开始提醒
	start(sendChan chan<- SenderMsg)
	GetReminderType() string
}

func GetReminderInterfaceByModel(r *model.Reminder) ReminderInterface {
	if r == nil {
		return nil
	}
	way := query.GetRemindQuery().GetReminderWayById(r.ReminderWayId)
	switch way.WayName {
	case "perminute":
		duration, err := time.ParseDuration(way.Params)
		if err != nil {
			return nil
		}
		return NewReminderPer(*r, duration)
	case "cron":
		return NewReminderCron(*r, way.Params)
	}
	return nil
}
