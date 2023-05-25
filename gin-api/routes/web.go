package routes

import (
	. "gin-api/app/http/controllers/web"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Web() {
	router := gin.Default()

	// 加载模板
	// router.LoadHTMLGlob("resources/views/web/*.html")
	// router.LoadHTMLFiles("./resources/views/web/api_list.html", "./resources/views/web/search.html", "./resources/views/web/register.html", "./resources/views/web/templates/layout.html")

	// 加载./resources/views目录下所有html文件
	var files []string
	filepath.Walk("./resources/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)

	// 静态文件挂载方法
	router.Static("/public", "./public")

	// 路由
	router.GET("/ping", Pong)
	router.GET("/index", Index)

	// api列表 页面
	api_list := router.Group("/api_list")
	{
		api_list.GET("", ApiList)
		api_list.GET("/search/:id", Search)
		api_list.GET("/doc/:id", Doc)
	}

	// 注册登录 验证码
	user := router.Group("/user")
	{
		user.POST("/getEmailCaptcha", GetEmailCaptcha)
		user.POST("/register", Register)
		user.POST("/login", Login)
	}

	router.Run("127.0.0.1:8080")
}
