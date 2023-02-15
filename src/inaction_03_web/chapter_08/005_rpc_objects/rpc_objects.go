package _05_rpc_objects

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (mine *Mine) SayHello(param *Mine, result *string) error {
	*result = "Hello, Mine.Name = " + param.Name + ", Mine.Pass = " + param.Pass
	return nil
}
