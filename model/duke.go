package model

import (
	"github.com/shopspring/decimal"
	"time"
)

//DukeCampaign 公爵竞选表-
type DukeCampaign struct {
	ID         uint32          `gorm:"primary_key"`
	UID        uint32          `gorm:"column:uid"`                  //玩家ID
	OpenID     int64           `gorm:"not null"`                    //玩家的openid
	Ticket     decimal.Decimal `gorm:"type:decimal(20,4);not null"` //票数
	WeightToID uint32          `gorm:""`                            //话语权投票给了谁
	Period     int64           `gorm:"not null"`                    //竞选公爵第N期
	Status     int             `gorm:"not null"`                    //竞选状态
	UpdateAt   time.Time       //竞选公爵时，票数最后更新时间
}
