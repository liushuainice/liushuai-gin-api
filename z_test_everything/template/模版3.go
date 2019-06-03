package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	//r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	//r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	r.AddFromFiles("index2", "asset/socket/test/index2.html")
	r.AddFromFiles("index3", "asset/index3.html")
	r.AddFromFiles("index4", "asset/index3.html", "asset/socket/test/index2.html") //第二个文件路径不知道有啥用
	return r
}

//templates
func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index2", gin.H{
			"msg": "Html5 Template Engine",
		})
	})
	router.GET("/a", func(c *gin.Context) {
		c.HTML(200, "index3", gin.H{
			"msg": "Html5 Article Engine",
		})
	})
	router.GET("/b", func(c *gin.Context) {
		c.HTML(200, "index4", gin.H{
			"msg": "Html5 Hello World !!!",
		})
	})
	router.Run(":8080")
}
