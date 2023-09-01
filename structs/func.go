package structs

import "math"

func Sum(list []int) int {
	var sum = 0
	for _, n := range list {
		sum += n
	}

	return sum
}

func Max(list ...int) int {
	var ans = math.MinInt
	for _, n := range list {
		if ans < n {
			ans = n
		}
	}

	return ans
}

func Min(list ...int) int {
	var ans = math.MaxInt
	for _, n := range list {
		if ans > n {
			ans = n
		}
	}

	return ans
}
