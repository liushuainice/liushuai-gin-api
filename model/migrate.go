package model

import (
	"github.com/jinzhu/gorm"
	"liushuai-gin-api/util/cli/db"
	"liushuai-gin-api/util/golog"
)

func must(sql *gorm.DB) {
	if sql.Error != nil {
		log.Fatal("migrate db:", sql.Error)
	}
}

//判断表是否为空
func tableIsEmpty(table string) bool {
	var count int64
	must(db.DB.Table(table).Count(&count))
	return count == 0
}

// AutoMigrate auto create or modify tables
func AutoMigratex() {
	log.Info("auto migrate db")
	must(db.DB.AutoMigrate(
		new(User),
		//new(Robot),
		//new(DukeCampaign),
		//new(BannedList),
	))

	// Avatar表手动生成并填充数据
	/*	if tableIsEmpty("avatars") {
		if !app.Config.Testing {
			models.InitAvatars()
			log.Info("avatars.csv created, import it by 'LOAD DATA LOCAL INFILE...' command")
			os.Exit(1)
		} else {
			models.InitAvatarsTest()
		}
	}*/
	log.Info("auto migrate db done")
}
