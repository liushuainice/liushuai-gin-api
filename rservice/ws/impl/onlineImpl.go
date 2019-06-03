package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/json-iterator/go"
	"liushuai-gin-api/app"
	"liushuai-gin-api/rservice/ws"
	"liushuai-gin-api/util/golog"
	"sync"
	"text/template"
	"time"
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
	conn *websocket.Conn
}

// online 在线相关逻辑服务
type onlinex struct {
	sockets        map[string]*websocket.Conn // 在线信息
	chanAddUser    chan *onlineUser           // 用户上线
	chanRemoveUser chan string                // 用户下线
	chanPushMsg    chan *onlineMessage        // 推消息
	chanIsOnline   chan *isOnlineReq          // 是否在线
}

type onlineUser struct {
	UID    string
	Socket *websocket.Conn
}
type isOnlineReq struct {
	UIDs []string
	Resp chan map[string]bool
}
type onlineMessage struct {
	UID     string `json:"uid"`
	Event   string `json:"event"`
	Content string `json:"content"`
}

var once sync.Once
var singleton *onlinex

// NewOnline 创建
func NewOnline() ws.OnlineService {
	log.Println(1)
	once.Do(func() {
		singleton = &onlinex{
			sockets:        make(map[string]*websocket.Conn),
			chanAddUser:    make(chan *onlineUser),
			chanRemoveUser: make(chan string),
			chanPushMsg:    make(chan *onlineMessage),
			chanIsOnline:   make(chan *isOnlineReq),
		}
		singleton.Run()
	})
	//这里注意，return的是结构体，而要返回的参数是接口，这里相当于
	// ws.OnlineService=singleton
	// return ws.OnlineService
	//singleton 是私有结构体，无法传递出去，外部只能用接口调用
	return singleton
}

func (s *onlinex) Run() {
	go func() {
		for {
			select {
			case onlineUser := <-s.chanAddUser:
				s.processAddUserx(onlineUser)
			case uid := <-s.chanRemoveUser:
				s.processRemoveUserx(uid)
			case req := <-s.chanIsOnline:
				s.processIsOnline(req)
			case msg := <-s.chanPushMsg:
				s.processPushMsg(msg)
			}
		}
	}()
}

//========================== websocket ========================== websocket ==========================
func ServeWsV(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err, "client:ServeWs")
		return
	}
	client := &Client{conn: conn}
	go client.connPump()
	go client.readPump()
}

//跳转测试页面
func ServeHomeV(c *gin.Context) {
	t, _ := template.ParseFiles("public/home.html")
	type Parm struct {
		Info string
	}
	t.Execute(c.Writer, Parm{"hello ws"})
}
func (c *Client) readPump() {
	var uid string
	defer func() {
		fmt.Println("readPump stop,", c.conn)
		app.ServiceOnline.RemoveUserx(uid) //注销
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) //读取连接有效期延时，超时失效返回err
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	f := true
	for {
		_, message, err := c.conn.ReadMessage() //读取到信息
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("client:readPump:error: ", err)
			}
			break
		}

		//fmt.Println(11111, "readPump : ", string(message))
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1)) //没看出具体效果
		//fmt.Println(22222, "readPump : ", string(message))

		msg := onlineMessage{}
		jsonx := jsoniter.ConfigCompatibleWithStandardLibrary //一个高效的json包，比原生的好
		er := jsonx.Unmarshal(message, &msg)
		if er == nil {
			if f { //注册
				uid = msg.UID
				app.ServiceOnline.AddUserx(uid, c.conn)
				f = false
			}
			//下面 做业务逻辑
			app.ServiceOnline.Broadcastx(msg.Event, msg.Content)
			app.ServiceOnline.Pushx(msg.UID, msg.Event, msg.Content)
		}
	}
}
func (c *Client) connPump() {
	ticker := time.NewTicker(pingPeriod) //定时器-间隔
	defer func() {
		fmt.Println("connPump stop", c.conn)
		ticker.Stop()
		c.conn.Close() //主要是这句，下线关闭连接
	}()
	for {
		select {
		case <-ticker.C: //定时器启动
			fmt.Println("心跳连接", time.Now().Format("15.04.05"), c.conn)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//============================  下面是接口业务逻辑--接口封装单方法  =====================================

// AddUser 上线
func (s *onlinex) AddUserx(uid string, socket *websocket.Conn) {
	s.chanAddUser <- &onlineUser{UID: uid, Socket: socket}
}

// RemoveUser 下线
func (s *onlinex) RemoveUserx(uid string) {
	s.chanRemoveUser <- uid
}

// GetAllUsersOnlinex 全部在线人员
func (s *onlinex) GetAllUsersOnlinex() map[string]bool {
	req := &isOnlineReq{UIDs: []string{}, Resp: make(chan map[string]bool)}
	s.chanIsOnline <- req
	return <-req.Resp
}

// IsUserOnline 多人在线判断
func (s *onlinex) IsUsersOnlinex(uids []string) map[string]bool {
	req := &isOnlineReq{UIDs: uids, Resp: make(chan map[string]bool)}
	s.chanIsOnline <- req
	return <-req.Resp
}

// IsUserOnline 单人在线判断
func (s *onlinex) IsUserOnlinex(uid string) bool {
	req := &isOnlineReq{UIDs: []string{uid}, Resp: make(chan map[string]bool)}
	s.chanIsOnline <- req
	res := <-req.Resp
	return res[uid]
}

// Push 推送消息-单发
func (s *onlinex) Pushx(uid string, event string, content string) {
	s.chanPushMsg <- &onlineMessage{UID: uid, Event: event, Content: content}
}

// Broadcast 推送消息-群发
func (s *onlinex) Broadcastx(event string, content string) {
	s.chanPushMsg <- &onlineMessage{Event: event, Content: content}
}

//=========================  下面是接口业务逻辑--底层与ws操作的方法  ========================================

//添加上线连接
func (s *onlinex) processAddUserx(onlineUser *onlineUser) {
	s.sockets[onlineUser.UID] = onlineUser.Socket
}

//移除上线连接
func (s *onlinex) processRemoveUserx(uid string) {
	delete(s.sockets, uid)
}

//判断连接是否存在
func (s *onlinex) processIsOnline(req *isOnlineReq) {
	result := make(map[string]bool)
	if len(req.UIDs) > 0 {
		for _, uid := range req.UIDs {
			_, ok := s.sockets[uid]
			result[uid] = ok
		}
	} else {
		for k := range s.sockets {
			result[k] = true
		}
	}

	req.Resp <- result
}

//推送消息
func (s *onlinex) processPushMsg(msg *onlineMessage) {
	b, e := json.Marshal(msg)
	if e != nil {
		return
	}
	if len(msg.UID) > 0 {
		so, ok := s.sockets[msg.UID]
		if !ok {
			return
		}
		if err := sent(so, b); err != nil {
			so.Close()
			delete(s.sockets, msg.UID)
		}
	} else {
		for _, so := range s.sockets {
			if err := sent(so, b); err != nil {
				so.Close()
				delete(s.sockets, msg.UID)
			}
		}
	}
	/*
		//这段是直接发送json信息段
		fmt.Println(" processPushMsg-->online:   ", 1, msg.UID, 2, msg.Event, 3, msg.Content)
			if msg.UID > 0 {
				so, ok := s.sockets[msg.UID]
				if !ok {
					return
				}
				err := so.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					so.Close()
					delete(s.sockets, msg.UID)
				}
			} else {
				for _, so := range s.sockets {
					err := so.WriteJSON(msg)
					if err != nil {
						log.Printf("error: %v", err)
						so.Close()
					}
				}
			}*/

}
func sent(c *websocket.Conn, b []byte) error {
	c.SetWriteDeadline(time.Now().Add(writeWait))
	if len(b) == 0 {
		c.WriteMessage(websocket.CloseMessage, []byte{})
		return nil
	}
	w, err := c.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	w.Write(b)
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
