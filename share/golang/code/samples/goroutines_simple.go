package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello")
}

func main() {
	go hello()
	go func(s string) {
		fmt.Println(s)
	}("world")
	fmt.Println("!!")

	time.Sleep(100 * time.Millisecond)
}
