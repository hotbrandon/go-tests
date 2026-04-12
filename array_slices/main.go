package main

func SumArray(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}
	return sum
}
