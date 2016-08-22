// +build OMIT
// OMIT
// OMIT
package main

import "fmt"

type Foo struct {
	field int
}

type Bar struct {
	field int
}

func main() {
	var foo Foo
	bar := Bar{field: 1}
	foo = bar

	foo.field = 5
	fmt.Printf("foo: [%T] %+v\n", foo, foo)
}
