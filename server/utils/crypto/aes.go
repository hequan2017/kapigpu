package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AesEncrypt AES加密（CBC模式）
// key 必须是16、24或32字节长度
func AesEncrypt(plaintext, key string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 填充数据
	plaintextBytes := pkcs7Padding([]byte(plaintext), block.BlockSize())

	// 初始化向量
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintextBytes)

	// Base64编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesDecrypt AES解密（CBC模式）
// key 必须是16、24或32字节长度
func AesDecrypt(ciphertext, key string) (string, error) {
	if ciphertext == "" {
		return "", nil
	}

	// Base64解码
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 分离IV和密文
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	// 去除填充
	plaintext, err := pkcs7Unpadding(ciphertextBytes)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// pkcs7Padding PKCS7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7Unpadding PKCS7去除填充
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is empty")
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

// GenerateKey 生成指定长度的随机密钥
func GenerateKey(length int) (string, error) {
	if length != 16 && length != 24 && length != 32 {
		return "", errors.New("key length must be 16, 24 or 32")
	}
	key := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
