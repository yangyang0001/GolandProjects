package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("------------------------------------ 1、创建单个文件夹 ------------------------------------")
	path1 := "./maosan"
	err := os.Mkdir(path1, 0777)
	if err != nil {
		fmt.Printf("os mkdir method error, the error = %v \n", err)
	}

	fmt.Println("------------------------------------ 2、创建嵌套文件夹 ------------------------------------")
	path2 := "./gousi/test1/test2"
	path3 := "./gousi/test1/test3"
	err = os.MkdirAll(path2, 0777)
	err = os.MkdirAll(path3, 0777)
	if err != nil {
		fmt.Printf("os mkdirall method error, the error = %v \n", err)
	}

	time.Sleep(time.Second * 1)

	fmt.Println("------------------------------------ 3、删除单文件夹 ------------------------------------")
	err = os.Remove("./maosan")
	if err != nil {
		fmt.Printf("os remove method error, the error = %v \n", err)
	}

	fmt.Println("------------------------------------ 4、删除嵌套件夹 ------------------------------------")
	err = os.RemoveAll("./gousi")
	if err != nil {
		fmt.Printf("os removeall method error, the error = %v \n", err)
	}
}

