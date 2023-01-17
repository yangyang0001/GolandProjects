package main

import "fmt"

func main() {

	fmt.Println("--------------------------------- 1、同级 catch error ---------------------------------")
	fmt.Println("1-recover_error_same_level method invoke start ...")
	recover_error_same_level()
	fmt.Println("1-recover_error_same_level method invoke end!")

	fmt.Println("--------------------------------- 2、上级 catch error ---------------------------------")
	fmt.Println("2-recover_error_upper_level method invoke start ...")
	recover_error_upper_level()
	fmt.Println("2-recover_error_upper_level method invoke end!")


}

func recover_error_same_level()  {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("error = %v \n", e)
		}
	}()
	panic("throw error to same level!")

}


func recover_error_upper_level()  {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("error = %v \n", e)
		}
	}()

	func() {
		panic("throw error to upper level!")
	}()
}

