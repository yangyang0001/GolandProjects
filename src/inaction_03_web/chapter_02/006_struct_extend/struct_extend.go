package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employe struct {
	Human
	company string
}

func main() {

	student := Student{Human{"Mar", 25, "18765828888"}, "MIT"}
	employe := Employe{Human{"Sam", 40, "18765826666"}, "Google"}

	student.Human.SayHi()
	employe.Human.SayHi()
	employe.SayHi()

}

func (h *Human) SayHi()  {
	fmt.Printf("Hello everyone, I am %v, you can call me on %v \n", h.name, h.phone)
}

func (e * Employe) SayHi()  {
	fmt.Printf("Hello everyone, I am %v, I work at %v, you can call me on %v \n", e.name, e.company, e.phone)
}
