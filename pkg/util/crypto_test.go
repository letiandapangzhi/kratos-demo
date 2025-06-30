package util

import (
	"fmt"
	"testing"
)

func TestBcryptHash(t *testing.T) {
	// 密码加密
	password := "mySecurePassword123"

	hashPassword, err := BcryptHash(password)

	fmt.Printf("Hashed password: %s Err:%v\n", hashPassword, err)
}

func TestCompareBcryptHash(t *testing.T) {
	// 密码加密
	password := "mySecurePassword123"

	hashPassword, _ := BcryptHash(password)

	// 密码验证
	err := CompareBcryptHash(hashPassword, password)
	if err != nil {
		fmt.Println("Password mismatch", err) // 验证失败
	} else {
		fmt.Println("Password correct")
	}

	err = CompareBcryptHash(hashPassword, "123")
	if err != nil {
		fmt.Println("Password mismatch", err) // 验证失败
	} else {
		fmt.Println("Password correct")
	}
}

func TestSHA3256Hash(t *testing.T) {
	// 在线sha3 256
	// https://emn178.github.io/online-tools/sha3_256.html

	// 待哈希数据
	data := "Hello, SHA3-256"

	// 输出十六进制格式
	fmt.Printf("原始数据: %s\n", data)
	fmt.Printf("SHA3-256哈希: %s\n", SHA3256Hash(data))
}

func TestSM3Hash(t *testing.T) {
	// 示例1：分步计算哈希
	data := "Hello SM3"

	fmt.Println(SM3Hash(data))
}

func TestAes256GCM(t *testing.T) {
	// 生成32字节随机密钥（256位）
	key, err := RandomByte(32)

	// 待加密明文
	plaintext := []byte("机密数据123")

	// 加密
	ciphertext, nonce, err := Aes256GCMEncrypt(key, plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密结果:\nNonce: %x\n密文: %x\n", nonce, ciphertext)

	// 解密
	decrypted, err := Aes256GCMDecrypt(key, nonce, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密结果: %s\n", decrypted)
}
