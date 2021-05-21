package ws

import (
	"gin_skeleton/constant"
	"log"
	"sync"
)

// 处理程序
type Func func(uid uint, p string, data string)

// 路由列表
type ClientHandler struct {
	Handlers map[string]Func
	SyMu     sync.Mutex
}

// 路由实例
var HandlerRouter = &ClientHandler{
	Handlers: make(map[string]Func),
}

// 执行程序
// 防止panic,增加处理
func (handler *ClientHandler) Run(uid uint, p string, data string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("client handler err:", err)
		}
	}()
	handler.SyMu.Lock()
	clientHandler, ok := handler.Handlers[p]
	handler.SyMu.Unlock()
	if !ok {
		SendDataByUid(uid, FormatWsResponse(constant.P_handler_err))
		return
	}
	clientHandler(uid, p, data)
}

// 添加路由
func (handler *ClientHandler) Add(p string, f Func) {
	handler.SyMu.Lock()
	handler.Handlers[p] = f
	handler.SyMu.Unlock()
}
