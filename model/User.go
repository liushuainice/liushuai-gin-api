package model

type User struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(80)"`
	Pwd  string `gorm:"type:varchar(80)"`
}
