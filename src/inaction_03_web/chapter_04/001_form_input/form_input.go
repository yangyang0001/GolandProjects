package main

import (
	"crypto"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var LoginFilePath = "/Users/yangjianwei/GolandProjects/src/inaction_03_web/chapter_04/001_form_input/login.html"
var UploadFilePath = "/Users/yangjianwei/GolandProjects/src/inaction_03_web/chapter_04/001_form_input/upload.html"

var UploadPath = "/Users/yangjianwei/GolandProjects/src/inaction_03_web/chapter_04/upload/"

func main() {

	fmt.Println("Http Server starting... ------------------------------------------------------------------------")
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/upload", UploadHandler)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func HelloHandler(response http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(response, "Hello everyone, the web server is ready ...")
}

func LoginHandler(response http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	if request.Method == http.MethodGet {
		// 生成 token
		hash := crypto.MD5.New()
		hash.Write([]byte(strconv.Itoa(time.Now().Nanosecond())))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		// 将 token 写入表单中!
		temp, err := template.ParseFiles(LoginFilePath)
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(response, token)
	} else {
		//CheckForm(request)
		//FormMethd(request)

		token := request.FormValue("token")
		if token == "" {
			// 非法 token
			fmt.Println("token is not legitimate!")
			return
		} else {
			// 再次验证 token 是否合法 .....

			// 继续进行下面的代码
			PrintForm(request)

			// 转义输出到 client 端!
			template.HTMLEscape(response, []byte(request.FormValue("username")))

			// 非转义输出到 客户端, 如果 username 为 <script>alert()</script>, 则弹窗!
			//io.WriteString(response, request.FormValue("username"))
		}
	}

}


func UploadHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// 生成 token
		hash := crypto.MD5.New()
		hash.Write([]byte(strconv.Itoa(time.Now().Nanosecond())))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		// 将 token 写入表单中!
		temp, err := template.ParseFiles(UploadFilePath)
		if err != nil {
			log.Fatal(err)
		}
		temp.Execute(response, token)

	} else {
		sourcefile, header, err := request.FormFile("uploadfile")
		if err != nil {
			log.Fatal(err)
		}
		defer sourcefile.Close()

		fmt.Printf("header = %v \n", header)
		uploadfile, err := os.OpenFile(UploadPath + header.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer uploadfile.Close()

		io.Copy(uploadfile, sourcefile)
	}
}


func FormMethd(request *http.Request)  {
	fmt.Println("FormMethd invoke start ------------------------------------------------------------------------")
	values := request.Form
	values.Set("age", "18")				// 设置 key = value 值
	values.Add("username", "xiaoming")	// 往同一个 key 的 value 中添加值!
	values.Add("gender", "1")

	username := values.Get("username")	// 根据源码来看 return vs[0], 只能获取到 第一个参数!
	fmt.Printf("------------ values get method invoke, username = %v ------------ \n", username)

	key := "username"
	fmt.Printf("------------ values del method invoke, the key is %v ------------ \n", key)
	values.Del(key)

	for key, value := range values {
		fmt.Printf("key = %v, value = %v \n", key, value)
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

func CheckForm(request *http.Request)  {
	fmt.Println("CheckForm invoke start ------------------------------------------------------------------------")

	/* 必填校验 */
	if len(request.Form["username"][0]) == 0 {
		fmt.Println("username is nil")
		return
	}
	if len(request.Form["password"][0]) == 0 {
		fmt.Println("password is nil")
		return
	}
	if len(request.Form["age"][0]) == 0 {
		fmt.Println("age is nil")
		return
	}

	/* 数字校验, 当然了可以用正则表达式, 在 GoWeb.txt 中提供了 几种正则表达验证的样本代码 */
	age, err := strconv.Atoi(request.Form["age"][0])
	if err != nil {
		fmt.Println("age is not number!")
		return
	}
	if age > 100 {
		fmt.Println("age is too large, the age can not large than 100!")
		return
	}

	/* 下拉列表的校验 */
	fruits := []string {"apple", "pear", "banana"}
	fruit  := request.FormValue("fruit")
	if !IsExist(fruits, fruit) {
		fmt.Println("fruit is not exist!")
		return
	}

	/* 单选校验 */
	genders := []string {"0", "1"}
	gender  := request.FormValue("gender")
	if !IsExist(genders, gender) {
		fmt.Println("gender is not exist!")
		return
	}

	/* 复选校验, 这个校验很少做! 此处不做了 */

	/* 日期校验, 这个一般用前端的日期组件完成, 后端很少校验 */

	/* 手机号校验 */
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, request.FormValue("mobile")); !m {
		fmt.Println("mobile is not legitimate!")
		return
	}

	/* 邮件校验 */
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, request.FormValue("email")); !m {
		fmt.Println("email is not legitimate!")
		return
	}

	/* 身份证号校验 */
	if len(request.FormValue("usercard")) == 15 {
		//验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, request.FormValue("usercard")); !m {
			fmt.Println("usercard length is 15, but it is not legitimate!")
			return
		}
	} else {
		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, request.FormValue("usercard")); !m {
			fmt.Println("usercard length is 18, but it is not legitimate!")
			return
		}
	}

}


func IsExist(collections []string, value string) bool {
	for _, item := range collections {
		if item == value {
			return true
		}
	}
	return false
}

