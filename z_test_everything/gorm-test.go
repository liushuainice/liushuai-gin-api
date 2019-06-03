package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"liushuai-gin-api/util/cli/db"
	"liushuai-gin-api/util/golog"
)

func init() {
	config := db.MysqlConfig{
		Addr:     "127.0.0.1:3306",
		User:     "root",
		Password: "root",
		//DB:       "youzi",
		DB: "test",
	}
	//建立连接
	if err := db.InitMysqlCli(&config); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(40)"`
	Pwd  string `gorm:"type:varchar(40)"`
}

//创建表
func AutoMigrate() {
	log.Info("auto migrate db")
	must(db.DB.AutoMigrate(
		new(User),
	))
}

//判断异常
func must(sql *gorm.DB) {
	if sql.Error != nil {
		log.Fatal("migrate db:", sql.Error)
	}
}

func main() {
	//执行原生sql,参数用站位符
	/*	type result struct {
			Id int
			Name string
			Pwd string
		}
		var r []result
		e := db.DB.Raw(`select * from users `).Scan(&r).Error
		fmt.Println(e)
		fmt.Println(r)
	*/
	//两表连查，注意 '_'和字段字母大写相关
	/*		type Test struct {
				//UID      uint32
				//Nickname string
				//Ticket   decimal.Decimal
				//ID uint32
			}
			var t []Test
			var i int
			//scan 和 count不能连用，无效
			_= db.DB.Raw(`SELECT d.id , d.uid , d.ticket , u.nickname FROM duke_campaigns d , users u WHERE u.uid=d.uid AND d.status=4 ORDER BY d.id asc LIMIT ? OFFSET ?`,4,2).Scan(&t).Count(&i).Error
			for _, v := range t {
				fmt.Println(v)
			}
			fmt.Println(i)//无效
	*/

	//AutoMigrate() //创建表--》ok

	//测试crud
	//user := User{
	//	Name: "testname",
	//	Pwd:  "testpwd",
	//}
	//err:= db.DB.Create(&user).Error
	//fmt.Println(err)

	//查列表
	/*	var users []User
		//db.DB.Order("id desc").Find(&users)
		db.DB.Not("id",[]string{"41","42"}).Order(gorm.Expr("rand()")).Find(&users)
		fmt.Println(users)
	*/
	//查一条
	//user := User{Id:2}
	//user := User{}
	//db.DB.First(&user,"id=?",1)
	//fmt.Println(user)1
	//删除
	//	db.DB.Delete(&User{}, "id=?", 1)

	/*var user User
	for i := 1; i < 10; i++ {
		j := strconv.Itoa(i)
		user = User{
			Name: "test_name_" + j,
			Pwd:  "test_pwd_" + j,
		}
		err := db.DB.Create(&user).Error
		fmt.Println(err)
	}*/

	//修改
	//db.DB.Model(&User{Id:5}).Update("name","test666")

	//外键关联重看
	//http://gorm.book.jasperxu.com/associations.html
	//查最后一条
	/*var user User
	db.DB.Last(&user)
	fmt.Println(user)*/
	//指定 字段
	//count 记录条数
	/*var i int
	db.DB.Select("name").Find(&users).Count(&i)
	fmt.Println(i)*/

	//Next()方法没引出来
	//rows, err := db.Table("orders").
	// Select("date(created_at) as date, sum(amount) as total").
	// Group("date(created_at)").Rows()
	//for rows.Next() {
	//}
	/*user:=User{
		Id:   9,
		Name: "adsf",
		Pwd:  "",
	}
	db.DB.Save(&user)//所有字段全部更改*/
	//db.DB.Model(&User{}).UpdateColumns(User{Pwd:"pwd666"})

	//model 只会byid修改，没用id的化，数据库会全部数据修改，加条件的化，先用where
	/*	var u User
		//u.Id=45//-->ok
		//u.Name="66666"//--err
		m:=map[string]interface{}{}
		m["name"]="555555"
		//db.DB.Model(&u).Updates(m)
		db.DB.Where("id=?",44).Model(&u).Updates(m)//--ok*/

	//条件删除
	//db.DB.Delete(User{},"pwd=? and name=?","test_pwd_6","test_name_6")//-->ok
	//db.DB.Where("pwd=?","test_pwd_7").Delete(User{})//-->ok,一定要where在前
	//if err := db.DB.Model(&User{}).
	//	Where("id = ?", 44).

	/*i := 44
	if err := db.DB.Model(&User{Id: i}).
		UpdateColumns(map[string]interface{}{"name": "66666"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(666)
	}*/

}
