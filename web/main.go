package main

import (
	"log"
	"net/http"
	"web/routes"
)

func main() {
	// 路由访问
	routes.Web()

	// web http服务器
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
