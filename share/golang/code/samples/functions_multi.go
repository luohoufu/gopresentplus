// +build OMIT

package main

import "fmt"

func FullName() (string, string) {
	return "Adam", "Schaub"
}

func FullNameSuffix() (suffix string, first string, last string) {
	suffix = "Ser"
	first = "Adam"
	last = "Schaub"
	return
}

func main() {
	first, last := FullName()
	_, lastOnly := FullName()

	if last == lastOnly {
		fmt.Printf("The %s's family\n", first)
	}
	
	fmt.Println(FullNameSuffix())
}
