package learning_go_module_a

type Values struct {
	A int
	B int
}

func (v Values) Sum() int {
	return v.A + v.B
}

func (v Values) Sub() int {
	return v.A - v.B
}

func (v Values) Mul() int {
	return v.A * v.B
}

func (v Values) Div() int {
	return v.A / v.B
}
