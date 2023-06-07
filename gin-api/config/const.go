package config

const (
	// app信息
	APP_NAME = "anyApi"         // 项目名称
	DOMAIN   = "www.anyApi.com" // 域名

	// 邮件信息
	SERVER_NAME = "smtp.qq.com:465"
	FROM        = "185510765@qq.com"
	PASSWORD    = "ljvirisgnujpbjgj"

	// redis key值
	USER_REG_CODE        = "user_reg_code:" // 用户注册邮件验证码前缀
	USER_REG_CODE_EXPIRE = "10"             // 注册验证码过期时间 单位：分
)
