package main

import (
	"github.com/BoyChai/Guard/config"
	"github.com/BoyChai/Guard/controller/http"
	"github.com/BoyChai/Guard/controller/middle"
	"github.com/BoyChai/Guard/dao"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置
	config.InitConfig()
	// 初始化数据库链接
	dao.InitClient()
	// 自动加载数据表
	dao.AutoTables()
	// 初始化路由
	g := gin.Default()
	g.Use(middle.CORS())
	g.Use(middle.JWTAuth())
	http.Router.InitApiRouter(g)
	// 启动
	g.Run(":" + viper.GetString("Settings.Port"))
}
