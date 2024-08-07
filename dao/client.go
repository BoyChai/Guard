package dao

import (
	"log"

	"github.com/BoyChai/Guard/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitClient 数据库链接
func InitClient() {
	var err error
	if viper.Get("DATABASE.DB_TYPE") == "sqlite" {
		Dao.db, err = gorm.Open(sqlite.Open("guard.db"), &gorm.Config{})
		if err != nil {
			log.Fatalln("链接数据库失败")
		}
	} else {
		log.Fatalln("数据库类型为支持")
	}

}

// AutoTables 自动加载数据表
func AutoTables() {
	var err error
	// 用户表
	err = Dao.db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln("加载数据表错误 -- User")
	}
	// 卡密
	err = Dao.db.AutoMigrate(&Card{})
	if err != nil {
		log.Fatalln("加载数据表错误 -- Card")
	}
	superUserCreate()
}

// 超级用户维护
func superUserCreate() {
	// 检查数据库中是否存在admin用户没有则创建
	var count int64
	tx := Dao.db.Model(&User{}).Where("name = ?", "admin").Count(&count)
	if tx.Error != nil {
		log.Fatalln("UpserUser检查错误")
	}
	if count == 0 {
		Dao.db.Create(&User{
			Name: viper.GetString("Settings.SuperUserName"),
			Pass: utils.CalculateMD5Hash(viper.GetString("Settings.SuperUserPass")),
		})
	}
}
