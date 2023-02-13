package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type VPN struct {
	XMLName xml.Name 	`xml:"servers"`
	Version string		`xml:"version,attr"`
	Servers []Server 	`xml:"server"`
}

type Server struct {
	XMLName xml.Name 	`xml:"server"`
	ServerName string 	`xml:"serverName"`
	ServerIP string 	`xml:"serverIP"`
}

func main() {

	fmt.Println("------------------------------------ 1、xml 解析 案例 ------------------------------------")
	xmlfile, err := os.Open("./vpn.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlfile.Close()

	data, err := ioutil.ReadAll(xmlfile)
	if err != nil {
		log.Fatal(err)
	}

	var vpn VPN
	err = xml.Unmarshal(data, &vpn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("VPN -------------------------------------------------------------------")
	fmt.Printf("vpn.XMLName = %v \n", vpn.XMLName)
	fmt.Printf("vpn.Version = %v \n", vpn.Version)

	fmt.Println("Servers ---------------------------------------------------------------")
	servers := vpn.Servers
	for index, server := range servers {
		fmt.Printf("index = %v, server.XMLName    = %v \n", index, server.XMLName)
		fmt.Printf("index = %v, server.ServerName = %v \n", index, server.ServerName)
		fmt.Printf("index = %v, server.ServerIP   = %v \n", index, server.ServerIP)
	}

	fmt.Println("------------------------------------ 2、xml 打印 案例 ------------------------------------")
	indent, err := xml.MarshalIndent(vpn, "", "")
	if err != nil {
		log.Fatal(err)
	}
	// 打印出 xml 头部格式
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(indent)

}


func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}
