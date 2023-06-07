package model_web

import (
	// . "gin-api/app/common/localtime"
	"time"
)

type User struct {
	// gorm.Model

	ID               uint       `json:"id" gorm:"id"`
	Username         string     `json:"username" gorm:"username"`
	Salt             string     `json:"salt" gorm:"salt"`
	Password         string     `json:"password" gorm:"password"`
	RelaName         string     `json:"rela_name" gorm:"rela_name"`
	Email            string     `json:"email" gorm:"email"`
	Phone            string     `json:"phone" gorm:"phone"`
	Qq               string     `json:"qq" gorm:"qq"`
	RegisterIpString string     `json:"register_ip_string" gorm:"register_ip_string"`
	RegisterIpInt    int        `json:"register_ip_int" gorm:"register_ip_int"`
	RegisterTime     *time.Time `json:"register_time" gorm:"register_time"`
	Status           int8       `json:"status" gorm:"status"`
	LoginErrorCount  int8       `json:"login_error_count" gorm:"login_error_count"`
	LoginErrorTime   *time.Time `json:"login_error_time" gorm:"login_error_time"`
	Remark           string     `json:"remark" gorm:"remark"`
	// CreatedAt      *LocalTime `json:"created_at" gorm:"created_at"`
	// UpdatedAt      *LocalTime `json:"updated_at" gorm:"updated_at"`
	// DeletedAt      *LocalTime `json:"deleted_at" gorm:"deleted_at"`
}

// form参数 ***********************************************************************************

type EmailParams struct {
	Email string `form:"email" json:"email" gorm:"email" binding:"required,email" label:"邮箱"`
}

type RegisterParams struct {
	Username        string `form:"username" json:"username" gorm:"username" binding:"required,min=5,max=50" label:"用户名"`
	Password        string `form:"password" json:"password" gorm:"password" binding:"required" label:"密码"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" gorm:"confirm_password" binding:"required" label:"确认密码"`
	Email           string `form:"email" json:"email" gorm:"email" binding:"required,email" label:"邮箱"`
	Captcha         string `form:"captcha" json:"captcha" gorm:"captcha" binding:"required" label:"验证码"`
}
