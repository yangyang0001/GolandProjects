package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	fmt.Println("------------------------------------ 1、正则表达式 校验案例 ------------------------------------")
	ip := "192.168.188.255"
	fmt.Printf("IP = %v, IsIP = %v \n", ip, IsIP(ip))

	num := "1000"
	fmt.Printf("num = %v, IsNum = %v \n", num, IsNum(num))

	han := "模式"
	fmt.Printf("han = %v, IsHan = %v \n", han, IsHan(han))

	eng := "zhangsan"
	fmt.Printf("eng = %v, IsEng = %v \n", eng, IsEng(eng))

	email := "2017585616@qq.com"
	fmt.Printf("email = %v, IsEmail = %v \n", email, IsEmail(email))

	mobile := "18765829090"
	fmt.Printf("mobile = %v, IsMobile = %v \n", mobile, IsMobile(mobile))

	card := "371523199001066253"
	fmt.Printf("card = %v, IsCard = %v \n", card, IsCard(card))

	fmt.Println("------------------------------------ 2、正则表达式 获取内容 ------------------------------------")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	src = re.ReplaceAllString(src, "")
	//去除HTMLUnscape的STYLE
	re, _ = regexp.Compile(`&lt;style[\S\s]+?&lt;/style&gt;`)
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	src = re.ReplaceAllString(src, "")
	//去除HTMLUnsapce的SCRIPT
	re, _ = regexp.Compile(`&lt;script[\S\s]+?&lt;/script&gt;`)
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`<[\S\s]+?>`)
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile(`\s{2,}`)
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))


}

/* IP 校验 */
func IsIP(ip string) bool {
	match, err := regexp.MatchString("^((25[0-5]|2[0-4]\\d|[01]?\\d\\d?)\\.){3}(25[0-5]|2[0-4]\\d|[01]?\\d\\d?)$", ip)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

/* 数字校验 */
func IsNum(num string) bool {
	match, err := regexp.MatchString("^[0-9]+$", num)
	if err != nil {
		log.Fatal(err)
	}
	return match
}


/* 中文校验 */
func IsHan(han string) bool {
	match, err := regexp.MatchString("^\\p{Han}+$", han)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

/* 英文校验 */
func IsEng(eng string) bool {
	match, err := regexp.MatchString("^[a-zA-Z]+$", eng)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

/* 邮件校验 */
func IsEmail(email string) bool {
	match, err := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, email)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

/* 手机校验 */
func IsMobile(mobile string) bool {
	match, err := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

/* 身份证校验 */
func IsCard(card string) bool {
	var err error
	var flag bool
	if len(card) == 15 {
		flag, err = regexp.MatchString(`^(\d{15})$`, card)
	} else {
		flag, err = regexp.MatchString(`^(\d{17})([0-9]|X)$`, card)
	}
	if err != nil {
		log.Fatal(err)
	}
	return flag
}
