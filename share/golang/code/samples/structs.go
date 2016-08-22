// +build OMIT
// OMIT
package main

import "fmt"

type Bar struct {
	Field1 int
	Field2 float32
	Name   string
}

func main() {
	var bar Bar
	barOne := Bar{1, 2, "Bar One"}  // unnamed assignment
	barTwo := Bar{Name: "Bar None"} // named assignment

	fmt.Printf("bar   : [%T] %+v\n", bar, bar)
	fmt.Printf("barOne: [%T] %#v\n", barOne, barOne)
	fmt.Printf("barTwo.Name: %s\n", barTwo.Name)
}
