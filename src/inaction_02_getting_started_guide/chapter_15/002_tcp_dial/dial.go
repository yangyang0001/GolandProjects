package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:50000") // tcp ipv4
	checkconn(conn, err)

	conn, err = net.Dial("udp", "localhost:50000") // udp
	checkconn(conn, err)

	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:50000") // tcp ipv6
	checkconn(conn, err)

}

func checkconn(conn net.Conn, err error)  {

	if err != nil {
		fmt.Printf("connect error is : %v \n", err.Error())
	}

	fmt.Printf("connection is %v \n", conn)

}