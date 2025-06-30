// Package util 加解密封装
package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"io"
)

// 单向

// BcryptHash bcrypt算法
func BcryptHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // 推荐使用默认成本值12
	)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CompareBcryptHash bcrypt算法对比
func CompareBcryptHash(hashPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
}

// SHA3256Hash sha3-256算法
func SHA3256Hash(data string) string {
	// 创建SHA3-256哈希对象
	hasher := sha3.New256()

	// 写入数据
	hasher.Write([]byte(data))

	// 计算哈希值
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash)
}

// SM3Hash sm3算法
func SM3Hash(data string) string {
	hash := sm3.New()
	hash.Write([]byte(data))
	result := hash.Sum(nil)

	return hex.EncodeToString(result)
}

// 对称加密

// Aes256GCMEncrypt aes256-gcm加密
func Aes256GCMEncrypt(key, plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// 生成12字节随机nonce（GCM标准长度）
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	// 创建GCM模式加密器
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// 加密并附加认证标签
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// Aes256GCMDecrypt aes256-gcm解密
func Aes256GCMDecrypt(key, nonce, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 解密并验证认证标签
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("解密失败: 认证标签无效")
	}
	return plaintext, nil
}

// AccessToken 整合加密生成访问令牌
func AccessToken(text string) {
	// 配置读取密钥
	// aes256加密
	// 拼接打乱替换
}

// 非对称加密
