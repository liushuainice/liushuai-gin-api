package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

// DB global mysql operation instance
var DB *gorm.DB

// MysqlConfig 用于解析mysql配置
type MysqlConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func newMysqlClient(addr, user, pwd, dbname string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&readTimeout=10s&timeout=10s", user, pwd, addr, dbname)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(200)
	db.DB().SetMaxIdleConns(200)
	return db, nil
}

// InitMysqlCli create a handler
func InitMysqlCli(conf *MysqlConfig) error {
	client, err := newMysqlClient(conf.Addr, conf.User, conf.Password, conf.DB)
	if err != nil {
		return err
	}

	DB = client
	return nil
}
