package main

import (
	"fmt"
	logx "ginWebTest/util/golog"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
)

var SocX *socketio.Server

func route(r *gin.Engine) {
	//r.Static("/asset", "./asset")
	r.Static("/vsocket", "./vsocket")
	rr := multitemplate.NewRenderer()
	rr.AddFromFiles("socket", "vsocket/index.html")
	r.HTMLRender = rr
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "socket", gin.H{"msg": "hello world xxx"})
	})
	r.GET("/socket.io/", func(c *gin.Context) {
		SocX.ServeHTTP(c.Writer, c.Request)
	})

}

//现在还是失败，可以启动，但是页面的js和后台连接不上
//socket-可以启动测试-8004
func main() {

	logx.Info(1)
	//r := gin.Default()
	r := gin.New() // 启动站点
	socket()
	route(r)
	err := r.Run("0.0.0.0:8004")
	if err != nil {
		logx.Error(err)
	}
	SocX.Close()
}

//var SocX *socketio.Server

func socket() {
	var err error
	SocX, err = socketio.NewServer(nil)
	if err != nil {
		logx.Fatal(err)
	}
	SocX.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("0")
		fmt.Println(1, "connected:", s.ID())
		return nil
	})
	SocX.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println(2, "notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	SocX.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	SocX.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	SocX.OnError("/", func(e error) {
		fmt.Println(3, "meet error:", e)
	})
	SocX.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println(4, "closed", msg)
	})
	go SocX.Serve()
}
