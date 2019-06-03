package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

//templates
func main() {
	router := gin.Default()
	//<div><img src="/s/profile-pic.jpg" alt=""></div>//http://localhost:10020/s/test.txt 可以访问到静态资源
	router.Static("/s", "static_test")
	router.GET("/html", serveHome)
	router.Run(":8080")
}
func serveHome(c *gin.Context) {
	t, _ := template.ParseFiles("z_test_everything/websocket/example-dcm/view/home.html")
	type Parm struct {
		Info string
	}
	t.Execute(c.Writer, Parm{"hello ws"})
}
