package util

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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
	fmt.Printf("密钥: %x\n", key)

	// 待加密明文
	plaintext := []byte("机密数据123")

	// 加密
	ciphertext, nonce, err := Aes256GCMEncrypt(key, plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密结果:\nNonce: %x\n密文: %x\n", nonce, ciphertext)
	fmt.Printf("%d%dz%x%xz%x", 1, 1+int64(nonce[0]), ^nonce[0], nonce[1:], ciphertext)
	// 解密
	decrypted, err := Aes256GCMDecrypt(key, nonce, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密结果: %s\n", decrypted)
}

func TestDecryptAccessToken(t *testing.T) {
	token := `157zc7394a1f9df5245508e3e9d1z8094a351dbcbd9ab734d8e3cd635b960893669c79ee3166f7f17840b624a97`
	tokenSlice := strings.Split(token, "z")
	//typeValue := tokenSlice[0][0:1]
	id, err := strconv.ParseInt(tokenSlice[0][1:], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	temp, err := hex.DecodeString(tokenSlice[1][0:2])
	if err != nil {
		fmt.Println(err)
		return
	}
	num := ^temp[0]
	nonce, _ := hex.DecodeString(fmt.Sprintf("%x%s", num, tokenSlice[1][2:]))
	// aes256解密
	key, _ := hex.DecodeString("6d7e77155406a25f106a25d25eed24bc2b701ca712ffe327f3c0df6302266db3")
	ciphertext, _ := hex.DecodeString(tokenSlice[2])
	plaintext, err := Aes256GCMDecrypt(key, nonce, ciphertext)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密结果: %d %s\n", id-int64(num), plaintext)
}
