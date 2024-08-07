package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/spf13/viper"
)

func CalculateMD5Hash(input string) string {
	// 创建MD5哈希对象
	h := md5.New()
	// 将字符串转换为字节数组并计算哈希值
	pass := input + viper.GetString("Settings.PassSalt")
	h.Write([]byte(pass))
	// 获取MD5哈希的字节数组
	hashBytes := h.Sum(nil)
	// 将字节数组转换为16进制字符串
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
