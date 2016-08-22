package main

import "fmt"

func main() {
	b := make([]bool, 1)
	fmt.Printf("Len(%d) Cap(%d): %v\n", len(b), cap(b), b)

	for i := 0; i < 10; i++ {
		b = append(b, i%2 == 0)
		fmt.Printf("Len(%d) Cap(%d): %v\n", len(b), cap(b), b)
	}
}
