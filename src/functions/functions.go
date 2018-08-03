package functions

// Sum two int numbers and return result
func Sum(a int, b int) int {
	return a + b
}

// Sum3 int numbers and return result
func Sum3(a, b, c int) int {
	return a + b + c
}

// MultipleReturn function to Sum and Subtract two int numbers and return sum and sub
func MultipleReturn(a int, b int) (sum int, sub int) {
	return a + b, a - b
}

/*func main() {

	res := sum(1, 2)
	fmt.Println("1 + 2 :", res)

	res = sum3(1, 2, 3)
	fmt.Println("1 + 2 + 3: ", res)

	p, m := multipleReturn(2, 1)
	fmt.Printf("2+1=%d AND 2-1=%d\n", p, m)
}
*/
