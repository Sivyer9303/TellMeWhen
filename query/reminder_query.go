package query

import (
	"fmt"
	"log"
	"sync"
	"tellMeWhen/model"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	fmt.Println("开始初始化reminder")
	once.Do(func() {
		rq := &ReminderQuery{
			state: STATE_INITIALING,
		}
		var conStr = "root:root@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
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
	fmt.Println("查询所有提醒结果为:", res)
	for i := range res {
		v := res[i]
		// 关联查询
		v.ReminderType = *rq.getReminderTypeById(v.ReminderTypeId)
		v.ReminderWay = *rq.getReminderWayById(v.ReminderWayId)
		fmt.Println(v)
	}
	return res
}

func (rq *ReminderQuery) getReminderTypeById(i int) *model.ReminderType {
	t := &model.ReminderType{}
	rq.db.Model(t).Find(t, "id = ?", i)
	return t
}

func (rq *ReminderQuery) getReminderWayById(i int) *model.ReminderWay {
	t := &model.ReminderWay{}
	rq.db.Model(t).Find(t, "id = ?", i)
	return t
}
