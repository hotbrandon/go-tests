package main

import (
	"slices"
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
	t.Run("testing sum of two arrays", func(t *testing.T) {
		arr1 := []int{4, 5, 6}
		arr2 := []int{7, 8, 9}

		expected := []int{15, 24}
		got := SumAll(arr1, arr2)

		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

}

func TestSumAllTails(t *testing.T) {
	t.Run("testing sum of slices", func(t *testing.T) {
		arr1 := []int{4, 5, 6}
		arr2 := []int{7, 8, 9}

		expected := []int{11, 17}
		got := SumAllTails(arr1, arr2)

		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("testing sum of empty slices", func(t *testing.T) {
		arr1 := []int{}
		arr2 := []int{7, 8, 9}

		expected := []int{17}
		got := SumAllTails(arr1, arr2)

		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

}
