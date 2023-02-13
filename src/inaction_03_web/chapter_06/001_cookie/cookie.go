package main

import (
	"crypto"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Http Server starting... ------------------------------------------------------------------------")
	http.HandleFunc("/login", LoginHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func LoginHandler(response http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// 获取 cookies
	GetCookies(request)

	if request.Method == http.MethodGet {
		// 生成 token
		hash := crypto.MD5.New()
		hash.Write([]byte(strconv.Itoa(time.Now().Nanosecond())))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		// 将 token 写入表单中!
		temp, err := template.ParseFiles("./login.html")
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(response, token)
	} else {

		token := request.FormValue("token")
		if token == "" {
			// 非法 token
			fmt.Println("token is not legitimate!")
			return
		} else {
			// 再次验证 token 是否合法 .....

			// 继续进行下面的代码
			PrintForm(request)

			// 设置 cookie 进行保存!
			expiration := time.Now()
			expiration = expiration.AddDate(0, 0, 1)
			fmt.Printf("expiration = %v \n", time.Unix(expiration.Unix(), 0).Format("2006-01-02 15:04:05"))

			cookie := http.Cookie{Name: "username", Value: request.FormValue("username"), Expires: expiration}
			http.SetCookie(response, &cookie)

			// 转义输出到 client 端!
			template.HTMLEscape(response, []byte(request.FormValue("username")))

			// 非转义输出到 客户端, 如果 username 为 <script>alert()</script>, 则弹窗!
			//io.WriteString(response, request.FormValue("username"))
		}
	}
}

func PrintForm(request *http.Request) {
	fmt.Println("PrintForm invoke start ------------------------------------------------------------------------")
	values := request.Form
	for key, _ := range values {
		//fmt.Printf("key = %v, value = %v \n", key, value)
		// 下面的代码使用了 转义:
		fmt.Printf("key = %v, value = %v \n", key, template.HTMLEscapeString(request.FormValue(key)))
	}
}

func GetCookies(request *http.Request)  {
	fmt.Println("GetCookies invoke start -----------------------------------------------------------------------")
	cookies := request.Cookies()
	for index, cookie := range cookies {
		fmt.Printf("index = %v, cookie is: %v \n", index, jsonParse(cookie))
	}
}

func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)

	if err != nil {
	log.Fatal(err)
	}

	return string(marshal)
}