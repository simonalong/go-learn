package mik

var test = 12

type TestStruct2 interface {
	Namess() []string
}

type TestStruct33 struct {
}

func (TestStruct33) Namess() []string {
	return nil
}
