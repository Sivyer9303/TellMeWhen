package model

import (
	"time"

	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	ReminderTypeId  int       `json:"reminderTypeId"`
	PlanRemindTime  time.Time `gorm:"plan_remind_time",json:"planRemindTime"`
	CircleStartTime time.Time `gorm:"circle_start_time",json:"circleStartTime"`
	CircleEndTime   time.Time `gorm:"circle_end_time",json:"circleEndTime"`
	ReminderWayId   int       `json:"reminderWayId"`
	ReminderPerson  string    `gorm:"person",json:"reminderPerson"`
}

func (Reminder) TableName() string {
	return "reminder"
}

//type ReminderType struct {
//	gorm.Model
//	Name  string `gorm:"type_name",json:"name"`
//	Desc  string `gorm:"type_desc",json:"desc"`
//	Param string
//}
//
//func (ReminderType) TableName() string {
//	return "reminder_type"
//}

type ReminderWay struct {
	gorm.Model
	Name    string `gorm:"name",json:"name"`
	Desc    string `gorm:"desc",json:"desc"`
	Params  string `gorm:"params",json:"params"`
	Type    string `gorm:"type"`
	WayName string `gorm:"way_name"`
	// `last_reminder_time` datetime DEFAULT NULL COMMENT '上次提醒时间',
	//  `reminder_count` bigint(255) DEFAULT NULL COMMENT '总计提醒次数',
	LastReminderTime time.Time `gorm:"last_reminder_time"`
	ReminderCount    int       `gorm:"reminder_count"`
}

func (ReminderWay) TableName() string {
	return "reminder_way"
}
