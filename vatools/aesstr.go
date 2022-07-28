package vatools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var aesKey = []byte("JimmyFzm19800217")
var commonIV = []byte("JimmyZsp19840227")

func AESEnStr(val string) string {
	plaintext := []byte(val)

	// 创建加密算法
	c, err := aes.NewCipher(aesKey)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return ""
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func AESUnStr(val string) string {
	c, err := aes.NewCipher(aesKey)
	if err != nil {
		fmt.Println("Error AesUnStr:", err.Error())
		return ""
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	ciphertext, _ := base64.StdEncoding.DecodeString(val)
	unchiptext := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(unchiptext, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, unchiptext)
	return string(unchiptext)
}

func SetPrivateAESKey(key string) error {
	btVal := []byte(key)
	if len(btVal) != 16 {
		return fmt.Errorf("密钥长度需要16位")
	}
	aesKey = btVal
	return nil
}
