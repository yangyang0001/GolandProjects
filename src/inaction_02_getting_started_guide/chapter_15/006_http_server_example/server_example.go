package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
当请求 http://localhost:9999/hello/Name 时，响应：hello Name（Name 需是一个合法的姓，比如 Chris 或者 Madeleine）
当请求 http://localhost:9999/shouthello/Name 时，响应：hello NAME
 */
func main() {

	http.HandleFunc("/", HttpServer)

	err := http.ListenAndServe("localhost:9999", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func HttpServer(response http.ResponseWriter, request *http.Request)  {

	paths := strings.Split(request.URL.Path, "/")

	if paths[1] == "hello" || paths[1] == "shouthello" {
		fmt.Fprintf(response, "hello " + paths[2])

	} else {
		fmt.Fprintf(response, "welcome to go web!")
	}

}
