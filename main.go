package main

import (
	"gin_skeleton/g"
	"gin_skeleton/initialize"
)

func main()  {

	//初始化服务器
	r := initialize.CreateServer()

	//启动服务器
	err := r.Run(g.Config.GetString("server.addr"))

	if err != nil {
		panic(err)
	}

}
