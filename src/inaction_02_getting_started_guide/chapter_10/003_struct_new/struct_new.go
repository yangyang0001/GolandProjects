package struct_new

type person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func NewPerson() *person {
	p := new(person)
	return p
}
