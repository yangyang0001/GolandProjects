package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println("-------------------------- 1、练习题1      ----------------------------")
	var sl []byte
	data := []byte{1, 3, 5, 7, 9, 2, 4, 6, 8, 10, 100}
	sl = Append(sl, data)
	fmt.Println(sl)

	fmt.Println("-------------------------- 2、切片扩容方式  ----------------------------")
	var arr []int
	for i := 0; i < 20; i++ {
		arr = append(arr, i)
		fmt.Printf("len(arr) = %v, cap(arr) = %v, arr = %v \n", len(arr), cap(arr), arr)
	}

	fmt.Println("-------------------------- 3、buffer的使用方式  -----------------------")
	var buf bytes.Buffer
	buf.WriteString("zhangsan, lisi, wangwu")
	slice1, slice2 := buf.Bytes()[:4], buf.Bytes()[4:]
	fmt.Printf("slice1 =:%v, slice2 =:%v \n", string(slice1), string(slice2))

}

func Append(slice, data []byte) []byte {

	for _, b := range data {
		slice = append(slice, b)
	}

	return slice
}
