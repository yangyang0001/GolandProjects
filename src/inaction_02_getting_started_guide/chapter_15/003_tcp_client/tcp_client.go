package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("client is connecting ...")

	connect, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		fmt.Printf("connect error is %v \n", err.Error())
		return
	}

	go doWrite(connect)

	go doRead(connect)

	time.Sleep(time.Second * 1000)


}

func doWrite(connect net.Conn)  {

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("do write error is %v \n", err.Error())
			return
		}

		line = strings.Trim(line, "\n")

		if line == "exit" {
			fmt.Println("do write is end, the line is exit!")
			return
		}

		io.WriteString(connect, line)
	}

}

func doRead(connect net.Conn)  {

	for {
		buff := make([]byte, 1024)

		read_num, err := connect.Read(buff)
		if err != nil {
			fmt.Printf("connect error is %v \n", err.Error())
			return
		}

		server_data := strings.Trim(string(buff[0:read_num]), "\n")
		fmt.Printf("receive server data is %v \n", server_data)
	}

}
