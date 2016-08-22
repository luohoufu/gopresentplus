// +build OMIT
package main

import "fmt"

type Address struct {
	HouseNumber int
	StreetName  string
}

type Person struct {
	Name    string
	Address Address
}

func main() {
	var addr Address
	addr = ...?
	var p Person
	p = ...?

	fmt.Println("My name is", p.Name)
	fmt.Println("My address is", p.Address)
}
