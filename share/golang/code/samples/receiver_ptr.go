package main

import "fmt"

type MyInt int

func (i MyInt) Set(n MyInt) {
	i = n
}

func main() {
	var i MyInt = 2
	i.Set(18)
	fmt.Println(i)
}
