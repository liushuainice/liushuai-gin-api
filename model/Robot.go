package model

type Robot struct {
	Id     uint32  `gorm:"primary_key"`
	Ip     string  `gorm:"type:varchar(80)"`
	Power  float64 `gorm:"type:float"`
	Status int     `gorm:"type:int"`
}
