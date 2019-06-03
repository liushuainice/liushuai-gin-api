package main

import (
	"github.com/gin-gonic/gin"
	"liushuai-gin-api/app"
	"liushuai-gin-api/model"
	"liushuai-gin-api/webs"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) //产生随机数种子
	app.Parse()                      // 解析配置
	webs.ThirdClient()               //连库
	model.AutoMigratex()             //创建model相关表
}

//8083
func main() {
	webs.Services()            //启动ws的管道推送接口
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，debug会多很多日志，线上环境为gin.ReleaseMode
	r := gin.New()             //启动站点
	webs.MiddleT(r)            //加载中间件
	webs.Route(r)              //加载路由
	webs.Runing(r)
}
