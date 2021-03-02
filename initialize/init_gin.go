package initialize

import (
	"gin_skeleton/g"
	"github.com/gin-gonic/gin"
)

func InitGin()  {

	// 日志写入文件
	writer := GetLogWriter("server.ginLog")
	gin.DefaultWriter = writer
	// 设置gin运行模式
	gin.SetMode(g.Config.GetString("server.ginModel"))
	// 屏蔽命令行颜色
	gin.DisableConsoleColor()
}