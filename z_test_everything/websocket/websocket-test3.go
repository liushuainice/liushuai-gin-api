package main

/*
//gm指令测试websocket的demo
import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logx "github.com/panjiang/golog"
	"log"
	"net/http"
	"strings"
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
type commendx struct {
	Commend string
	Msg     string
}

func getCommandxList(comm string) func(*commendx) bool {
	maps := map[string]func(*commendx) bool{
		"help": gmHelp,
	}
	return maps[comm]
}
func gmHelp(command *commendx) bool {
	command.Msg = command.Msg + `  test--commdend`
	return true
}
func handleConnections2(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil) //获得连接
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true //保存连接
	for { //管道监听信息--》信息给携程，携程做信息处理
		var msg Message
		err := ws.ReadJSON(&msg) //ws接收信息
		if err != nil {
			log.Printf("error: %v", err)
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
		//=================
		//gm指令测试
		split := strings.Split(msg.Message, "@")
		if len(split) == 2 {
			comm := commendx{
				Commend: split[1],
				Msg:     split[0],
			}
			commFunc := getCommandxList(comm.Commend)
			if commFunc != nil { //查不到方法会返回nil
				fmt.Println(0, commFunc(&comm), 0)
			}
			msg.Message = comm.Msg
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

//socket-可以启动测试-8001
func main() {
	logx.Info(1)
	r := gin.Default()
	route(r)
	r.Run("0.0.0.0:8001")
}
*/
