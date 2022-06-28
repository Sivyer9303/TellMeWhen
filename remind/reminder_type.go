package remind

import (
	"tellMeWhen/model"
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
	switch r.ReminderType.Name {
	case "perminute":
		duration, err := time.ParseDuration(r.ReminderType.Param)
		if err != nil {
			return nil
		}
		return NewReminderPer(*r, duration)
	}
	return nil
}
