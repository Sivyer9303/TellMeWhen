package remind

type ReminderInterface interface {
	// 开始提醒
	start(sendChan chan<- SenderMsg)
	GetReminderType() string
}
