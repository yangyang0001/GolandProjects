package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	fmt.Println("udp server is starting, the server port is 8888 ...")
	addr, err := net.ResolveUDPAddr("udp4", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", addr)

	for {
		DoRead(conn)
	}

}

func DoRead(conn *net.UDPConn)  {
	buff := make([]byte, 1024)
	read, addr, err := conn.ReadFromUDP(buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client network = %v, address = %v, data = %v \n",
		addr.Network(), addr.String(), string(buff[0:read]))

	DoSend(conn, addr)
}

func DoSend(conn *net.UDPConn, addr *net.UDPAddr)  {
	nowtime := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	conn.WriteToUDP([]byte(nowtime), addr)
}

func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}
