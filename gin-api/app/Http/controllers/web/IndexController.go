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
	info, tips := GetOneList(int_id)

	searchInput := c.Query("searchInput")
	if searchInput == "" {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"info":   info,
			"status": 0,
			"msg":    tips + "不能为空",
		})
		return
	}

	// 查询限制（根据ip），限制查询频率（5秒）、24小时总的查询次数（100次）
	ip := c.ClientIP()
	var keySuffix string = id + ":" + ip
	status, msg := ValidateSearch(c, searchInput, keySuffix)
	if status == 0 {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"info":   info,
			"status": 0,
			"msg":    msg,
		})
		return
	}

	// 查询接口数据

	c.HTML(http.StatusOK, "search.html", gin.H{
		"info":   info,
		"status": status,
		"msg":    msg,
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
