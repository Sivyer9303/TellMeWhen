package query

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"tellMeWhen/model"
)

type ReminderQuery struct {
	state int
	db    *gorm.DB
}

const (
	STATE_INITIALING = iota
	STATE_READY
	STATE_ERROR
)

var once sync.Once
var RQ *ReminderQuery

func initReminderQuery() {
	once.Do(func() {
		rq := &ReminderQuery{
			state: STATE_INITIALING,
		}
		var conStr = "root:root@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
		if err != nil {
			log.Fatal("create database fail")
			rq.state = STATE_ERROR
		} else {
			rq.state = STATE_READY
			db.AutoMigrate(&model.Reminder{})
			rq.db = db
		}
		RQ = rq
	})
}

func GetRemindQuery() *ReminderQuery {
	initReminderQuery()
	return RQ
}

// 查询所有提醒
func (rq *ReminderQuery) GetAllReminder() []*model.Reminder {
	var res []*model.Reminder
	err := rq.db.Find(&res).Error
	if err != nil {
		log.Fatal("get all reminder fail")
		return nil
	}
	for _, v := range res {
		// 关联查询
		t := new(model.ReminderType)
		w := new(model.ReminderWay)
		rq.db.Model(v).Association("ReminderType").Find(t)
		rq.db.Model(v).Association("ReminderWay").Find(w)
		v.ReminderType = *t
		v.ReminderWay = *w
	}
	return res
}
