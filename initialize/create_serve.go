package initialize

import (
	"gin_skeleton/routers"
	"github.com/gin-gonic/gin"
)

func InitServe() {

	// 基础初始化
	InitLogDir()
	// 初始化配置
	InitConfig()
	// 初始化日志
	InitLogger()
	// 初始化orm
	//InitGorm()
	// 初始化redis
	//InitRedis()
	// 初始化定时器
	InitTask()
	// gin初始化配置
	InitGin()

}

// 启动服务器
func CreateServer() *gin.Engine {
	// 初始化
	InitServe()
	// gin Server
	r := gin.Default()
	// 注册路由
	routers.InitRouter(r)
	return r
}
