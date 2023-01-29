package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("the tcp server starting ...")
	listen, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Printf("error listen, the error is : %v \n", err.Error())
		return
	}

	for  {
		connect, err := listen.Accept()

		if err != nil {
			fmt.Printf("error listen accept, the error is : %v \n", err.Error())
			return
		}

		go doAccept(connect)
	}

}

func doAccept(connect net.Conn) {

	for {
		buff := make([]byte, 1024)
		len, err := connect.Read(buff)

		if err != nil {
			fmt.Printf("error read, the error is : %v \n", err.Error())
			return
		}

		fmt.Printf("recevie data is : %v", string(buff[0:len]))
	}

}