package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
)

func main() {
	var result []byte
	hash := sha1.New()

	fmt.Println("--------------------------- 1、为 test 字符串产生 hash 值! ---------------------------")
	io.WriteString(hash, "test")
	fmt.Printf("hash result = %x \n", hash.Sum(result))
	fmt.Printf("hash result = %v \n", hash.Sum(result))


	fmt.Println("--------------------------- 2、为 zhangsan 字符串产生 hash ---------------------------")
	hash.Reset()
	data := []byte("zhangsan")
	n, err := hash.Write(data)
	if n != len(data) || err != nil {
		log.Printf("hash write error: %v / %v", n, err)
	}
	checksum := hash.Sum(result)
	fmt.Printf("hash result = %x \n", checksum)

	fmt.Println("--------------------------- 3、为 zhangsan 字符串产生 hash256 ---------------------------")
	sha256 := sha256.New()
	sha256.Write([]byte("zhangsan"))

	var buff []byte
	sum := sha256.Sum(buff)
	fmt.Printf("sum = %x \n", sum)



}
