package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main method invoke start ...")

	longwait()

	shortwait()

	fmt.Println("sleep in main method, time is 10s!")

	time.Sleep(time.Second * 1)

	fmt.Println("main method invoke end!")
}

func longwait()  {

	fmt.Println("longwait method invoke start ...")

	time.Sleep(time.Second * 5)

	fmt.Println("longwait method invoke end!")

}

func shortwait()  {
	fmt.Println("shortwait method invoke start ...")

	time.Sleep(time.Second * 2)

	fmt.Println("shortwait method invoke end!")
}
