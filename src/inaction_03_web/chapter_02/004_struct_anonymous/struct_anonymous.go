package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
	phone string
}

type Student struct {
	Human
	int
	speciality string
	phone string
}


func main() {

	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120, "18765826666"}, 100, "Computer Science", "18765828888"}
	fmt.Println("------------------------------------ show Huame properties --------------------------")
	fmt.Printf("mark.Human.name   = %v \n", mark.Human.name)
	fmt.Printf("mark.Human.age    = %v \n", mark.Human.age)
	fmt.Printf("mark.Human.weight = %v \n", mark.Human.weight)
	fmt.Printf("mark.Human.phone  = %v \n", mark.Human.phone)

	fmt.Println("------------------------------------ show mark  properties --------------------------")
	fmt.Printf("mark.int    = %v \n", mark.int)
	fmt.Printf("mark.specia = %v \n", mark.speciality)
	fmt.Printf("mark.phone  = %v \n", mark.phone)


	fmt.Println("------------------------------------ change speciality ------------------------------")
	mark.speciality = "AI"
	fmt.Printf("mark.specia = %v \n", mark.speciality)

	fmt.Println("------------------------------------ change age -------------------------------------")
	mark.age = 28
	fmt.Printf("mark.Human.age    = %v \n", mark.Human.age)

	fmt.Println("------------------------------------ change weight ---------------------------------")
	mark.weight += 60
	fmt.Printf("mark.Human.weight = %v \n", mark.Human.weight)



}
