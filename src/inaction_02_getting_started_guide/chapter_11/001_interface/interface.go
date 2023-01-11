package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	length float64
	width  float64
}

type valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

func main() {

	fmt.Println("-------------------------------------------- 1、接口案例1 ----------------------------------------")
	square := Square{5.0}
	var shaper Shaper = &square
	area := shaper.Area()
	fmt.Printf("area = %v \n", area)

	fmt.Println("-------------------------------------------- 2、接口案例2 ----------------------------------------")
	square1 := Square{6.0}
	rectang := Rectangle{10, 5}

	shapers := []Shaper{&square1, &rectang}

	for _, shaper := range shapers {
		fmt.Printf("shaper.Area() = %v \n", shaper.Area())
	}

	fmt.Println("-------------------------------------------- 3、接口案例3 ----------------------------------------")
	var able valuable = &stockPosition{"GOOG", 577.20, 4}
	showValue(able)
	able = &car{"BMW", "M3", 66500}
	showValue(able)
}

func (square *Square) Area() float64 {
	return square.side * square.side
}

func (rect *Rectangle) Area() float64 {
	return rect.length * rect.width
}

func (position *stockPosition) getValue() float32 {
	return position.sharePrice * position.count
}

func (car *car) getValue() float32 {
	return car.price
}

func showValue(able valuable) {
	fmt.Printf("valuable get value = %v \n", able.getValue())
}
