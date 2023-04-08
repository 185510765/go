package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.HTML(http.StatusOK, "api_list.html", gin.H{})
}

// api查询页面
func Search(c *gin.Context) {
	id := c.Param("id")

	c.HTML(http.StatusOK, "search.html", gin.H{
		"id":    id,
		"title": "",
	})
}

// api文档页面
func Doc(c *gin.Context) {
	id := c.Param("id")

	c.HTML(http.StatusOK, "doc.html", gin.H{
		"id":    id,
		"title": "",
	})
}

// *****************************************************************************************
// 查询api列表数据
func getApiList() {

}
