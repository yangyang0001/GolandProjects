package main

import (
	"deepblue.com/inaction_02/chapter_10/struct_new"
	"fmt"
)

func main() {

	p := struct_new.NewPerson()
	p.Name = "zhangsan"
	p.Pass = "12345678"

	fmt.Println(p)

}
