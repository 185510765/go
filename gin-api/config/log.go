package config

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Log() {
	// 强制 log 高亮
	// gin.ForceConsoleColor()
	// // 关闭高亮
	// gin.DisableConsoleColor()

	// 配置将日志打印到文件内
	t := time.Now().Unix()
	// date := time.Unix(t, 0).Format("2006-01-02 15:04:05")
	date := time.Unix(t, 0).Format("2006-01-02")
	file, _ := os.Create("storage/logs/go-gin-notes-" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// 错误日志输出到文件配置 如果想将日志都放到一个文件中，只需要将DefaultWriter和DefaultErrorWriter指定到一个文件上即可
	errFile, _ := os.Create("storage/logs/go-gin-error-" + date + ".log")
	gin.DefaultErrorWriter = io.MultiWriter(errFile, os.Stdout)
}
