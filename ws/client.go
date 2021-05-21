package ws

import (
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"sync"
	"time"
)

type Client struct {
	Uid         uint
	Conn        websocket.Conn
	Send        chan []byte
	ConnectTime time.Time
	LastTime    time.Time
	CloseOnce   sync.Once
}

// 解析client发送过来的数据
// 数据格式定义,json格式
// ex:{p:101,data{f1:1,f2:2}}
// p协议号, data,提交数据

func (c *Client) ParserData(data []byte) {
	// 更新client时间
	c.LastTime = time.Now()

	// todo:此处可处理限流,限速,其他中间件

	// 解析数据
	p := gjson.GetBytes(data, "0").String()
	formData := gjson.GetBytes(data, "1").String()
	HandlerRouter.Run(c.Uid, p, formData)

}

// 读取client发送过来的数据
func (c *Client) Read() {
	// 离线关闭client
	defer func() {
		Manage.Offline <- c
	}()
	for {
		_, revData, err := c.Conn.ReadMessage()
		if err != nil {
			// log.Println("read err:",err)
			return
		}
		c.ParserData(revData)
	}
}

// 对client写入数据
func (c *Client) Write() {
	// 离线关闭client
	defer func() {
		Manage.Offline <- c
	}()
	for {
		select {
		case sendData := <-c.Send:
			err := c.Conn.WriteMessage(websocket.TextMessage, sendData)
			if err != nil {
				// log.Println("write err:",err)
				return
			}
		}
	}
}
