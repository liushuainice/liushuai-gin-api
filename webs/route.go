package webs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"liushuai-gin-api/controllers"
	"liushuai-gin-api/rservice/ws/impl"
	logx "liushuai-gin-api/util/golog"
)

// Route 路径映射
func Route(r *gin.Engine) {
	logx.Info(1)
	fmt.Println(1)
	//=============== websocket--->线在main里-->webs.Services()
	r.GET("/ws", impl.ServeWsV)
	r.GET("/so", impl.ServeHomeV)
	r.GET("/all", controllers.GetAllUsersOnlineC)
	//===============

	r.GET("/", controllers.HelloIndexGet)
	r.POST("/", controllers.HelloIndexPost)
	r.Any("/msg", controllers.AnyTest)
	r.POST("/sys-chatban", controllers.SysChatban)
	/*	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", "") //暂时没用html文件
		// or
		t, _ := template.ParseFiles("path/404.html") //暂时没用html文件
		t.Execute(c.Writer, nil)
	})*/
}
