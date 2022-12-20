package main

import "fmt"

type Person struct {
	Name string `json:"name"`
	Age  int
}

func (person *Person) String() string {
	return fmt.Sprintf("%v (%v years old)", person.Name, person.Age)
}

func main() {

	person := Person{"zhangsan", 10}
	fmt.Println(&person)

}
