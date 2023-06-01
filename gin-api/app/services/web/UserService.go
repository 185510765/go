package service_web

import (
	"fmt"
	"gin-api/app/common/cache"
	"gin-api/app/common/common"
	"gin-api/app/common/validate"
	"strconv"
	"time"

	// . "gin-api/app/models/web"
	"gin-api/config"
)

type UserService struct{}

// 校验email
func (userService *UserService) ValidateEmail(email string) (int, string) {
	if email == "" {
		return 0, "邮箱不能为空"
	}

	isEmail := validate.IsEmail(email)
	if !isEmail {
		return 0, "邮箱格式不正确"
	}

	return 1, ""
}

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
	redisRegCodeKey := "userRegCode:" + email
	expireInt, _ := strconv.Atoi(register_email_code_expire)
	cache.RedisClient.Set(redisRegCodeKey, code, time.Duration(expireInt*60)*time.Second)

	subject := app_name + "注册验证码"
	body := "<p>您好！欢迎注册" + app_name + "</p>" +
		"<p>您的验证码是：<strong>" + code + "</strong>， " + register_email_code_expire + "分钟内有效，如果您未注册，请忽略此消息~</p>" +
		"<p>官网网址：" + domain + "</p>"
	if err := common.SendMail(servername, from, password, to, subject, body, "html"); err != nil {
		return 0
	}

	return 1
}

// **********************************************************************************************************************
