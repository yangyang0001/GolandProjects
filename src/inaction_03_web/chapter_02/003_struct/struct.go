package main

import "fmt"

type person struct {
	name string
	age int
}

func Older(p1, p2 *person) (*person, int) {

	if p1.age >= p2.age {
		return p1, p1.age - p2.age
	} else {
		return p2, p2.age - p1.age
	}

}


func main() {

	fmt.Println("-------------------------------------- 1、结构体案例1 ----------------------------------------")
	var p person
	p.name = "zhangsan"
	p.age = 25
	fmt.Printf("p.name = %v, p.age = %v \n", p.name, p.age)

	fmt.Println("-------------------------------------- 2、结构体案例2 ----------------------------------------")
	tom  := person{"tom", 18}
	bob  := person{"bob", 25}
	paul := person{"paul", 43}

	tb_Older, tb_diff := Older(&tom, &bob)
	tp_Older, tp_diff := Older(&tom, &paul)
	bp_Older, bp_diff := Older(&bob, &paul)

	fmt.Printf("Of %s and %s, %s is older by %d years \n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years \n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years \n", bob.name, paul.name, bp_Older.name, bp_diff)

}
