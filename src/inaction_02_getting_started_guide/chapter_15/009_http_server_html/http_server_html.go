package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form =
	`<html>
		<body>
			<form action="#" method="post" name="bar">
				<input type="text" name="in" />
				<input type="submit" value="submit"/>
			</form>
		</body>
	</html>`

func main() {

	http.HandleFunc("/simple", SimpleHandler)
	http.HandleFunc("/form", FormHandler)

	err := http.ListenAndServe("localhost:8088", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func SimpleHandler(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("simple handler invoke ...")
	io.WriteString(response, "<h1>Welcome to SimpleHandler!</h1>")
}

func FormHandler(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("form handler invoke ...")
	response.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case http.MethodGet:
		io.WriteString(response, form)
	case http.MethodPost:
		io.WriteString(response, request.FormValue("in"))
	default:
		io.WriteString(response, "Welcome to FormHandler!")
	}
}


