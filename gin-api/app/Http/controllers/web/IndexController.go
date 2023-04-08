package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

// 主页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		// "code": http.StatusOK,
		// "msg":  "<p>该项目是gin框架的学习笔记</p>",
	})
}

// 条形码查询页面
func Barcode(c *gin.Context) {
	c.HTML(http.StatusOK, "barcode.html", gin.H{})
}
