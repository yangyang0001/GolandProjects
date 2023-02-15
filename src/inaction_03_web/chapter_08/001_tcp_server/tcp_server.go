package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	fmt.Println("tcp server is starting, the server port is 8888 ...")

	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go DoAccept(conn)
	}

}

func DoAccept(conn net.Conn)  {
	for {
		buff := make([]byte, 1024)
		read, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("client network = %v, address = %v, data = %v \n",
			conn.RemoteAddr().Network(), conn.RemoteAddr().String(), string(buff[0:read]))

		DoSend(conn)
	}
}

func DoSend(conn net.Conn)  {
	nowtime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	io.WriteString(conn, "Hello " + conn.RemoteAddr().String() + ", Now time = " + nowtime)
}
