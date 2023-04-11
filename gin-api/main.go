package main

import (
	config "gin-api/config"
	routes "gin-api/routes"
)

func main() {
	// // db
	// dsn := "root:@tcp(127.0.0.1:3306)/api?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err) // 异常处理
	// }

	// config设置
	config.Log()
	config.App()
	config.Database()

	// 路由
	routes.Web()

}
