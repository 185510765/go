package service_web

import (
	"errors"
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
	// 生成随机验证码
	code := common.RandomCode("int", 6)

	// 验证码存redis
	redisUserRegKey := config.USER_REG_CODE + email
	expireInt, _ := strconv.Atoi(config.USER_REG_CODE_EXPIRE)
	cache.RedisClient.Set(redisUserRegKey, code, time.Duration(expireInt*60)*time.Second)

	subject := config.APP_NAME + "注册验证码"
	body := "<p>您好！欢迎注册" + config.APP_NAME + "</p>" +
		"<p>您的验证码是：<strong>" + code + "</strong>， " + config.USER_REG_CODE_EXPIRE + "分钟内有效，如果您未注册，请忽略此消息~</p>" +
		"<p>官网网址：" + config.DOMAIN + "</p>"
	if err := common.SendMail(config.SERVER_NAME, config.FROM, config.PASSWORD, email, subject, body, "html"); err != nil {
		return 0
	}

	return 1
}

// 校验注册数据
func (userService *UserService) ValidateRegister(regParams RegisterParams, rsaDecPwd string, rsaDecConPwd string) (int, string) {
	// 判断验证码
	redisUserRegKey := config.USER_REG_CODE + regParams.Email
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

	// 密码
	pwdLen := len(rsaDecPwd)
	if pwdLen < 6 || pwdLen > 32 {
		return 0, "密码长度在6至32位之间"
	}

	// 确认密码
	if rsaDecPwd != rsaDecConPwd {
		return 0, "确认密码和密码不一致"
	}

	// 邮箱唯一
	userEmailRes := db.Model.Where("email = ?", regParams.Email).First(&user)
	if !errors.Is(userEmailRes.Error, gorm.ErrRecordNotFound) {
		return 0, "邮箱已存在，请更换邮箱或找回密码"
	}

	return 1, ""
}

// 校验成功后操作
func (userService *UserService) RegisterToDo(regParams RegisterParams, rsaDecPwd string) {
	// 清除redis
	redisUserRegKey := config.USER_REG_CODE + regParams.Email
	cache.RedisClient.Del(redisUserRegKey).Result()

	// 添加用户
	salt := common.RandomCode("all", 8)                  // 盐值
	password := common.PasswordEncrypte(rsaDecPwd, salt) // 密码

	location, _ := time.LoadLocation("Asia/Shanghai")
	cTime := time.Now().In(location)
	db.Model.Create(&User{
		Username:     regParams.Username,
		Salt:         salt,
		Password:     password,
		Email:        regParams.Email,
		RegisterTime: &cTime,
	})
}

// **********************************************************************************************************************
