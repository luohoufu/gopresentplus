// +build OMIT

package main

import (
	"fmt"
	"sync"
)

func do(s string, i int, wg *sync.WaitGroup) {
	fmt.Printf("%s_%d\n", s, i)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go do("Hello", i, &wg)
		go do("World", i, &wg)
	}

	wg.Wait()
}
