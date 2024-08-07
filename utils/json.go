package utils

import "github.com/gin-gonic/gin"

func HtpJson(msg, data any) gin.H {
	return gin.H{"msg": msg, "data": data}
}
