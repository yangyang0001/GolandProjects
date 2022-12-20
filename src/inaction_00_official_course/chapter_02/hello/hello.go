package main

import (
	"deepblue.com/greetings"
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {

	fmt.Println("main method invoke ...")
	fmt.Println(greetings.Hello("zhangsan"))
	fmt.Println(stringutil.Reverse("zhangsan"))
	fmt.Println(stringutil.UpperCase("zhangsan"))

}
