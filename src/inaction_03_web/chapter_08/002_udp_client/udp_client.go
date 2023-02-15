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

	addr, err := net.ResolveUDPAddr("udp4", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	go DoSend(conn)

	go DoRead(conn)

	time.Sleep(time.Hour)
}

func DoSend(conn *net.UDPConn) {
	fmt.Printf("the send method invoke start, network = %v, address = %v \n",
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

func DoRead(conn *net.UDPConn) {
	fmt.Printf("the read method invoke start, network = %v, address = %v \n",
		conn.LocalAddr().Network(), conn.LocalAddr().String())
	for {
		buff := make([]byte, 1024)
		read, _, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Back Message is : %v \n", string(buff[0:read]))
	}
}
