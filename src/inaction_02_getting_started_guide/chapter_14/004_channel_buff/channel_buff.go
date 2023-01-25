package main

import "fmt"

func main() {

	channel := make(chan int)
	fmt.Printf("none buff channel cap = %v \n", cap(channel))

	channel = make(chan int, 100)
	fmt.Printf("have buff channel cap = %v \n", cap(channel))


}

func send()  {



}


func receive() {

}
