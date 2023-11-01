package main

import (
	"fmt"
)

// Person struct
 type Person struct {// Add your solution here!}

// Greet method
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func main() {
	person := Person{Name: "John Doe", Age: 30}
	person.Greet()
}