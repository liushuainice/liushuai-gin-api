package webs

import "github.com/gin-contrib/multitemplate"

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	//r.AddFromFiles("index2", "asset/socket/test/index2.html")
	r.AddFromFiles("index3", "asset/index3.html")
	r.AddFromFiles("user/index.html", "asset/user/index.html")
	r.AddFromFiles("socket", "asset/socket/index.html", "asset/socket/jquery.js", "asset/socket/socket.io.js")
	return r
}
