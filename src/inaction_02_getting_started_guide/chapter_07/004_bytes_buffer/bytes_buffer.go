package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	var buffer bytes.Buffer
	var str = []string{"zhangsan", "lisi", "wangwu", "zhaoliu"}

	for _, val := range str {
		buffer.WriteString(val)
		buffer.WriteString("-")
	}

	result := buffer.String()
	result = strings.TrimRight(result, "-")

	fmt.Printf("result = %v \n", result)

}
