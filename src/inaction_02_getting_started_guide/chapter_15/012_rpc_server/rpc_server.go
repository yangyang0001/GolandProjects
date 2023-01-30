package main

import (
	_12_rpc_objects "deepblue.com/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {

	object := new(_12_rpc_objects.MineObject)
	rpc.Register(object)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)

}


