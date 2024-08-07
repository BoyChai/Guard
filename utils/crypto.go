package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/google/uuid"
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

// UUID生成
func GenerateUUID() string {
	return uuid.New().String()
}

// 私钥
var PrivateKey *rsa.PrivateKey

// parsePrivateKey 解析 PEM 文件中的 RSA 私钥
func ParsePrivateKey(privateKeyFile string) (*rsa.PrivateKey, error) {
	// 读取 PEM 文件内容
	pemData, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read PEM file: %v", err)
	}

	// 解码 PEM 数据块
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// 解析私钥
	var parsedKey interface{}
	if block.Type == "RSA PRIVATE KEY" {
		parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS1 private key: %v", err)
		}
	} else if block.Type == "PRIVATE KEY" {
		parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS8 private key: %v", err)
		}
	} else {
		return nil, fmt.Errorf("unknown PEM block type")
	}

	// 将解析后的私钥转换为 *rsa.PrivateKey 类型
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not an RSA private key")
	}

	return privateKey, nil
}

// 通过私钥解密
func DecryptWithPrivateKey(encryptedData string) (decryptedData string, err error) {

	// 解码 base64 编码的加密数据
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 encoded encrypted data: %v", err)
	}

	// 使用私钥解密数据
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, PrivateKey, encryptedBytes, nil)
	if err != nil {
		return "", fmt.Errorf("decryption error: %v", err)
	}

	return string(decryptedBytes), nil
}

// 通过私钥加密
func SignData(data string) (signature string, err error) {
	// 计算数据的 SHA-256 散列值
	hashed := sha256.Sum256([]byte(data))

	// 使用私钥对数据进行签名
	signatureBytes, err := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("signing error: %v", err)
	}

	// 将签名结果编码为 base64 字符串
	signature = base64.StdEncoding.EncodeToString(signatureBytes)

	return signature, nil
}
