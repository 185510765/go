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
	var params EmailParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(validator.Translate(err), c)
		return
	}

	// 发送
	resCode := userService.GetEmailCaptcha(params.Email)
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

	// 校验
	status, msg := userService.ValidateRegister(regParams)
	if status == 0 {
		response.FailWithMessage(msg, c)
		return
	}

	// 校验成功后操作
	userService.RegisterToDo(regParams)

	response.Ok(c)
}

// 登录
func Login(c *gin.Context) {}
