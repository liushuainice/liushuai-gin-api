package model

import "time"

// BannedList 封禁名单
type BannedList struct {
	ID              uint32    `gorm:"primary_key"` // 用户ID
	PublicBannedTo  time.Time // 公共聊天封禁到的时间
	TeamBannedTo    time.Time // 战队聊天封禁到的时间
	PrivateBannedTo time.Time // 私聊封禁到的时间
}
