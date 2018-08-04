// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.
package main

import (
	"fmt"
)

// This person struct type has name and age fields
type person struct {
	name string
	age  int
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 20})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{"Ann", 21})

	// Access struct fields with a dot.
	p := person{name: "Sean", age: 55}
	fmt.Printf("Name: %s\n", p.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	pp := &p
	fmt.Println("Age:", pp.age)

	// Structs are mutable.
	pp.age = 51
	fmt.Println("New Age:", pp.age)
}
