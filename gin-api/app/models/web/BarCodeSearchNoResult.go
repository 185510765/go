package model_web

import (
	. "gin-api/app/common/localtime"
)

type BarCodeSearchNoResult struct {
	Id        uint       `json:"id" gorm:"id"`
	BarCode   string     `json:"bar_code" gorm:"bar_code"`
	UserId    uint       `json:"user_id" gorm:"user_id"`
	IpString  string     `json:"ip_string" gorm:"ip_string"`
	IpInt     uint32     `json:"ip_int" gorm:"ip_int"`
	Status    uint8      `json:"status" gorm:"status"`
	CreatedAt *LocalTime `json:"created_at" gorm:"created_at"`
	UpdatedAt *LocalTime `json:"updated_at" gorm:"updated_at"`
	DeletedAt *LocalTime `json:"deleted_at" gorm:"deleted_at"`
}

// // TableName 解决gorm表名映射 解决表名后缀加s问题
// func (BarCodeSearchNoResult) TableName() string {
// 	return "bar_code_search_no_result"
// }
