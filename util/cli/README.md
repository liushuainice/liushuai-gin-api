# cli

## Init a mysql client (with 'jinzhu/GORM')
```
import "github.com/panjiang/cli/db"

// Init instance
conf := db.MysqlConfig{Addr: "localhost:3306", User: "root", Password: "root", DB: "test"}
if err := db.InitMysqlCli(&conf); err != nil {
	log.Fatal(err)
}

// Use it 
if err := db.DB.Where("email=?", "test@qq.com").First(&account).Error; err != nil {
	Panic(err)
}

// ...
```
