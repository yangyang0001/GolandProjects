package main

import "fmt"

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

type CameraPhone struct {
	Camera
	Phone
}

type Base struct {
}

type Voodoo struct {
	Base
}

func main() {
	fmt.Println("------------------------------ 1、方法多继承案例实现 ---------------------------")
	cp := CameraPhone{Camera{}, Phone{}}
	fmt.Printf("cp.TakeAPicture()) = %v \n", cp.TakeAPicture())
	fmt.Printf("cp.Call()) = %v \n", cp.Call())

	fmt.Println("------------------------------ 2、多继承的疑问问题 ---------------------------")
	voo := Voodoo{Base{}}
	voo.Base.Magic()
	voo.Magic()
	voo.MoreMagic()

}

func (p *Phone) Call() string {
	return "Ring Ring"
}

func (base *Base) Magic() {
	fmt.Println("base magic")
}

func (self *Base) MoreMagic() {
	fmt.Println("more magic invoke start ...")
	self.Magic()
	self.Magic()
	fmt.Println("more magic invoke end   ...")
}

func (voo *Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
