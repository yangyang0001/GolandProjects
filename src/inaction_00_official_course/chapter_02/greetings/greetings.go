package greetings

import "fmt"

func Hello(name string) string {

	return fmt.Sprintf("hello %v, welcome to the multi-module!", name)

}
