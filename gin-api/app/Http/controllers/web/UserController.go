package controller_web

import (
	"gin-api/app/common/response"
	. "gin-api/app/services/web"

	"github.com/gin-gonic/gin"
)

var userService UserService

// 发送邮件验证码
func GetEmailCaptcha(c *gin.Context) {
	email := c.DefaultPostForm("email", "")

	// 校验邮箱
	status, msg := userService.ValidateEmail(email)
	if status == 0 {
		response.FailWithMessage(msg, c)
		return
	}

	// 发送
	resCode := userService.GetEmailCaptcha(email)
	if resCode == 0 {
		response.FailWithMessage("邮件发送失败", c)
		return
	}

	response.Ok(c)

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
