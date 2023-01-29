package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	connect, err := net.Dial("tcp", "localhost:50000")

	if err != nil {
		fmt.Printf("client dial error is : %v \n", err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("reader string error is : %v \n", err.Error())
			return
		}

		if strings.Trim(line, "\n") == "exit" {
			return
		}

		_, err = connect.Write([]byte(line))

		if err != nil {
			fmt.Printf("connect write error is : %v \n", err.Error())
		}

	}

}
