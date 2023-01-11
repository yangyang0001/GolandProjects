package main

import "fmt"

type List []int

type Appender interface {
	Append(param int)
}

type Lener interface {
	Len() int
}

func main() {

	var list List
	CountInto(&list, 1, 10)

	for _, val := range list {
		fmt.Printf("val = %v \n", val)
	}
}

func (list *List) Len() int {
	return len(*list)
}

func (list *List) Append(param int) {
	*list = append(*list, param)
}

func CountInto(appender Appender, start int, end int) {
	for i := start; i < end; i++ {
		appender.Append(i)
	}
}
