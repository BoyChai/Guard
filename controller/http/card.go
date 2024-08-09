package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BoyChai/Guard/dao"
	"github.com/BoyChai/Guard/utils"
	"github.com/gin-gonic/gin"
)

var Card card

type card struct {
}

// 获取所有卡密
func (c *card) GetCardList(ctx *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	cardList, err := dao.Dao.ListCard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("获取卡密列表失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.HtpJson("获取卡密列表成功", cardList))

}

// CreateCard 卡密创建
func (c *card) CreateCard(ctx *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	id, err := utils.GetAuthInfo(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	params := new(struct {
		Key  string `form:"key" `
		Time int64  `form:"time"  binding:"required"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	timeObj := time.Unix(params.Time, 0)
	// 如果key等于空则生成uuid当成key
	if params.Key == "" {
		params.Key = utils.GenerateUUID()
	}

	idNum, err := strconv.Atoi(utils.Trim(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("claims信息转换失败", err.Error()))
		return
	}
	if err := dao.Dao.CreateCard(params.Key, timeObj, uint(idNum)); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("创建卡密失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.HtpJson("创建卡密成功", nil))
}

// DeleteCardByID 删除卡密
func (c *card) DeleteCardByID(ctx *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	params := new(struct {
		ID uint `form:"id" binding:"required"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	if err := dao.Dao.DeleteCardByID(params.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("删除卡密失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.HtpJson("删除卡密成功", nil))

}

// 修改卡密有效期
func (c *card) UpdateCardEndDateByID(ctx *gin.Context) {
	// 获取 JWTAuth 中间件设置的 claims
	_, err := utils.GetAuthInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("授权信息获取失败", err.Error()))
		return
	}
	params := new(struct {
		ID   uint  `form:"id" binding:"required"`
		Time int64 `form:"time" binding:"required"`
	})

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}

	timeObj := time.Unix(params.Time, 0)

	if err := dao.Dao.UpdateCardEndDateByID(params.ID, timeObj); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("修改卡密有效期失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.HtpJson("修改卡密有效期成功", nil))
}

// CheckCard卡密的校验
func (c *card) CheckCard(ctx *gin.Context) {
	params := new(struct {
		Msg string `form:"msg" binding:"required"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("绑定参数错误", err.Error()))
		return
	}
	data, err := utils.DecryptWithPrivateKey(params.Msg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HtpJson("数据解析失败", err.Error()))
	}
	isTrue, err := dao.Dao.CheckCard(data)
	if !isTrue {
		ctx.JSON(http.StatusOK, utils.HtpJson("卡密校验失败", err.Error()))
	}
	signData, err := utils.SignData(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.HtpJson("签名失败", err.Error()))
	}
	ctx.JSON(http.StatusOK, signData)
}
