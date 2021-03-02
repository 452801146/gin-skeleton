package routers

import (
	"gin_skeleton/routers/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 公共中间件
	// r.Use(middleware.Common())
	// 跨域支持
	r.Use(middleware.Cors())
	// ip限速qps
	r.Use(middleware.IpRateLimit(10))
	// 错误路由
	r.NoRoute(middleware.MissRouter())

	//注册用户路由
	Register(r)

}

func Register(r *gin.Engine) {



}