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
	// ID用户的唯一标识
	ID uint `gorm:"primaryKey;auto_increment"`
	// 用户名
	Name string `gorm:"not null"`
	// 密码(md5(Pass))
	Pass string `gorm:"not null"`
}

// Card 卡密表
type Card struct {
	gorm.Model
	// 卡密
	Key string
	// 有效租约
	Lease time.Time
	// 创建用户
	UserID uint
}
