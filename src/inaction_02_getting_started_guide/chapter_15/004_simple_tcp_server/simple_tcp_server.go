package main

import (
	"fmt"
	"log"
	"net"
)


var clients []string

func main() {

	listener := initserver("tcp", "localhost:8888")

	for {
		connect, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go doaccept(connect)
	}

}

func initserver(protocol string, address string) net.Listener {

	fmt.Println("the tcp server starting ...")

	listener, err := net.Listen(protocol, address)

	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func doaccept(connect net.Conn)  {

	client := connect.RemoteAddr().String()
	clients = append(clients, client)
	fmt.Printf("clients = %v \n", clients)

	for {
		buff := make([]byte, 1024)
		read_num, err := connect.Read(buff)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v : %v \n", client, string(buff[0:read_num]))
	}

}


