package variadicfunctions_test

import (
	"testing"

	variadicfunctions "github.com/pedromsmoreira/learning-go/src/variadic-functions"
)

// TODO: Use test tables
func TestSumVaridic3Ints(t *testing.T) {
	expected := 3

	total := variadicfunctions.VariadicSum(1, 1, 1)

	if total != expected {
		t.Errorf("Sum failed!. Got: %d, Expected: %d", total, expected)
	}
}

// If you already have multiple args in a slice,
// apply them to a variadic function using func(slice...) like this.
func TestSumVaridicIntSlice(t *testing.T) {
	expected := 3

	s := []int{1, 1, 1}
	total := variadicfunctions.VariadicSum(s...)

	if total != expected {
		t.Errorf("Sum failed!. Got: %d, Expected: %d", total, expected)
	}
}
