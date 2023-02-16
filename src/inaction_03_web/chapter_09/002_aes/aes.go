package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {

	plaintext := "Hello everyone, I am Yang!"

	keytext := "Yang199001"
	sha256 := sha256.New()
	sha256.Write([]byte(keytext))
	var buff []byte
	secretkey := sha256.Sum(buff)

	c, err := aes.NewCipher(secretkey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("c.BlockSize = %v \n", c.BlockSize())

	block := make([]byte, c.BlockSize())
	cfb := cipher.NewCFBEncrypter(c, block)

	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, []byte(plaintext))
	fmt.Printf("plaintext = %s, ciphertext = %x \n", plaintext, string(ciphertext))

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, block)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("ciphertext = %x, plaintextCopy = %s \n", ciphertext, string(plaintextCopy))

}
