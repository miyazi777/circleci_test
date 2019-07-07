package calc

import "testing"

func TestIncrement(t *testing.T) {
	result := Increment(1)
	if result != 2 {
		t.Fatal("failed test")
	}
}
