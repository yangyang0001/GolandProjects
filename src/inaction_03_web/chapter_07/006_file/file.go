package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("------------------------------------ 1、文件的创建 ------------------------------------")
	_, err := os.Create("./football_player.txt")
	if err != nil {
		fmt.Printf("create method invoke error, the error = %v \n",  err)
	}

	fmt.Println("------------------------------------ 2、文件的删除 ------------------------------------")
	filepath := "./football_player.txt"
	err = os.Remove(filepath)
	if err != nil {
		fmt.Printf("remove method invoke error, the error = %v \n",  err)
	}

	// 文件的 读写操作这里就不提了, 参考 file_input_output.go 或 input_output_review.go

}
