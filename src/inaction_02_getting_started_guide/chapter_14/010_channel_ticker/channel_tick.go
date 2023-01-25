package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(time.Second * 1)			//  每隔多长时间 执行!
	afte := time.After(time.Second * 20)		//  执行多长时间后, 只执行一次!

	for  {
		select {

		case <- tick:
			SayHello()

		case <-afte:
			AfterFunc()

		}
	}

}

func SayHello()  {
	fmt.Printf("hello world, time = %v \n", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}

func AfterFunc()  {
	fmt.Printf("after func invoke, time = %v", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
}
