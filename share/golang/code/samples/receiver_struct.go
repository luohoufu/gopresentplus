package main

import "fmt"

type Me struct {
	Tonny   bool
	Luo bool
}

func (me Me) String() string {
	if me.Tonny == true && me.Luo == true {
		return "Tonny Luo"
	}
	return "Other Person"
}

func main() {
	m := Me{Tonny: true, Luo: true}
	fmt.Println(m)
}
