package main

import "fmt"

type MineI interface {
	Getname() string
	Setname(param string)
}

type MineN interface {
	Show()
}

type MineM interface {
	MineI
	MineN
}

type Mine struct {
	Name string
}

func main() {

	var mineM MineM = &Mine{"zhangsan"}
	mineM.Show()
	mineM.Setname("wangwu")
	mineM.Show()

}

func (mine *Mine) Getname() string {
	return mine.Name
}

func (mine *Mine) Setname(name string) {
	mine.Name = name
}

func (mine *Mine) Show() {
	fmt.Printf("mine.Name = %v \n", mine.Name)
}
