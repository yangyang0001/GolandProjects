package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 访问路径为 http://localhost:8080/mine/hello/xxx, 这里的 pattern 为 addr 后面拼接的路径! addr 为 host:port 的格式!
	http.HandleFunc("/mine/hello/", HttpServer)

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func HttpServer(response http.ResponseWriter, request *http.Request)  {
	paths := strings.Split(request.URL.Path, "/")
	fmt.Printf("Path = %v, paths = %v, length = %v \n", request.URL.Path, paths, len(paths))
	fmt.Fprintf(response, "hello " + paths[len(paths)-1])
}




