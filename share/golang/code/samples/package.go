package mypackage

struct MyType int

func (mt MyType) String() string {
	return "foo"
}

struct secretType float32

func (st secretType) String() string {
	return "bar"
}