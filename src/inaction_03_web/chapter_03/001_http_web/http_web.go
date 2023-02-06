package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println("--------------------------- http web server starting ... -------------------------")
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe("localhost:9090", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func SayHello(response http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error is %v \n", err)
	}

	fmt.Printf("request.Form             = %v \n", request.Form)
	fmt.Printf("request.URL              = %v \n", request.URL)
	fmt.Printf("request.URL.Path         = %v \n", request.URL.Path)
	fmt.Printf("request.URL.Scheme       = %v \n", request.URL.Scheme)
	fmt.Printf("request.Form['url_long'] = %v \n", request.Form["url_long"])

	for key, value := range request.Form {
		fmt.Printf("key = %v, value = %v \n", key, value)
	}


	io.WriteString(response, "Hello everyone, I am web server!")
}


