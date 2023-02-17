package main

import (
	"errors"
	"fmt"
)

type MineError struct {
	message string
}

func (mineError *MineError) Error() string {
	return mineError.message
}


func main() {

	fmt.Println("-------------------------- 1、使用 errors.New() 自定义 error ----------------------------")
	err := errors.New("this is first error!")
	if err != nil {
		fmt.Printf("error = %v \n", err)
	}

	fmt.Println("-------------------------- 2、实现 error 接口, 自定义 error   ----------------------------")
	err = &MineError{"this is mine error!"}
	if err != nil {
		fmt.Printf("error = %v \n", err)
	}

}
