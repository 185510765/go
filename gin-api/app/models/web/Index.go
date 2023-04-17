package model_web

import (
	"time"

	. "gin-api/app/common/localtime"

	"gorm.io/gorm"
)

type Model struct {
	Id        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type List struct {
	Id          int        `json:"id" gorm:"primary_key"`
	Name        string     `json:"name" gorm:"name"`
	Img         string     `json:"img" gorm:"img"`
	Sort        int        `json:"sort" gorm:"sort"`
	IsShow      int        `json:"is_show" gorm:"is_show"`
	SearchInput string     `json:"search_input" gorm:"search_input"`
	CreatedAt   *LocalTime `json:"created_at" gorm:"created_at"`
	UpdatedAt   *LocalTime `json:"updated_at" gorm:"updated_at"`
	DeletedAt   *LocalTime `json:"deleted_at" gorm:"deleted_at"`
}
