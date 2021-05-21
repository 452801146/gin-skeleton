package ws_handler

import (
	"gin_skeleton/ws"
)

// 维持心跳
func Ping(uid uint, p string, data string) {

	ws.SendDataByUid(uid, ws.FormatWsResponse(p))
}
