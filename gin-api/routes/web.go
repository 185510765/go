package routes

import (
	web "gin-api/app/http/controllers/web"

	"github.com/gin-gonic/gin"
)

func Web() {
	router := gin.Default()

	// 加载模板
	router.LoadHTMLGlob("resources/**/**/*")
	// router.LoadHTMLFiles("resources/view/test.gohtml", ".gohtml")
	// 静态文件挂载方法
	router.Static("/public", "./public")

	// 路由
	router.GET("/ping", web.Pong)
	router.GET("/index", web.Index)

	router.Run("127.0.0.1:8080")
}
