package main

import (
	"fmt"
	"strconv"
)

type Mine struct {
	a int
	b int
}

type TwoInts struct {
	a int
	b int
}

func main() {

	mine := Mine{1, 2}
	fmt.Printf("mine = %v \n", mine.String())
	fmt.Printf("mine = %v \n", mine)

	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)

}

func (mine *Mine) String() string {
	return "a = " + strconv.Itoa(mine.a) + ", b = " + strconv.Itoa(mine.b)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
