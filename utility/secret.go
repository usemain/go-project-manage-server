package utility

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateUniqueID 生成唯一ID
func GenerateUniqueID(length int, secret string) string {
	// 创建一个指定长度的字节数组
	bytes := make([]byte, length)
	// 将密钥转换为字节数组
	secretBytes := []byte(secret)
	// 使用加密随机源和密钥混合生成随机字节
	for i := 0; i < length; i++ {
		bytes[i] = bytes[i] ^ secretBytes[i%len(secretBytes)]
	}
	// 从加密随机源中读取随机字节
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	// 将随机字节转换为十六进制字符串
	str := hex.EncodeToString(bytes)
	return str
}
