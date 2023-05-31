package model_web

import (
	. "gin-api/app/common/localtime"
)

type EmailCode struct {
	Id       int        `json:"id" gorm:"id"`
	Email    string     `json:"email" gorm:"email"`
	Code     string     `json:"code" gorm:"code"`
	SendTime *LocalTime `json:"send_time" gorm:"send_time"`
	UseTime  *LocalTime `json:"use_time" gorm:"use_time"`
}
