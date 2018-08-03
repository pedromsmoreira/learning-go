package recursion_test

import (
	"testing"

	"github.com/pedromsmoreira/learning-go/src/recursion"
)

func TestFact(t *testing.T) {
	exp := 6

	total := recursion.Fact(3)

	if total != exp {
		t.Errorf("Fact failed. Got: %d, Expected: %d", total, exp)
	}
}
