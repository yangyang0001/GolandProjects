package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Response struct {
	result int
}

type Request struct {
	param1 int `json:"param1"`
	param2 int `json:"param2"`
	channel chan *Response
}


func main() {

	var requests = make([]Request, 100)

	for i := 0; i < 10; i++ {
		var request Request
		request.param1 = i
		request.param2 = i
		requests[i] = request
		HttpRequest(&requests[i])
	}

	fmt.Println("----------------------------------- check response --------------------------------")

	for i := 0; i < 10; i++ {
		response := <- requests[i].channel
		fmt.Printf("i = %v, result = %v \n", i, response.result)
	}

	time.Sleep(time.Second * 2)

}

func HttpRequest(request *Request) {

	fmt.Printf("http request method start, param1 = %v, param2 = %v \n", request.param1, request.param2)

	go DoRequest(request)

	fmt.Printf("http request method   end, param1 = %v, param2 = %v \n", request.param1, request.param2)
}

func DoRequest(request *Request) {

	go func() {
		request.channel = make(chan *Response)
		request.channel <- &Response{request.param1 + request.param2}
	}()

}

func jsonParse(obj interface{}) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)

}


