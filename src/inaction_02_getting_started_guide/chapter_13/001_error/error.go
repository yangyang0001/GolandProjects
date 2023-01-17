package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	fmt.Println("--------------------------------- 1、自定义 error ---------------------------------")
	err := errors.New("this is mine error!")
	fmt.Printf("error = %v \n", err)

	fmt.Println("--------------------------------- 2、使用自定义 error ------------------------------")
	res, err := sqrt(-1)
	//res, err := sqrt(100)
	if err != nil {
		fmt.Printf("Error = %v \n", err)
		os.Exit(1)
	} else {
		fmt.Printf("sqrt method res = %v \n", res)
	}

}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("the param num cannot be less than 0!")
	}

	return math.Sqrt(num), nil
}
