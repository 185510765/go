package controller_web

import (
	"fmt"
	"gin-api/app/common/common"
	"gin-api/app/common/response"
	. "gin-api/app/services/web"
	"gin-api/config"

	"github.com/gin-gonic/gin"
)

// 发送邮件验证码
func GetEmailCaptcha(c *gin.Context) {
	email := c.DefaultPostForm("email", "")

	// 校验邮箱
	status, msg := ValidateEmail(email)
	if status == 0 {
		response.FailWithMessage(msg, c)
		return
	}

	// 配置
	emailConfig := config.EmailConfig()
	servername := fmt.Sprint(emailConfig["servername"])
	from := fmt.Sprint(emailConfig["from"])
	password := fmt.Sprint(emailConfig["password"])
	to := email

	appConfig := config.App()
	app_name := fmt.Sprint(appConfig["app_name"])
	domain := fmt.Sprint(appConfig["domain"])

	subject := app_name + "注册验证码"
	body := "<p>您好！欢迎注册" + app_name + "</p>" +
		"<p>您的验证码是：{activeCode}，10分钟内有效，如果您未注册，请忽略此消息~</p>" +
		"<p>官网网址：" + domain + "</p>"
	if err := common.SendMail(servername, from, password, to, subject, body, "html"); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// response.OkWithData([]interface{}{
	// 	gin.H{
	// 		"email":        email,
	// 		"emailIsExist": "",
	// 	},
	// 	gin.H{
	// 		"code": 2,
	// 		"msg":  "2",
	// 	}}, c)
}

// 注册

// 登录
