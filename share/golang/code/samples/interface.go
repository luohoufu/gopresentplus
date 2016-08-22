// OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	wd := time.Now().Weekday()
	fmt.Printf("%s (%d)\n", wd, wd)
}
