package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Person
	Company string
}

func main() {
	e := Employee{Person: Person{"Bruce Willis"}, Company: "Cisgoat"}

	fmt.Printf("e: [%T] %#v\n", e, e)

	fmt.Printf("Employee's name is %s\n", e.Name)
	fmt.Printf("Employee.Person's name is %s\n", e.Person.Name)
}
