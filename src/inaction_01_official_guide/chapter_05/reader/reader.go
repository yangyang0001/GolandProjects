package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	fmt.Println(len("Hello, World!"))
	reader := strings.NewReader("Hello, World!")

	bytes := make([]byte, 8)

	for {
		/**
		 * 返回值 n 表示: 每次实际读取的字节数; Read 方法是每次读取 8个 字节, 这里要 make 一个 8个字节的切片!
		**/
		n, err := reader.Read(bytes)

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("n = %v, bytes = %v \n", n, string(bytes[:n]))
		}

	}
}
