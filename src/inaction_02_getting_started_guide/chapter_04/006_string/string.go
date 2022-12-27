package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	fmt.Println("----------------------------- 1、字符串拼接 +  --------------------------------")
	string0 := "hello" + ", world!"
	fmt.Printf("string0 = %v \n", string0)

	fmt.Println("----------------------------- 2、字符串拼接 strings.Join  --------------------")
	strarr := []string{"zhangsan", "lisi", "wangwau"}
	joinre := strings.Join(strarr, "-")
	fmt.Printf("joinre = %v \n", joinre)

	fmt.Println("----------------------------- 3、字符串拼接 bytes.Buffer 后续补充, 这里先预留 ---")

	fmt.Println("----------------------------- 4、utf8 字符串统计函数 --------------------------")
	string1 := "asSASA ddd dsjkdsjs dk"
	//str2 := "0123456789012345678901"
	fmt.Printf("string1 length = %d \n", len(string1))
	fmt.Printf("string1 characters count = %d \n", utf8.RuneCountInString(string1))

	string2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("string2 length = %d \n", len(string2))
	fmt.Printf("string2 characters count = %d \n", utf8.RuneCountInString(string2))

	string3 := "こん"
	fmt.Printf("string3 length = %d \n", len(string3))
	fmt.Printf("string3 characters count = %d \n", utf8.RuneCountInString(string3))

}
