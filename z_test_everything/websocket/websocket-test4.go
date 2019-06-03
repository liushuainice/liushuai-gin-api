package main

//后台发送公告的websocket的demo
/*import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logx "github.com/panjiang/golog"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
func notice(c *gin.Context) {
	msg:= Message{
		Email:    "email",
		Username: "username",
		Message:  "background message",
	}
	broadcast <- msg
	fmt.Println("=============")
	c.JSON(200, gin.H{"msg":"ok"})
}

//先有的连接，再有的信息体，可以访问时直接带用户名参数
func handleConnections2(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil) //获得连接
	if err != nil {
		logx.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true //保存连接




	for { //管道监听信息--》信息给携程，携程做信息处理
		var msg Message
		err := ws.ReadJSON(&msg) //ws接收信息
		if err != nil {
			logx.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}
func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Println("handleMessages-->msg:", 1000, msg.Email, 2000, msg.Message, 3000, msg.Username)
		for client := range clients { //群发
			err := client.WriteJSON(msg)
			if err != nil {
				logx.Printf("error: %v", err)
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
	r.GET("/notice", notice)
}

//socket-可以启动测试-8001
func main() {
	logx.Info(1)
	r := gin.Default()
	route(r)
	r.Run("0.0.0.0:8001")
}
*/
