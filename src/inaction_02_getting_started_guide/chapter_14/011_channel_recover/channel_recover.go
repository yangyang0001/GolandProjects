package main

import "fmt"

func main() {

	defer func() {
		if e:= recover(); e != nil {
			fmt.Printf("panicing error = %v \n", e)
		}
	}()

	SayHello()

}

func SayHello()  {
	fmt.Println("hello world")
	panic("this is sayhello error!")
}


