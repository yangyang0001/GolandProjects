package main

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	fmt.Printf("random is %v \n", Random())
}

func Random() int {

	return int(C.random())

}

func Seed(i int) {
	C.srandom(C.uint(i))
}
