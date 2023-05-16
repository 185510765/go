package model_web

import (
	. "gin-api/app/common/localtime"
)

type BarCode struct {
	Id             int64      `json:"id" gorm:"id"`
	BarCode        string     `json:"bar_code" gorm:"bar_code"`
	Name           string     `json:"name" gorm:"name"`
	ShortName      string     `json:"short_name" gorm:"short_name"`
	Image          string     `json:"image" gorm:"image"`
	Price          float64    `json:"price" gorm:"price"`
	Specification  string     `json:"specification" gorm:"specification"`
	Brand          string     `json:"brand" gorm:"brand"`
	Supplier       string     `json:"supplier" gorm:"supplier"`
	Classification string     `json:"classification" gorm:"classification"`
	Status         int8       `json:"status" gorm:"status"`
	CreatedAt      *LocalTime `json:"created_at" gorm:"created_at"`
	UpdatedAt      *LocalTime `json:"updated_at" gorm:"updated_at"`
	DeletedAt      *LocalTime `json:"deleted_at" gorm:"deleted_at"`
}
