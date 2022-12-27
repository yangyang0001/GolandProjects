package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func main() {

	fmt.Println("----------------------------- 1、bool 类型 和 逻辑运算 -------------------------------")
	var T, F bool = true, false
	fmt.Printf("T && F = %v \n", T && F)
	fmt.Printf("T || F = %v \n", T || F)
	fmt.Printf("T = %v, !T = %v \n", T, !T)
	fmt.Printf("F = %v, !F = %v \n", F, !F)

	fmt.Println("----------------------------- 2、混合类型的 使用案例  --------------------------------")
	var a int
	var b int64
	// fmt.Printf("a + b = %v \n", a+b)
	// 混合类型的变量 不能使用 上面的 + , 如果想使用 则必须进行类型转换, 如下:
	fmt.Printf("a + b = %v \n", int64(a)+b)

	const c = 10
	const d = 100.0
	fmt.Printf("c + d = %v \n", c+d) // 混合类型的常量 可以使用 +

	fmt.Println("----------------------------- 3、随机数 和 随机数的获取  -----------------------------")
	// 设置一个随机数种子
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(8))
	}

	fmt.Println("----------------------------- 4、unicode 几个字符方法  -----------------------------")

	fmt.Printf("IsLetter = %v \n", unicode.IsLetter(int32('A'))) // 是否是 字母
	fmt.Printf("IsLower  = %v \n", unicode.IsLower(int32('a')))  // 是否是 小写字母
	fmt.Printf("IsDigit  = %v \n", unicode.IsNumber(10))         // 是否是 十进制数字
	fmt.Printf("IsSpace  = %v \n", unicode.IsSpace(int32(13)))   // 是否是 空格

}
