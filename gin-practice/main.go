package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 强制 log 高亮
	// gin.ForceConsoleColor()
	// // 关闭高亮
	// gin.DisableConsoleColor()

	// 配置将日志打印到文件内
	file, _ := os.Create("go-gin-notes.log")
	// gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// 错误日志输出到文件配置 如果想将日志都放到一个文件中，只需要将DefaultWriter和DefaultErrorWriter指定到一个文件上即可
	// errFile, _ := os.Create("go-gin-notes-error.log")
	errFile, _ := os.Create("go-gin-notes.log")
	gin.DefaultErrorWriter = io.MultiWriter(errFile, os.Stdout)

	// 路由
	router := gin.Default()

	// 定义路由日志格式
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	router.GET("/ping", pong)

	// 路由分组
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	// 从url中获取参数
	goodGroup := router.Group("/goods")
	{
		//获取商品id为1的详细信息 模式
		goodGroup.GET("/:id/", goodsDetail)
		goodGroup.GET("/:id/:action", goodsDetailX)
		goodGroup.GET("/:id/:action/add", goodsDetailY)
	}

	// 获取表单参数
	router.GET("/welcome", welcome)

	// post获取参数
	router.POST("/formPost", formPost)

	// 表单验证
	router.POST("/login", login)
	router.POST("/register", register)

	// 模板自定义函数 注册时间格式化方法,特别注意要放到加载模板之前
	router.SetFuncMap(template.FuncMap{
		"DateFormat": DateFormat,
	})

	// 加载模板
	router.LoadHTMLGlob("templates/**/*")
	// router.LoadHTMLFiles("templates/index.tmpl", "templates/goods.html")
	router.GET("/article/index", index)

	// 静态文件挂载方法
	router.Static("/static", "./static")

	router.Run("127.0.0.1:8080")

	// qiut := make(chan os.Signal)
	// //接收control+c
	// //当接收到退出指令时，我们向chan收数据
	// signal.Notify(qiut, syscall.SIGINT, syscall.SIGTERM)
	// <-qiut

	// //服务退出前做处理
	// fmt.Println("服务退出中")
	// fmt.Println("服务已退出")

}

func DateFormat(time time.Time) string {
	return time.Format("2006 01 02")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "article/index.html", gin.H{
		"code": http.StatusOK,
		// "msg":  "<p>该项目是gin框架的学习笔记</p>",
		"msg": "2023-02-3",
	})
}

type registerForm struct {
	User       string `form:"user" json:"user" xml:"user" binding:"required,min=5,max=10"`
	Password   string `form:"password" json:"password" xml:"password"  binding:"required"`
	RePassword string `form:"re_password" json:"re_password" xml:"re_password"  binding:"required,eqfield=Password"` //跨字段
	Age        string `form:"age" json:"age" xml:"age"  binding:"gte=1,lte=130"`
	Email      string `form:"email" json:"email" xml:"email" binding:"required,email"`
}

func register(c *gin.Context) {
	var registerForm registerForm
	err := c.ShouldBind(&registerForm)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

type loginForm struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required,min=5,max=10"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func login(c *gin.Context) {
	var loginForm loginForm
	// loginForm := LoginForm{}
	err := c.ShouldBind(&loginForm)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

func formPost(c *gin.Context) {
	msg := c.PostForm("msg")
	nick := c.DefaultPostForm("nick", "someone")
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"nick": nick,
	})
}

func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "zhangsan")
	// lastname := c.DefaultQuery("lastname", "test")
	// 不加默认值的话就直接使用Query
	lastname := c.Query("lastname")
	c.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
	})
}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func goodsDetailX(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func goodsDetailY(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"msg": "pong",
	})
}

func loginEndpoint(c *gin.Context) {
	// // 定义一个包含多个元素的数组
	// items := []string{"item1", "item2", "item3"}

	// // 使用 c.JSON 返回数组
	// c.JSON(http.StatusOK, gin.H{"items": items})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "loginEndpoint",
		// "data": map[string]interface{}{},
		"data": []string{},
	})
}
func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"msg": "submitEndpoint",
	})
}
func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"msg": "readEndpoint",
	})
}
