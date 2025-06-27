package util

import "crypto/rand"

// RandomByte
// @Description: 随机生成指定长度字节
// @param length
func RandomByte(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key) // 使用密码学安全随机源
	return key, err
}
