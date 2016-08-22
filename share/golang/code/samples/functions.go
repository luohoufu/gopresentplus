// +build OMIT

package main

import "fmt"

func SayMyName(name string) {
	fmt.Println(name)
}

func WhatIsMyName() string {
	return "Heisenberg"
}

func main() {
	name := WhatIsMyName()
	SayMyName(name)
}
