package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Mine interface {
	SayHi()
}

type None interface {
	SayHello()
}

type MineNone interface {
	Mine
	None
}

type User struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}


func main() {

	xiaoming := User{"xiaoming", "111111"}
	xiaohong := User{"xiaohong", "123456"}
	
	var mines []MineNone
	mines = append(mines, &xiaoming)
	mines = append(mines, &xiaohong)

	for index, mine := range mines {
		fmt.Printf("index = %v, json = %v \n", index, jsonParse(mine))
		mine.SayHi()
		mine.SayHello()
		fmt.Println("-----------------------------------------------")
	}



}

func (u *User) SayHi()  {
	fmt.Printf("Hi everyone, I am %v \n", u.Name)
}

func (u *User) SayHello()  {
	fmt.Printf("Hello everyone, I am %v \n", u.Name)
}

func jsonParse(obj interface{}) string {
	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)
}
