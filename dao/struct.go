package dao

import (
	"time"

	"gorm.io/gorm"
)

var Dao dao

type dao struct {
	db *gorm.DB
}

// User 用户表
type User struct {
	gorm.Model
	// 用户名
	Name string `gorm:"not null;unique"`
	// 密码(md5(Pass))
	Pass string `gorm:"not null"`
}

// Card 卡密表
type Card struct {
	gorm.Model
	// 卡密
	Key string `gorm:"unique"`
	// 结束时间
	EndDate time.Time
	// 创建用户
	UserID uint
}
