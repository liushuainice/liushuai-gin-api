package main

import (
	"flag"
	"ginWebTest/z_test_everything/websocket/example-dcm/ws"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8087", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("z_test_everything/websocket/example-dcm/view/home.html")
	type Parm struct {
		Info string
	}
	t.Execute(w, Parm{"hello ws"})
}

func main() {
	flag.Parse()
	hub := ws.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
