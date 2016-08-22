package main

import (
	"fmt"
	"mypackage"
)

func main() {
	var mt mypackage.MyType = 2
	fmt.Println(mt)

	var st mypackage.secretType // can't do it!
}
