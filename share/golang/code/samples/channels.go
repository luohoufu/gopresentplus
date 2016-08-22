package main

import (
	"fmt"
	"time"
)

func Dori(sig <-chan bool) {
	for keepGoing := true; keepGoing; {
		select {
		case keepGoing = <-sig:
			fmt.Println("...okay")
		default:
			fmt.Println("Just keep swimming")
		}
	}
}

func main() {
	c := make(chan bool)
	go Dori(c)
	time.Sleep(500 * time.Millisecond)
	c <- false
}
