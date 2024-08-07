package http

import "github.com/gin-gonic/gin"

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体w
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	// 用户相关api
	router.GET("/api/user/getList", User.GetUserList).
		POST("/api/user/login", User.Login).
		POST("/api/user/create", User.CreateUser).
		DELETE("/api/user/delete", User.DeleteUser).
		// 卡密相关api
		GET("/api/card/getList", Card.GetCardList).
		POST("/api/card/create", Card.CreateCard).
		DELETE("/api/card/delete", Card.DeleteCardByID).
		PUT("/api/card/update", Card.UpdateCardEndDateByID)
}
