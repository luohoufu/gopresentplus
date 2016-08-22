package main

import "fmt"

type Me struct {
	Adam   bool
	Schaub bool
}

func main() {
	m := Me{Adam: true, Schaub: true}
	fmt.Println(m)
}
