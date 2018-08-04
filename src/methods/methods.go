package main

import (
	"fmt"
)

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return r.width*2 + r.height*2
}

func main() {
	r := rect{height: 2, width: 2}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("Area: ", r.area())
	fmt.Println("Perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	rp := &r

	fmt.Println("Area: ", rp.area())
	fmt.Println("Perim: ", rp.perim())

}
