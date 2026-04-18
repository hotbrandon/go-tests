package main

func SumArray(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var result []int
	for _, arr := range arrays {
		result = append(result, SumArray(arr))
	}

	return result
}

func SumAllTails(arrays ...[]int) []int {
	var result []int
	for _, arr := range arrays {
		if len(arr) < 2 {
			continue
		}
		result = append(result, SumArray(arr[1:]))
	}

	return result
}
