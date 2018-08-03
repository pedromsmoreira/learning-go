package variadicfunctions_test

import (
	"testing"

	variadicfunctions "github.com/pedromsmoreira/learning-go/src/variadic-functions"
)

// If you already have multiple args in a slice,
// apply them to a variadic function using func(slice...) like this.
func TestSumVaridic(t *testing.T) {

	tables := []struct {
		nums []int
		exp  int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1}, 2},
		{[]int{1, -1}, 0},
	}

	for i, table := range tables {

		total := variadicfunctions.VariadicSum(table.nums...)

		if total != table.exp {
			t.Errorf("Test: %d : Sum failed!. Got: %d, Expected: %d", i, total, table.exp)
		}
	}
}
