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

	// operate := router.Group("/operate")
	// {
	// 	operate.GET("/getSearchInfo", GetSearchInfo)
	// }

	router.Run("127.0.0.1:8080")
}
