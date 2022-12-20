package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var c, python, java = true, false, "no!" // := 不能在函数外使用, 只能用 var 这种方式进行声明!
var a, b int = 10, 100

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("my favorite number is : %v \n", rand.Intn(10))
	fmt.Printf("math.PI is %v \n", math.Pi)
	fmt.Printf("add1 method invoke, %v + %v = %v \n", 42, 13, add1(42, 13))
	fmt.Printf("add2 method invoke, %v + %v = %v \n", 42, 13, add2(42, 13))

	x := "hello"
	y := "world"
	s1, s2 := swap(x, y)
	fmt.Printf("swap method invoke, [%v %v] swap is [%v %v] \n", x, y, s1, s2)

	fmt.Println(split(17))

	i, j := 1, 2
	fmt.Printf("c = %v, python = %v, java = %v, i = %v, j = %v \n", c, python, java, i, j)
	fmt.Println(a, b)

	fmt.Printf("ToBe   type is %T, value is %v \n", ToBe, ToBe)
	fmt.Printf("MaxInt type is %T, value is %v \n", ToBe, MaxInt)
	fmt.Printf("z      type is %T, value is %v \n", ToBe, z)

	// 类型转化
	transf()

	// 常量方法
	constMethod()

}

func init() {
	fmt.Println("init method invoke ...")
	rand.Seed(time.Now().UnixNano())
}

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func transf() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("x = %v, y = %v, z = %v \n", x, y, z)
}

func constMethod() {
	const Big = 1 << 100
	const Sma = Big >> 99

	fmt.Printf("Sma is %v \n", needInt(Sma))
	fmt.Printf("Sma is %v \n", needFloat(Sma))
	fmt.Printf("Big is %v \n", needFloat(Big))

}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
