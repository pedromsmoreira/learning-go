package functions

import "testing"

func TestSum(t *testing.T) {
	expected := 2
	total := sum(1, 1)

	if total != expected {
		t.Errorf("Sum was incorrect, got: %d, expected: %d", total, expected)
	}
}

func TestSum3(t *testing.T) {
	expected := 3
	total := sum3(1, 1, 1)

	if total != expected {
		t.Errorf("Sum3 was incorrect, got: %d, expected: %d", total, expected)
	}
}
