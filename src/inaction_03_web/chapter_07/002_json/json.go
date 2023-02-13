package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

func main() {

	fmt.Println("------------------------------------ 1、xml 解析 案例 ------------------------------------")
	servers := []Server{
		{"HongKong_VPN", "127.0.0.1"},
		{"Shanghai_VPN", "127.0.0.2"},
		{"Beijing_VPN", "127.0.0.3"},
	}

	fmt.Printf("servers json = %v \n", jsonParse(servers))
	fmt.Println("------------------------------------ 2、xml 解析 案例 ------------------------------------")

}

func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}
