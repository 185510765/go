package main

import (
	"log"
	"net/http"
	routes "web/routes"
)

func main() {
	// 路由访问
	routes.Web()

	// web http服务器
	err := http.ListenAndServe("127.0.0.1:9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// res := &common.Result{
	// 	Code: 200,
	// 	Msg:  "可以",
	// 	Data: make([]int, 0),
	// }

	// fmt.Println(res.Output(res.Code, res.Msg, res.Data))
	// fmt.Println(res.Success(make([]int, 0)))
	// fmt.Println(res.Error("error"))
}
