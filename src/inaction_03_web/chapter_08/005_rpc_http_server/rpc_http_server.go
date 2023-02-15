package main

import (
	objects "deepblue.com/inaction_03/chapter_08/rpc_objects"
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {

	param := new(objects.Mine)
	rpc.Register(param)
	rpc.HandleHTTP()

	err := http.ListenAndServe("localhost:1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
