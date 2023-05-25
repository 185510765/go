package controller_web

import (
	"gin-api/app/common/response"
	// "gin-api/app/response/response"

	"github.com/gin-gonic/gin"
	// . "gin-api/app/services/web"
)

// 发送邮件验证码
func GetEmailCaptcha(c *gin.Context) {
	response.Ok(c)

	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "GetEmailCaptcha",
	// })
}

// 注册

// 登录
