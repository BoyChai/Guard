package config

import (
	"log"
	"os"

	"github.com/BoyChai/Guard/utils"
	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd() //获取工作目录

	viper.SetConfigName("config")            // 设置config名字
	viper.SetConfigType("yml")               //设置配置文件类型
	viper.AddConfigPath(workDir + "/config") // 设置配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	// 解析私钥
	utils.PrivateKey, err = utils.ParsePrivateKey(viper.GetString("Settings.Private_Key_PATH"))
	if err != nil {
		log.Fatalln("通信私钥读取错误", err.Error())
	}
}
