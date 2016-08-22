// +build OMIT

package main

import (
	"fmt"
)

// START TYPE OMIT
type Foo int

// END TYPE OMIT

// START OMIT

func main() {
	var i int // Declare i, an integer, set to "zero value"
	var b bool
	var f float32
	var s string

	b = true
	f, s = 3.14, "hello"

	fmt.Printf("i: [%T] %v\n", i, i)
	fmt.Printf("b: [%T] %v\n", b, b)
	fmt.Printf("f: [%T] %v\n", f, f)
	fmt.Printf("s: [%T] %v\n", s, s)
}

// END OMIT
