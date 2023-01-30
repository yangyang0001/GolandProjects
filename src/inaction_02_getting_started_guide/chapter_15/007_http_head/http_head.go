package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
	"https://www.google.com/?hl=zh_CN",
	"http://golang.org/",
	"http://blog.golang.org/",
}


func main() {

	for _, url := range urls {
		response, err := http.Head(url)

		if err != nil {
			log.Fatal(err)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("response.status = %v \n", response.Status)
		fmt.Printf("response.body   = %v \n", string(data))
	}

}
