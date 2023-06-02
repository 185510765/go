package controller_web

import (
	"gin-api/app/common/response"
	"gin-api/app/common/validator"
	. "gin-api/app/services/web"

	. "gin-api/app/models/web"

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
}

// 注册
func Register(c *gin.Context) {
	var regParams RegisterParams
	if err := c.ShouldBind(&regParams); err != nil {
		response.FailWithMessage(validator.Translate(err), c)
		return
	}

	response.OkWithData(gin.H{
		"username":         regParams.Username,
		"password":         regParams.Password,
		"confirm_password": regParams.ConfirmPassword,
		"email":            regParams.Email,
		"captcha":          regParams.Captcha,
	}, c)
	return

	// ******************************************************

	username := c.PostForm("username")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")
	email := c.PostForm("email")
	captcha := c.PostForm("captcha")

	// 校验数据
	status, msg := userService.ValidateRegister()
	if status == 0 {
		response.FailWithMessage(msg, c)
		return
	}

	response.OkWithData(
		gin.H{
			"username":         username,
			"password":         password,
			"confirm_password": confirm_password,
			"email":            email,
			"captcha":          captcha,
		}, c)
}

// 登录
func Login(c *gin.Context) {}
