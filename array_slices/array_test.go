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
