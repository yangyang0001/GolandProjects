package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


/* 这个结构会保存解析后的返回数据. 他们会形成有层级的 XML, 可以忽略一些无用的数据 */
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {


	fmt.Println("----------------------------------- 1、http get 方法 ----------------------------------")
	http_get()

	fmt.Println("----------------------------------- 2、http get struct -------------------------------")
	http_get_struct()

}

func http_get()  {

	var url = "http://www.google.com"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("google data = %v \n", string(data))

}

func http_get_struct()  {

	// 发起请求查询推特 Goodland 用户的状态
	response, _ := http.Get("http://twitter.com/users/Googland.xml")

	// 初始化 XML 返回值的结构
	user := User{xml.Name{"", "user"}, Status{""}}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// 将 XML 解析为我们的结构
	xml.Unmarshal(data, &user)

	fmt.Printf("status: %v \n", user.Status.Text)
}



