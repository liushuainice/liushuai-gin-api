package main

import (
	"github.com/gin-gonic/gin"
)

type Userx struct {
	Id   int32
	Name string
}

func Sum(x, y int) int {
	return x + y
}

func main() {

	r := gin.Default()
	r.Static("/s", "static_test") //http://localhost:10020/s/test.txt
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	/*
		模板
	*/
	r.LoadHTMLGlob("asset/**/*")
	r.GET("/html", func(c *gin.Context) {
		data1 := Userx{1, "a1"}
		data2 := Userx{2, "a2"}
		data3 := Userx{3, "a3"}
		var dataList []Userx
		dataList = append(dataList, data2, data3)
		c.HTML(200, "user/index.html", gin.H{"title": "China", "data": data1, "dataList": dataList, "SumTest": Sum(1, 2)})
	})
	r.Run("0.0.0.0:10020")
}
