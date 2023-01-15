package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("---------------------------------------- 1、使用 Scanln, Sscanf ----------------------------------------")
	var firtname string
	var lastname string

	fmt.Scanln(&firtname, &lastname)
	fmt.Printf("firtname = %v, lastname = %v \n", firtname, lastname)

	var i int
	var j float32
	var k string

	fmt.Scanln(&i, &j, &k)
	fmt.Printf("i = %v, j = %v, k = %v \n", i, j, k)

	var m int
	var n float64
	var q string

	var input string = "10 , 100 , zhangsan"
	var format string = "%d , %f , %s"

	fmt.Sscanf(input, format, &m, &n, &q)
	fmt.Printf("m = %v, n = %v, q = %v \n", m, n, q)

	fmt.Println("------------------------------------------ 2、使用 reader ------------------------------------------")
	inputReader := bufio.NewReader(os.Stdin)

	readString, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal("error is :", err)
	} else {
		fmt.Printf("readString = %v \n", readString)
	}

	fmt.Println("------------------------------------------ 3、input_switch -----------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	input0, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("input0 = %v \n", input0)
	}

	switch input0 {
	case "zhangsan\n":
		fmt.Println("input string is zhangsan!")
	case "lisi\n":
		fmt.Println("input string is lisi!")
	default:
		fmt.Println("input string is other!")
	}

}
