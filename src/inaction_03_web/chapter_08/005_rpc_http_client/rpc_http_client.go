package main

import (
	objects "deepblue.com/inaction_03/chapter_08/rpc_objects"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	http, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	param := &objects.Mine{"zhangsan", "123456"}
	var result string

	err = http.Call("Mine.SayHello", param, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result = %v \n", result)
}
