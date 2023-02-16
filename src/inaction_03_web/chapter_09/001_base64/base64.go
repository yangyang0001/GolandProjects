package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	param := "你好, 世界! Hello World"

	encode := Base64Encode(param)
	fmt.Printf("encode = %v \n", encode)

	decode := Base64Decode(encode)
	fmt.Printf("decode = [%v] \n", decode)
}

func Base64Encode(param string) string {
	return base64.StdEncoding.EncodeToString([]byte(param))
}


func Base64Decode(param string) string {
	decode, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		log.Fatal(err)
	}
	return string(decode)
}
