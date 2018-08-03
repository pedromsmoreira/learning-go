package recursion

// Fact returns the factorial product of the given number
func Fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * Fact(n-1)
}
