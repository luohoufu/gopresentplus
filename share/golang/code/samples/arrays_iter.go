package main

import "fmt"

func main() {
	arr := [5]string{"hen", "ducks", "squawking geese", "limerick oysters", "corpulent porpoises"}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %s\n", i+1, arr[i])
	}
	fmt.Println("")
	for ind, phrase := range arr {
		fmt.Printf("%d %s\n", ind+1, phrase)
	}
	fmt.Println("")
	for _, phrase := range arr {
		fmt.Println(phrase)
	}
}
