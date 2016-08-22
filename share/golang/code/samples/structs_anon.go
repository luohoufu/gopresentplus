package main

import "fmt"

func main() {
	s := struct {
		field1 float64
		field2 bool
	}{
		3.14159,
		false,
	}

	fmt.Printf("s: [%T] %+v\n", s, s)
}
