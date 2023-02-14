package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Username string
	Password string
}

func main() {

	fmt.Println("http server starting ------------------------------------------------------------------------")
	http.HandleFunc("/login", LoginHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func LoginHandler(response http.ResponseWriter, request *http.Request)  {

	if request.Method == http.MethodGet {
		fmt.Println("request get method invoke ------------------------------------------------------------------")

		// 解析字段:
		temp := template.New("Login")
		temp.Parse("{{.Username}}")
		temp.Parse("{{.Password}}")

		temp, err := template.ParseFiles("./login.html")
		if err != nil {
			log.Fatal(err)
		}

		temp.Execute(response, Person{"wangwu", "123456"})
	} else {
		fmt.Println("request post method invoke -----------------------------------------------------------------")
		PrintForm(request)
	}


}

func PrintForm(request *http.Request) {

}
