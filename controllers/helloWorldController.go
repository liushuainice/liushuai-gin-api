package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"liushuai-gin-api/app"
	"liushuai-gin-api/model"
	"liushuai-gin-api/util"
	"liushuai-gin-api/util/cli/db"
	logx "liushuai-gin-api/util/golog"
	"time"
)

func HelloIndexGet(c *gin.Context) {
	//m := map[string]interface{}{"message": "pong"}
	//c.JSON(200, m)
	//c.JSON(200, gin.H{"message": err, "ok": ok, "er": er})
	//err := db.RDS.Set("ls", "贼特么帅！！！", 0).Err()
	//ok, er := db.RDS.Set("ls+1", "贼特么帅！！！+1", 0).Result()
	util.Success(c, gin.H{"message": "hello-world", "ok": 666, "er": "nothing"})

}
func HelloIndexPost(c *gin.Context) {
	type HeroIDJSON struct {
		HeroID   uint32 `json:"heroid" binding:"required"`
		NickName string `json:"nickname" binding:"required"`
	}
	var params HeroIDJSON
	if err := c.ShouldBindJSON(&params); err != nil {
		panic(err)
	}
	logx.Info(1, params, 2)
	logx.Error(1, params, 2)

	util.Success(c, gin.H{"message": "ok"})

}
func AnyTest(c *gin.Context) {
	type test struct {
		HeroID   uint32 `json:"heroid" binding:"required"`
		NickName string `json:"nickname" binding:"required"`
	}
	var params test
	if err := c.ShouldBindJSON(&params); err != nil {
		panic(err)
	}
	fmt.Println(44, params, 44)
	util.Success(c, gin.H{"message": "ok"})
}

type Userx struct {
	Id   int32
	Name string
}

func Sum(x, y int) int {
	return x + y
}
func Template(c *gin.Context) {
	data1 := Userx{1, "a1"}
	data2 := Userx{2, "a2"}
	data3 := Userx{3, "a3"}
	var dataList []Userx
	dataList = append(dataList, data2, data3)
	//"SumTest": Sum--》函数传递到前端，前端可以使用计算
	c.HTML(200, "user/index.html", gin.H{"title": "China", "data": data1, "dataList": dataList, "SumTest": Sum})
}

type sysChatbanReq struct {
	Password string `json:"password" binding:"required"`
	UID      uint32 `json:"uid" binding:"required,min=1"` // 要禁言的玩家uid
	Type     string `json:"type" binding:"required"`      // 禁言类型 公共/战队/私聊/all
	Time     uint64 `json:"time" binding:"required"`      // 禁言多少秒
}

// SysChatban 封禁用户
func SysChatban(c *gin.Context) {
	var req sysChatbanReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}
	uid := req.UID
	bannedTo := time.Now().Add(time.Duration(req.Time) * time.Second)
	var ban model.BannedList
	ban.ID = uid
	switch req.Type {
	case "public":
		ban.PublicBannedTo = bannedTo

	default:
		util.Success(c, gin.H{"msg": "err"})
		return
	}
	//mysql 添加或修改
	if err := db.DB.Assign(ban).FirstOrCreate(&ban).Error; err != nil {
		panic(err)
	}

	// 返回数据
	util.Success(c, gin.H{"msg": "ok"})
}
func GetAllUsersOnlineC(c *gin.Context) {
	allOnlinex := app.ServiceOnline.GetAllUsersOnlinex()
	for k, v := range allOnlinex {
		fmt.Println(k, v)
	}
	util.Success(c, gin.H{"msg": "ok"})
}
