package remind

import (
	"fmt"
	"tellMeWhen/common"
	"tellMeWhen/model"
)

type Sender struct {
	sendChan  chan SenderMsg
	reminders []ReminderInterface
}

func GetSender() *Sender {
	return &Sender{
		sendChan:  make(chan SenderMsg, 10),
		reminders: make([]ReminderInterface, 0),
	}
}

type SenderMsg struct {
	id     uint
	way    model.ReminderWay
	params map[string]interface{}
}

func (s *Sender) AddReminder(re ReminderInterface) {
	s.reminders = append(s.reminders, re)
	go re.start(s.sendChan)
	go s.Send()
}

func (s *Sender) Send() {
	for {
		select {
		case msg := <-s.sendChan:
			go sendMsg(msg)
		}
	}
}

// 钉钉 https://oapi.dingtalk.com/robot/send?access_token=a3e76af9d2e3e79423822fa2c742b89405db57481123582105407000e63af7e6
func sendMsg(msg SenderMsg) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("oh shit....sendMsg panic...")
		}
	}()
	way := msg.way
	name := way.Name
	switch name {
	case common.DingTalkText:
		fmt.Println("start dingtalk message")
		break
	case common.Bible:
		fmt.Println("start send bible message")
		formatBibleMessage(way)
	default:
		break
	}
	fmt.Println("start sendMsg,msg:", msg)
}

// https://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=004685&datefrom=2022-06-20&dateto=2022-06-28&page=1
func formatBibleMessage(way model.ReminderWay) string {
	//t := way.LastReminderTime
	//ctx := context.Background()
	//callMethod := "GET"
	//baseUrl := "https://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=004685"
	//endPoint := "https://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=004685&datefrom=2022-06-20&dateto=2022-06-28&page=1"
	//header := make(map[string]string)
	//header["Cookie"] = "STOCK7-FINANCE-SINA-COM-CN="
	//body := []byte{}
	//response, err := common.doRequest(ctx, callMethod, endPoint, header, body)
	return ""

}
