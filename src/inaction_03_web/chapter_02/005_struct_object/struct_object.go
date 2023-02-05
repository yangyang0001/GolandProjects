package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func main() {

	fmt.Println("-------------------------------------- 1、函数案例 ----------------------------------------")
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	fmt.Printf("the area of r1 is : %v \n", area(r1))
	fmt.Printf("the area of r2 is : %v \n", area(r2))

	fmt.Println("-------------------------------------- 2、结构体方法 --------------------------------------")
	r3 := Rectangle{12, 2}
	c1 := Circle{10.0}
	fmt.Printf("the area of r3 is : %v \n", r3.area())
	fmt.Printf("the area of c1 is : %v \n", c1.area())

}


func area(r Rectangle) float64 {
	return r.width*r.height
}

func (r *Rectangle) area() float64  {
	return r.height * r.height
}

func (c *Circle) area() float64  {
	return math.Pi * math.Pow(c.radius, 2)
}

