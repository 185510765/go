package main

import (
	config "gin-api/config"
	routes "gin-api/routes"
)

func main() {
	// config设置
	config.Log()
	config.App()
	config.Database()

	// 路由
	routes.Web()

}
