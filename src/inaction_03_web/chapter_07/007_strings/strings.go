package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("------------------------------------ 1、strings 相关方法 ------------------------------------")
	// 包含
	fmt.Println(strings.Contains("seafood", "foo"))		// true
	fmt.Println(strings.Contains("seafood", "bar"))		// false
	fmt.Println(strings.Contains("seafood", ""))			// true
	fmt.Println(strings.Contains("", ""))				// true

	// 拼接
	arr := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(arr, "-"))

	// 下标
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	// 重复字符串
	fmt.Println("ba" + strings.Repeat("na", 2))

	// 替换字符串: 在字符串s中，把 old 字符串 替换为 new 字符串, n表示替换的次数, 小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))


	fmt.Printf("%q \n", strings.Split("a,b,c", ","))
	fmt.Printf("%q \n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q \n", strings.Split(" xyz ", ""))
	fmt.Printf("%q \n", strings.Split("", "Bernardo O'Higgins"))

	// 去掉 左+右两边的 !
	fmt.Printf("[%v] \n", strings.Trim("!!! Achtung !!!", "!")) 	// [ Achtung ]

	fmt.Printf("fields = %q \n", strings.Fields("  foo bar  baz   "))	// fields = ["foo" "bar" "baz"]

	fmt.Println("------------------------------------ 2、字符串转换 相关方法 -----------------------------------")
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)

}

func checkError(err error){
	if err != nil{
		fmt.Println(err)
	}
}
