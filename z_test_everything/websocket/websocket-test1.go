package main

/*
//原始websocket的demo
import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logx "github.com/panjiang/golog"
	"log"
	"net/http"
	"strconv"
)



var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
//=================
type Cli struct {
	Ws *websocket.Conn
	Bo bool
	Name string
}
var Clis []Cli
//=================
// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}


func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		fmt.Println("handleConnections-->msg:",msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}
var i int
func handleConnections2(c *gin.Context) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true
	//=================
	i++
	cli := Cli{
		Ws:   ws,
		Bo:   true,
		Name: strconv.Itoa(i),
	}
	Clis= append(Clis, cli)
	//=================
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		fmt.Println("handleConnections-->msg:",msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			//delete struct
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		fmt.Println("handleMessages-->msg:",msg)
		//=================
		for _, v := range Clis {
			fmt.Println("Cli:",v.Name,v.Bo,v.Ws)
			err:=v.Ws.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
			}
		}
		//=================
		for client := range clients {
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
	r.HTMLRender =rr

	go handleMessages()

	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "socket", gin.H{"msg":"hello world xxx"})
	})
	r.GET("/ws",handleConnections2)

}
//socket-可以启动测试-8001
func main() {
	logx.Info(1)
	r := gin.Default()
	route(r)
	r.Run("0.0.0.0:8001")
}*/

/*func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	// Start listening for incoming chat messages
	go handleMessages()
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)



	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}*/
