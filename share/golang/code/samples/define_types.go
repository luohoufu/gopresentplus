// +build OMIT
// OMIT
// OMIT
package main

type Foo int

func main() {
	var foo Foo
	var i int
	foo = 1
	i = foo
	i = int(foo)
	f := 3.14
	i = int(f)
}
