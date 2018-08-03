package variadicfunctions

import "fmt"

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.

// VariadicSum takes an arbitrary number of args
func VariadicSum(nums ...int) int {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}
