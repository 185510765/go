package service_web

import (
	"errors"
	"fmt"
	"gin-api/app/common/cache"
	"gin-api/app/common/common"
	"gin-api/app/common/db"
	. "gin-api/app/common/localtime"
	. "gin-api/app/models/web"
	"strconv"
	"time"

	// . "gin-api/app/models/web"
	"gin-api/config"

	"gorm.io/gorm"
)

type UserService struct{}

var localTime LocalTime

// 发送邮件验证码
func (userService *UserService) GetEmailCaptcha(email string) int {
	// 配置
	emailConfig := config.EmailConfig()
	servername := fmt.Sprint(emailConfig["servername"])
	from := fmt.Sprint(emailConfig["from"])
	password := fmt.Sprint(emailConfig["password"])
	to := email

	appConfig := config.App()
	app_name := fmt.Sprint(appConfig["app_name"])
	domain := fmt.Sprint(appConfig["domain"])
	register_email_code_expire := fmt.Sprint(appConfig["register_email_code_expire"])

	// 生成随机验证码
	code := common.RandomCode("int", 6)

	// 验证码存redis
	redisUserRegKey := "userRegCode:" + email
	expireInt, _ := strconv.Atoi(register_email_code_expire)
	cache.RedisClient.Set(redisUserRegKey, code, time.Duration(expireInt*60)*time.Second)

	subject := app_name + "注册验证码"
	body := "<p>您好！欢迎注册" + app_name + "</p>" +
		"<p>您的验证码是：<strong>" + code + "</strong>， " + register_email_code_expire + "分钟内有效，如果您未注册，请忽略此消息~</p>" +
		"<p>官网网址：" + domain + "</p>"
	if err := common.SendMail(servername, from, password, to, subject, body, "html"); err != nil {
		return 0
	}

	return 1
}

// 校验注册数据
func (userService *UserService) ValidateRegister(regParams RegisterParams) (int, string) {
	// 判断验证码
	redisUserRegKey := "userRegCode:" + regParams.Email
	code, _ := cache.RedisClient.Get(redisUserRegKey).Result()
	if code == "" || code != regParams.Captcha {
		return 0, "验证码不正确或已过期"
	}

	// 用户名唯一
	user := User{}
	userRes := db.Model.Where("username = ?", regParams.Username).First(&user)
	if !errors.Is(userRes.Error, gorm.ErrRecordNotFound) {
		return 0, "用户名已存在，请更换用户名"

	}

	// 邮箱唯一
	userEmailRes := db.Model.Where("email = ?", regParams.Email).First(&user)
	if !errors.Is(userEmailRes.Error, gorm.ErrRecordNotFound) {
		return 0, "邮箱已存在，请更换邮箱或找回密码"
	}

	return 1, ""
}

// 校验成功后操作
func (userService *UserService) RegisterToDo(regParams RegisterParams) {
	// 清除redis
	redisUserRegKey := "userRegCode:" + regParams.Email
	cache.RedisClient.Del(redisUserRegKey).Result()

	// 添加用户
	location, _ := time.LoadLocation("Asia/Shanghai")
	cTime := time.Now().In(location)
	db.Model.Create(&User{
		Username:     regParams.Username,
		Password:     regParams.Password,
		Email:        regParams.Email,
		RegisterTime: &cTime,
	})
}

// **********************************************************************************************************************
