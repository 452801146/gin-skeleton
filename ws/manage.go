package ws

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type ClientManage struct {
	Clients sync.Map
	Online  chan *Client
	Offline chan *Client
}

var Manage = &ClientManage{
	Clients: sync.Map{},
	Online:  make(chan *Client),
	Offline: make(chan *Client),
}

// 服务启动
func (manage *ClientManage) Start() {
	// 1.启动manage
	go func() {
		for {
			select {
			case client := <-manage.Online:
				// 如果已存在client,需要先关闭原来的连接
				existsClient, hasExists := manage.Clients.Load(client.Uid)
				if hasExists {
					oldClient := existsClient.(*Client)
					manage.OfflineHandler(oldClient)
				}

				// 添加到连接列表
				manage.Clients.Store(client.Uid, client)

			case client := <-manage.Offline:
				manage.OfflineHandler(client)
			}
		}
	}()
}

// 管理连接离线操作
func (manage *ClientManage) OfflineHandler(c *Client) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("offline handler err:", err)
		}
	}()
	c.CloseOnce.Do(func() {
		// 1.关闭通道&连接
		close(c.Send)
		err := c.Conn.Close()
		if err != nil {
			log.Println(err)
		}
		// 2.删除在线列表
		manage.Clients.Delete(c.Uid)
		// 3.在线时长统计,兼容夸天统计
		//c.StatOnline()

	})

}

// gin连接到ws
func Connect() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := (&websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			}}).Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			// 连接失败处理
			c.JSON(http.StatusOK, gin.H{})
			c.Abort()
			log.Println(err)
			return
		}
		// jwt鉴权,在中间件有做abort操作
		uidInterface, exists := c.Get("uid")
		if !exists {
			c.Abort()
			return
		}
		uid := uidInterface.(uint)
		// 注册服务
		client := &Client{
			Uid:         uid,
			Conn:        *conn,
			Send:        make(chan []byte),
			ConnectTime: time.Now(),
			LastTime:    time.Now(),
			CloseOnce:   sync.Once{},
		}
		Manage.Online <- client

		// 启动客户端读&写服务
		go client.Read()
		go client.Write()
	}
}

// 群发信息
func SendDataByUids(uids []uint, msg []byte) {
	for _, uid := range uids {
		SendDataByUid(uid, msg)
	}
}

// 单独发送信息
func SendDataByUid(uid uint, msg []byte) {
	client, ok := Manage.Clients.Load(uid)
	if !ok {
		return
	}
	clientCov := client.(*Client)
	clientCov.Send <- msg
}

// 发送给所有在线
func SendAllOnline(msg []byte) {
	Manage.Clients.Range(func(key, value interface{}) bool {
		client, ok := value.(*Client)
		if !ok {
			return true
		}
		SendDataByUid(client.Uid, msg)
		return true
	})
}

// 格式协议输出
func FormatWsResponse(p string, data ...interface{}) []byte {
	var rsData []interface{}
	if len(data) > 0 {
		rsData = make([]interface{}, 2)
		rsData[0] = p
		rsData[1] = data[0]
	} else {
		rsData = make([]interface{}, 1)
		rsData[0] = p
	}
	rsByte, _ := json.Marshal(rsData)
	return rsByte
}

// 获取client内存地址
// 用于对client相关修改操作
func GetClient(uid uint) (c *Client, err error) {
	clientInterface, ok := Manage.Clients.Load(uid)
	if !ok {
		err = errors.New("client no exists")
		return
	}
	client, ok := clientInterface.(*Client)
	if !ok {
		err = errors.New("client cv err")
		return
	}
	c = client
	return
}
