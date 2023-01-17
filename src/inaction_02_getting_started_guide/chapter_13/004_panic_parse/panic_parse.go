package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ParseError struct {
	Index int
	Wrod string
	Err error
}

func main() {

	//input := "1 3 5 10 22"
	input := "1 3 zhangsan lisi"

	numbers, err := Parse(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers)

}

func Parse(input string) ([]int, error)  {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing error = %v \n", e)
		}

	}()

	fields := strings.Fields(input)
	return fields2numbers(fields)

}

func fields2numbers(params []string) ([]int, error)  {

	if len(params) == 0 {
		panic("params is nil!")
	}

	var numbers []int

	for index, param := range params {
		num, err := strconv.Atoi(param)
		if err != nil {
			panic(ParseError{index, param, errors.New(param + " is not number!")})
		}
		numbers = append(numbers, num)
	}
	
	return numbers, nil
}