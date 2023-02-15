
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)


func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var receiveMsg string
		if err = websocket.Message.Receive(ws, &receiveMsg); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Printf("ReceiveMsg = %v \n", receiveMsg)

		sendMsg := "Server send message " + receiveMsg

		fmt.Printf("Sending to client message:  %v, \n", sendMsg)

		if err = websocket.Message.Send(ws, sendMsg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}