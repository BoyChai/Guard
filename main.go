package main

import (
	"fmt"

	"github.com/BoyChai/Guard/config"
	"github.com/BoyChai/Guard/controller/http"
	"github.com/BoyChai/Guard/controller/middle"
	"github.com/BoyChai/Guard/controller/socket"
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
	// 启动socket
	if viper.GetBool("Settings.Socket") {
		go socket.StartSocket()
		fmt.Println("Socket客户端已启动 Port:" + viper.GetString("Settings.Socket_Port"))
	}
	// 启动http
	g.Run(":" + viper.GetString("Settings.Port"))
}
