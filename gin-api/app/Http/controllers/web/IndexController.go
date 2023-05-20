package controller_web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	. "gin-api/app/services/web"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

// 主页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// 登录
func Login(c *gin.Context) {}

// api列表页面
func ApiList(c *gin.Context) {
	list := GetApiList()

	c.HTML(http.StatusOK, "api_list.html", gin.H{
		"list": list,
	})
}

// api查询页面
func Search(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	info, tips := GetOneList(int_id)

	searchInput, inputIsExist := c.GetQuery("searchInput")

	// 查询校验 查询限制（根据ip），限制查询频率（5秒）、24小时总的查询次数（100次），每个类型限制时间不共用
	ip := c.ClientIP()
	var keySuffix string = id + ":" + ip
	status, msg := ValidateSearch(searchInput, inputIsExist, keySuffix, tips)
	if status == 0 {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"searchInput": searchInput,
			"info":        info,
			"status":      0,
			"msg":         msg,
		})
		return
	}

	// 查询操作
	searchRes := QueryData(int_id, searchInput)

	if searchRes["name"] == "" || searchRes["name"] == nil {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"searchInput": searchInput,
			"info":        info,
			"status":      0,
			"msg":         "没有查询到数据",
		})
		return
	}

	// 处理返回数据
	searchInitRes := InitRes(searchRes)
	searchResJsonString, _ := json.Marshal(searchRes)

	c.HTML(http.StatusOK, "search.html", gin.H{
		"searchInput":         searchInput,
		"info":                info,
		"status":              status,
		"msg":                 msg,
		"searchRes":           searchRes,                   // 英文key
		"searchInitRes":       searchInitRes,               // 中文key
		"searchResJsonString": string(searchResJsonString), // json 字符串
	})
}

// api文档页面
func Doc(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	info, tips := GetOneList(int_id)

	c.HTML(http.StatusOK, "doc.html", gin.H{
		"info": info,
		"tips": tips,
	})
}

// // 搜索操作
// func GetSearchInfo(c *gin.Context) {
// 	fmt.Println("GetSearchInfo")
// }
