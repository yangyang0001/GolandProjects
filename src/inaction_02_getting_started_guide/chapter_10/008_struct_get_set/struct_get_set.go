package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type Person struct {
	FirsName string `json:"firs_name"`
	LastName string `json:"last_name"`
}

type Point struct {
	x float64
	y float64
}

type NamePoint struct {
	name  string
	point Point
}

type Log struct {
	message string
}

type Customer struct {
	Name string
	log  Log
}

func main() {

	fmt.Println("------------------------------ 1、set get 方法使用 ------------------------------")
	var person = Person{"zhang", "san"}
	fmt.Printf("person json = %v \n", jsonParse(person))

	person.SetFirsName("xiao")
	fmt.Printf("person json = %v \n", jsonParse(person))

	fmt.Println("------------------------------ 2、内嵌 struct 的方法1 ---------------------------")
	point := Point{3.0, 4.0}
	nameP := NamePoint{"mineP", point}

	fmt.Printf("nameP.point.Abs() = %v \n", nameP.point.Abs())

	fmt.Println("------------------------------ 3、内嵌 struct 的方法2 ---------------------------")
	log := Log{"1 - Yes we can!"}
	cus := Customer{"Mine Log", log}
	cus.log.Add("2 - After me the world will be a better place!")
	fmt.Println(cus.log.String())
}

func (p *Person) GetFirsName() string {
	return p.FirsName
}

func (p *Person) SetFirsName(firsname string) {
	p.FirsName = firsname
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func jsonParse(obj any) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(marshal)
}

func (log *Log) Add(msg string) {
	log.message += "\n" + msg
}

func (log *Log) String() string {
	return log.message
}

func (c *Customer) Log() *Log {
	return &c.log
}
