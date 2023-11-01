package main

import (
	"fmt"
)

// Person struct
 type Person struct {
	Name string
	Age  int	 
 }

// Greet method
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func main() {
	person := Person{Name: "John Doe", Age: 30}
	person.Greet()
}
