package main

type M map[string]string

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	/**
	结构体 只能使用 new()  不能使用 make()
	map   只能使用 make() 不能使用 new()
	*/

}
