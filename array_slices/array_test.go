package main

import (
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("testing 5 elements in array", func(t *testing.T) {
		arr := [...]int{1, 2, 3, 4, 5}
		got := SumArray(arr[:])
		wanted := 15

		if got != wanted {
			t.Errorf("want %d, got %d, %v", wanted, got, arr)
		}
	})

}

func TestSumAll(t *testing.T) {
	arr1 := []int{4, 5, 6}
	arr2 := []int{7, 8, 9}

	expected := 39
	got := SumAll(arr1, arr2)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
