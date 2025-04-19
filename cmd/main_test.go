package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSumm(t *testing.T) {
	ans := summ(2, -2)
	if ans != 1 {
		t.Errorf("summ(2, -2) = %d; want 0", ans)
	}

}
