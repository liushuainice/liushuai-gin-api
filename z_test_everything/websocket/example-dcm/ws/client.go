package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'} //换行
	space   = []byte{' '}  //空格
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub            //连接ws服务
	conn *websocket.Conn //自己的连接
	send chan []byte
}
type Msg struct {
	Info string `json:"info"`
	In   int    `json:"in"`
}

func (c *Client) readPump() {
	defer func() {
		fmt.Println("readPump stop,", c.conn)
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) //读取连接有效期延时，超时失效返回err
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage() //读取到信息
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("clent:readPump:error: ", err)
			}
			break
		}
		//信息去掉里换行和空格，然后给里ws的广播服务，最后ws会群发到每个cli
		fmt.Println(11111, "readPump : ", string(message))
		//===========页面发送{"info":"aaa","in":666}可以解析
		msg := Msg{}
		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Println("msg json err")
		} else {
			fmt.Println(0, msg, 0)
		}
		//===========
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}
func (c *Client) writePump(ssssss string) {
	ticker := time.NewTicker(pingPeriod) //定时器-间隔
	defer func() {
		fmt.Println("writePump stop", c.conn)
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send: //ws服务将msg给每个cli，cli将信息发送
			fmt.Println(11111, "writePump : ", string(message), ok)

			//=============
			msg := Msg{}
			if err := json.Unmarshal(message, &msg); err != nil {
				fmt.Println("msg err")
			} else {
				fmt.Println(0, msg, 0)
				msg.Info = msg.Info + " is ok"
				mess, er := json.Marshal(msg)
				if er == nil {
					message = mess
				}
			}
			//=============

			c.conn.SetWriteDeadline(time.Now().Add(writeWait)) //写入连接有效期延时，超时失效返回err
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message) //这里就把信息拼接好了
			//w.Write([]byte("end"))
			//下面的for知道有啥具体作用,注释了也不影响什么
			/*	n := len(c.send)
				for i := 0; i < n; i++ { //拼接管道里的信息
					w.Write(newline)
					w.Write(<-c.send)
				}*/

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C: //定时器启动
			fmt.Println("连接心跳数据", c.conn)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println("下线时触发return", time.Now().Format("2019.01.02.15.04.05"))
				return
			}
		}
	}
}
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err, "client:ServeWs")
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client //ws收到注册信息
	go client.writePump("abc")    //abc 是没用的
	go client.readPump()
}
