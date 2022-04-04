package model

import (
	"gorm.io/gorm"
	"time"
)

type Reminder struct {
	gorm.Model
	ReminderType    ReminderType `json:"reminderType"`
	ReminderTypeId  int          `json:"reminderTypeId"`
	PlanRemindTime  time.Time    `gorm:"plan_remind_time",json:"planRemindTime"`
	CircleStartTime time.Time    `gorm:"circle_start_time",json:"circleStartTime"`
	CircleEndTime   time.Time    `gorm:"circle_end_time",json:"circleEndTime"`
	ReminderWay     ReminderWay  `json:"reminderWay"`
	ReminderWayId   int          `json:"reminderWayId"`
	ReminderPerson  string       `gorm:"person",json:"reminderPerson"`
}

func (Reminder) TableName() string {
	return "reminder"
}

type ReminderType struct {
	gorm.Model
	Name string `gorm:"type_name",json:"name"`
	Desc string `gorm:"type_desc",json:"desc"`
}

func (ReminderType) TableName() string {
	return "reminder_type"
}

type ReminderWay struct {
	gorm.Model
	Name   string `gorm:"name",json:"name"`
	Desc   string `gorm:"desc",json:"desc"`
	Params string `gorm:"params",json:"params"`
}

func (ReminderWay) TableName() string {
	return "reminder_way"
}
