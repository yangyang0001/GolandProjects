package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("------------------------------------ 1、os_args 使用 ------------------------------------")
	args:= os.Args

	for index, arg := range args {
		fmt.Printf("os.args index = %v, arg = %v \n", index, arg)
	}

	fmt.Printf("hello %v \n", args[1])

	// 使用命令: go run . zhangsan 则会出现 Hello zhangsan 的结果!

	fmt.Println("------------------------------------ 2、flag 包使用 ------------------------------------")
	f := flag.Bool("flag", true, "mine flag")
	num_args := flag.NArg()
	fmt.Printf("*f = %v, num_args = %v \n", *f, num_args)



}
