package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("I am before next 1")
		c.Header("Access-Control-Allow-Origin", "test")
		c.Set("name", "test")
		/*
		   c.Next()后就执行真实的路由函数(i am here 2)，路由函数执行完成之后继续执行后续的代码(i am after next 3)
		*/
		c.Next()
		fmt.Println("I am after next 3")
	}
}

//中间件demo
func main() {
	r := gin.Default()
	r.Use(middleware()) //自己写中间件
	r.GET("/test", func(c *gin.Context) {
		fmt.Println("I am here 2")
		name, _ := c.Get("name")
		c.JSON(200, gin.H{"name": name})
	})
	r.Run("0.0.0.0:10022")
}
