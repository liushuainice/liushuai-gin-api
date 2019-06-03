package webs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"liushuai-gin-api/app"
	"liushuai-gin-api/rservice/ws/impl"
	"liushuai-gin-api/util/cli/db"
	"liushuai-gin-api/util/golog"
	"liushuai-gin-api/webs/midware"
)

// ThirdCli 第三方客户端
func ThirdClient() {
	log.Info("init clients")
	// Redis操作句柄
	if err := db.InitRedisCli(app.Config.Redis); err != nil {
		log.Fatal(err)
	}
	// Mysql操作句柄
	if err := db.InitMysqlCli(app.Config.Mysql); err != nil {
		log.Fatal(err)
	}

	// 开启DB日志模式
	db.DB.LogMode(app.Config.Debug)
}
func MiddleT(r *gin.Engine) {
	r.Use(gin.Logger()) //控制台打印信息

	// 创建记录日志的文件
	//r.Use(gin.Recovery())         //gin 恢复恐慌的中间件--》带测试
	r.Use(midware.RecoverPanic()) // 捕获Panic异常
	//r.Use(midware.CORS())         //http 跨域访问什么的
}

// Run 监听端口运行服务
func Runing(r *gin.Engine) {
	addr := fmt.Sprintf(":%d", app.Config.BindPort)
	mode := "development"
	if !app.Config.Debug {
		mode = "production"
	}
	log.Infof("listening on %s (%s)", addr, mode)
	r.Run(addr)
}

// Services 服务
func Services() {
	log.Info("run websocket online service")
	app.ServiceOnline = impl.NewOnline()
}
