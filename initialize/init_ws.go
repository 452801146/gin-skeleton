package initialize

import "gin_skeleton/ws"

func InitWs() {
	ws.Manage.Start()
}
