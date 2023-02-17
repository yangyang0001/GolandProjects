package main

import "fmt"

type MineError struct {
	message string
}

func (mineError *MineError) Error() string {
	return mineError.message
}

func main() {


	fmt.Println("-------------------------- 1、同级捕获 error ----------------------------")

	SameLevelCatch()
	fmt.Println("the first  error catched!")

	fmt.Println("-------------------------- 2、冒泡捕获 error ----------------------------")
	BubbleLevelCatch()
	fmt.Println("the bubble error catched!")
}

func SameLevelCatch() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error = %v \n", err)
		}
	}()
	panic(MineError{"this is the same level error!"})
}

func BubbleLevelCatch()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error = %v \n", err)
		}
	}()

	func() {
		panic(MineError{"this is the bubble level error!"})
	}()

}

