package main

import (
	"fmt"
	"sort"
)

// Go’s sort package implements sorting for builtins and user-defined types.
// We’ll look at sorting for builtins first.

func main() {
	// Sort methods are specific to the builtin type; here’s an example for strings.
	// Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
	strs := []string{"c", "a", "d", "a"}
	sort.Strings(strs)
	fmt.Println("Sorted: ", strs)

	ints := []int{5, 2, 7, 0, 1, 3}
	sort.Ints(ints)
	fmt.Println("Sorted: ", ints)

	areSorted := sort.IntsAreSorted(ints)
	fmt.Println("Are Sorted: ", areSorted)

}
