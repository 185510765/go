package routes

import (
	"net/http"
	controller "web/app/http/controller"
)

func Web() {
	http.HandleFunc("/", controller.Index)      // 首页
	http.HandleFunc("/login", controller.Login) // 登录
}
