package utils

type Interface_ interface{
	F() int
}

type Struct_ struct {
	v int
}

var _ Interface_ = &Struct_{}

func (s *Struct_) F() int {
	return s.v
}

