package main

import "fmt"

func main() {
	var slice1 []int
	fmt.Printf("slice1: [%T] %v\n", slice1, slice1)

	slice2 := make([]int, 0)
	fmt.Printf("slice2: [%T] %v\n", slice2, slice2)

	fmt.Printf("slice2: Len(%d) Cap(%d)\n", len(slice2), cap(slice2))
}
