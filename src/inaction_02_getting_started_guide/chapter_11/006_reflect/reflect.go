package main

import (
	"fmt"
	"reflect"
)

type Mine struct {
	s1, s2, s3 string
}

type T struct {
	A int
	B string
}

func main() {

	//fmt.Println("------------------------------------------ 1、反射案例1 ------------------------------------------")
	//var i int = 5
	//
	//typeI := reflect.TypeOf(i)
	//valuI := reflect.ValueOf(i)
	//kindI := reflect.Kind(i)
	//
	//fmt.Printf("typeI = %v, valuI = %v, kindI = %v \n", typeI, valuI, kindI)
	//
	//var x float64 = 3.4
	//typeX := reflect.TypeOf(x)
	//valuX := reflect.ValueOf(x)
	//kindX := reflect.Kind(x)
	//inteX := valuX.Interface()
	//fmt.Printf("typeX = %v, valuX = %v, kindX = %v, inteX = %v \n", typeX, valuX, kindX, inteX)
	//
	//fmt.Println("------------------------------------------ 2、通过反射设置 ------------------------------------------")
	//var m float64 = 3.4
	//valuM := reflect.ValueOf(m)
	//setmF := valuM.CanSet()
	//
	//if setmF {
	//	fmt.Println("m can set, m = %v \n", m)
	//}
	//
	//valuPM := reflect.ValueOf(&m)
	//typePM := reflect.TypeOf(&m)
	//setmpF := valuPM.CanSet()
	//elemPM := valuPM.Elem()
	//fmt.Printf("valuPM = %v, typePM = %v, setpmF = %v, elemPM = %v \n", valuPM, typePM, setmpF, elemPM)
	//
	//// 这是设置值的 格式: 必须按照以下方式进行修改!
	//setmF = reflect.ValueOf(&m).Elem().CanSet()
	//if setmF {
	//	reflect.ValueOf(&m).Elem().SetFloat(3.1415926)
	//	fmt.Printf("m = %v \n", m)
	//}
	//
	//fmt.Println("------------------------------------------ 3、反射结构体 ------------------------------------------")

	var mine Mine = Mine{"Ada", "Go", "Oberon"}

	valuMine := reflect.ValueOf(mine)
	typeMine := reflect.TypeOf(mine)
	kindMine := valuMine.Kind()

	fmt.Printf("valuMine = %v, typeMine = %v, kindMine = %v \n", valuMine, typeMine, kindMine)

	numField := valuMine.NumField()
	for i := 0; i < numField; i++ {
		fmt.Printf("index = %v, value = %v \n", i, valuMine.Field(i))
	}

	numMethd := valuMine.NumMethod()
	//numMethd := typeMine.NumMethod()
	fmt.Printf("numMethd = %v \n", numMethd)
	for i := 0; i < numMethd; i++ {
		valuMine.Method(i).Call(nil)
	}

	fmt.Println("------------------------------------------ 4、反射设置结构体 -------------------------------------")
	var t T = T{1, "zhangsan"}
	reflect.ValueOf(&t).Elem().Field(0).SetInt(100)
	reflect.ValueOf(&t).Elem().Field(1).SetString("wangwu")

	fmt.Println(t)
}

func (mine Mine) Show() {
	fmt.Printf("show method invoke, mine = %v \n", mine)
}
