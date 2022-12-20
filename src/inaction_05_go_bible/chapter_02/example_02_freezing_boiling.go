package main

import "fmt"

func main() {

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g℉  = %g℃ \n", freezingF, transF2C(freezingF))
	fmt.Printf("%g℉  = %g℃ \n", boilingF, transF2C(boilingF))

	var a string = "zhangsan"

	fmt.Println(a)
}

func transF2C(f float64) float64 {
	return (f - 32) * 5 / 9
}
