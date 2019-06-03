package main

import (
	"fmt"
	logx "ginWebTest/util/golog"
	"github.com/googollee/go-socket.io"
	"net/http"
)

//socket-可以启动测试-8002
func main() {
	logx.Info(1)
	server, err := socketio.NewServer(nil)
	if err != nil {
		logx.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println(11, "connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./vsocket")))
	logx.Println("Serving at localhost:8002...")
	logx.Fatal(http.ListenAndServe(":8002", nil))
}
