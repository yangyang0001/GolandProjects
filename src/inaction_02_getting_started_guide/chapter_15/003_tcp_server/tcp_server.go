package main

import (
	"fmt"
	"net"
	"strings"
)

var clients []string

func main() {

	fmt.Println("the server port is 8888, the server is starting ...")

	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Printf("connect error is %v \n", err.Error())
		return
	}

	for {
		// 接收客户端 连接
		connect, err := listen.Accept()
		if err != nil {
			fmt.Printf("client error is %v \n", err.Error())
			return
		}

		// 协程处理 链接
		go doAccept(connect)

	}

}

func doAccept(connect net.Conn)  {

	client_network := connect.RemoteAddr().String()
	clients = append(clients, client_network)
	fmt.Printf("clients = %v \n", clients)

	for {
		buff := make([]byte, 1024)

		read_num, err := connect.Read(buff)
		if err != nil {
			fmt.Printf("%v connect error is %v \n", client_network, err.Error())
			return
		}

		client_data := strings.Trim(string(buff[0:read_num]), "\n")
		fmt.Printf("%v data is :%v \n", client_network, client_data)

		if client_data == "WHO" {
			for _, client := range clients {
				//io.WriteString(connect, client + " ")
				connect.Write([]byte(client + " "))
			}
		}
	}

}
