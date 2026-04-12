package integers

import "testing"

func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("Expected: %d, Got: %d", expected, sum)
	}
}
