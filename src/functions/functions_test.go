package functions_test

import "testing"
import . "github.com/pedromsmoreira/learning-go/src/functions"

func TestSum(t *testing.T) {
	expected := 2
	total := Sum(1, 1)

	if total != expected {
		t.Errorf("Sum was incorrect, got: %d, expected: %d", total, expected)
	}
}

func TestSum3(t *testing.T) {
	expected := 3
	total := Sum3(1, 1, 1)

	if total != expected {
		t.Errorf("Sum3 was incorrect, got: %d, expected: %d", total, expected)
	}
}

func TestMultipleReturn(t *testing.T) {
	expectedSum, expectedSub := 2, 0

	totalSum, totalSub := MultipleReturn(1, 1)

	if totalSum != expectedSum {
		t.Errorf("totalSum is incorrect. Got: %d, Expected: %d", totalSum, expectedSum)
	}

	if totalSub != expectedSub {
		t.Errorf("totalSub is incorrect. Got: %d, Expected: %d", totalSub, expectedSub)
	}
}
