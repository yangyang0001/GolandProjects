package main

import (
	"fmt"
	"time"
)

func main() {

	go say("Hello")
	go say("World")

	time.Sleep(time.Second * 2)

}

func say(s string)  {
	fmt.Println(s)
}
