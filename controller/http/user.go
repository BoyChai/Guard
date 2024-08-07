package http

import (
	"fmt"
	"net/http"

	"github.com/BoyChai/Guard/dao"
	"github.com/BoyChai/Guard/utils"
	"github.com/gin-gonic/gin"
)

var User user

type user struct {
}

// Login 登录
func (u *user) Login(c *gin.Context) {
	params := new(struct {
		Name string `form:"name" binding:"required"`
		Pass string `form:"pass"  binding:"required"`
	})
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	if params.Name == "" || params.Pass == "" {
		c.JSON(http.StatusBadRequest, utils.HtpJson("用户名或密码不能为空", nil))
		return
	}
	userID, err := dao.Dao.CheckUser(params.Name, params.Pass)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HtpJson("用户名或密码错误", err.Error()))
		return
	}
	if token, err := utils.GenerateToken(fmt.Sprintln(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("生成token失败", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, utils.HtpJson("登录成功", token))
		return
	}

}

// CreateUser 创建用户
func (u *user) CreateUser(c *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	params := new(struct {
		Name string `form:"name" binding:"required"`
		Pass string `form:"pass"  binding:"required"`
	})
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	id, err := dao.Dao.CreateUser(params.Name, params.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("创建用户错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.HtpJson("创建用户成功", id))
}

// DeleteUser删除用户
func (u *user) DeleteUser(c *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	params := new(struct {
		ID uint `form:"id" binding:"required"`
	})
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	if err := dao.Dao.DeleteUserByID(params.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("删除用户错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.HtpJson("删除用户成功", params.ID))
}

// GetUserList 列出用户
func (u *user) GetUserList(c *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	users, err := dao.Dao.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HtpJson("获取用户列表失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.HtpJson("获取用户列表成功", users))
}
