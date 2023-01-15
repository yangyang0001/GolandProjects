package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	/**
	练习 12.1: word_letter_count.go

		编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
		i) 输入的字符的个数，包括空格，但不包括 '\n'
		ii) 输入的单词的个数
		iii) 输入的行数

	练习 12.2: calculator.go

		编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
		输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
		当用户输入字符 'q' 时，程序结束。请使用您在练习 11.13 中开发的 stack 包。
	*/

	fmt.Println("----------------------------------- 1、练习1 -------------------------------------")
	reader := bufio.NewReader(os.Stdout)
	input, err := reader.ReadString('S')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("input = %v \n", input)
	}

	var words int = len(strings.Fields(input))
	var lines int = strings.Count(input, "\n")
	if lines == 0 {
		lines = 1
	}

	input = strings.ReplaceAll(input, "\n", "")
	var count int = len(input) - 1

	fmt.Printf("words = %v, lines = %v, count = %v \n", words, lines, count)

	fmt.Println("----------------------------------- 2、练习2 -------------------------------------")
	bufferReader := bufio.NewReader(os.Stdin)
	inputstring, err := bufferReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	//inputstring = "10 ENTER 2 ENTER + ENTER"
	//inputstring = "10 ENTER 2 ENTER - ENTER"
	//inputstring = "10 ENTER 2 ENTER * ENTER"
	//inputstring = "10 ENTER 2 ENTER / ENTER"
	inputstring = strings.ReplaceAll(inputstring, " ", "")
	fmt.Println(inputstring)
	split := strings.Split(inputstring, "ENTER")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])

	switch split[2] {
	case "+":
		fmt.Printf("%v + %v = %v \n", a, b, a+b)
	case "-":
		fmt.Printf("%v - %v = %v \n", a, b, a-b)
	case "*":
		fmt.Printf("%v * %v = %v \n", a, b, a*b)
	case "/":
		fmt.Printf("%v / %v = %v \n", a, b, a/b)

	default:
		fmt.Println("this is error!")
	}

}
