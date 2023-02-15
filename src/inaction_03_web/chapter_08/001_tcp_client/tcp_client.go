package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	go DoSend(conn)

	go DoRead(conn)

	time.Sleep(time.Hour)
}

func DoSend(conn net.Conn) {
	fmt.Printf("the send data method invoke start, network = %v, address = %v \n",
		conn.LocalAddr().Network(), conn.LocalAddr().String())

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line = strings.Trim(line, "\n")
		if line == "exit" {
			fmt.Printf("the send data method invoke end,   network = %v, address = %v \n",
				conn.LocalAddr().Network(), conn.LocalAddr().String())
			return
		}
		io.WriteString(conn, line)
	}
}

func DoRead(conn net.Conn) {
	for {
		buff := make([]byte, 1024)
		read, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Back Message is : %v \n", string(buff[0:read]))
	}
}
