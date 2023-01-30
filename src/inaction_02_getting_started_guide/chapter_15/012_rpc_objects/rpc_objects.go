package _12_rpc_objects

type MineObject struct {
	M int
	N int
}

func (*MineObject) Multi(param *MineObject, result *int) error {
	*result = param.M * param.N
	return nil
}
