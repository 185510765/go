package controller_web

import (
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
	info := GetOneList(int_id)

	// 查询限制（根据ip），限制查询频率（5秒）、一天中总的查询次数（100次）
	// ip := c.ClientIP()

	// searchInput := c.Query("searchInput")

	c.HTML(http.StatusOK, "search.html", gin.H{
		"info": info,
	})
}

// api文档页面
func Doc(c *gin.Context) {
	id := c.Param("id")
	int_id, _ := strconv.Atoi(id)
	info := GetOneList(int_id)

	c.HTML(http.StatusOK, "doc.html", gin.H{
		"info": info,
	})
}

// // 搜索操作
// func GetSearchInfo(c *gin.Context) {
// 	fmt.Println("GetSearchInfo")
// }
