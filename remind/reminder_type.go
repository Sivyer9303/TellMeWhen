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
