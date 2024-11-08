package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 转小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5Encode 转大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(plainPwd, salt string) string {
	return MD5Encode(Md5Encode(plainPwd + salt))
}

// ValidPassword 解密
func ValidPassword(plainPwd, salt, password string) bool {
	return MD5Encode(Md5Encode(plainPwd+salt)) == password
}
