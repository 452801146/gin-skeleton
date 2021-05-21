package main

import (
	"fmt"
	"gin_skeleton/g"
	"gin_skeleton/initialize"
)

func main() {

	//初始化服务器
	r := initialize.CreateServer()

	fmt.Println(g.Redis.ClientList().String())
	//启动服务器
	err := r.Run(g.Config.GetString("server.addr"))
	if err != nil {
		panic(err)
	}

}
