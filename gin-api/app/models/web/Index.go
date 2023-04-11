package model_web

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type List struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"name"`
	Img       string    `json:"img" gorm:"img"`
	Sort      int       `json:"sort" gorm:"sort"`
	IsShow    int       `json:"is_show" gorm:"is_show"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
}
