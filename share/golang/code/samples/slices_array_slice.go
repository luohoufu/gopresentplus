package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:4]

	fmt.Printf("array: [%T] %v\n", arr, arr)
	fmt.Printf("slice: [%T] %v\n", slice, slice)
}
