package main

import "fmt"

type Pokemon struct {
	Type string
	No   int
}

type Squirtle struct {
	Pokemon
	Name string
}

func main() {
	s := Squirtle{Pokemon{"Water", 7}, "Squirtle"}
	var p Pokemon = s
	fmt.Printf("%#v\n", p)
}
