package main

//单对单的聊天websocket的demo
import (
	"fmt"
	logx "ginWebTest/util/golog"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
//=================
var cliMap = make(map[string]*websocket.Conn)

//=================
// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	//====之间的代码 模拟 私聊
	Email    string `json:"email"`    //现在假设这是账号
	Username string `json:"username"` //这是对话的人的账号
	Message  string `json:"message"`  //这是对话内容
}

//socket-可以启动测试-8001
func main() {
	logx.Info(1)
	r := gin.Default()
	route(r)
	r.Run("0.0.0.0:8001")
}

//先有的连接，再有的信息体，可以访问时直接带用户名参数
func handleConnections2(c *gin.Context) {
	fmt.Println(1)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil) //获得连接
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("ws.Close()", 6, ws)
		ws.Close()
	}()
	//判断是否重复连接
	if _, ok := clients[ws]; !ok {
		clients[ws] = true //保存连接
		fmt.Println(ws)
	} else {
		fmt.Println("重复连接")
	}
	//clients[ws] = true //保存连接
	//=================
	flag := true
	//=================
	for { //管道监听信息--》信息给携程，携程做信息处理
		fmt.Println(2)
		var msg Message
		err := ws.ReadJSON(&msg) //ws接收信息
		//=========
		if flag {
			fmt.Println(3)
			cliMap[msg.Email] = ws
			flag = false
		}
		//=========
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			//======
			delete(cliMap, msg.Email)
			//======
			break
		}
		// Send the newly received message to the broadcast channel
		fmt.Println(4)
		broadcast <- msg
		fmt.Println(5)
	}
}
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		fmt.Println("handleMessages-->msg:", 1000, msg.Email, 2000, msg.Message, 3000, msg.Username)
		//=================
		msg1 := msg
		msg1.Message = msg.Message + " chat by " + msg.Email
		v, ok := cliMap[msg.Username]
		if ok {
			err := v.WriteJSON(msg1)
			if err != nil {
				log.Printf("error: %v", err)
				cliMap[msg.Username].Close()
				delete(cliMap, msg.Username)
			}
		}
		//=================
		for client := range clients { //群发
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func route(r *gin.Engine) {
	r.Static("/public", "./public")
	rr := multitemplate.NewRenderer()
	rr.AddFromFiles("socket", "public/index.html")
	r.HTMLRender = rr

	go handleMessages()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "socket", gin.H{"msg": "hello world xxx"})
	})
	r.GET("/ws", handleConnections2)
	r.GET("/a", func(c *gin.Context) {
		g := c.DefaultQuery("name", "test")
		c.JSON(200, gin.H{"message": g})
	})
}
