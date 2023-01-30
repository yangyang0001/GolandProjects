package main

import (
	_12_rpc_objects "deepblue.com/rpc_objects"
	"fmt"
	"log"
	"net/rpc"
)
func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	// Synchronous call
	object := &_12_rpc_objects.MineObject{7, 8}
	var result int
	err = client.Call("MineObject.Multi", object, &result)

	if err != nil {
		log.Fatal("MineObject error:", err)
	}

	fmt.Printf("MineObject: %d * %d = %d \n", object.N, object.M, result)
}
