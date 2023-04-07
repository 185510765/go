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
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"code": http.StatusOK,
		// "msg":  "<p>该项目是gin框架的学习笔记</p>",
		"msg": "2023-02-3",
	})
}
