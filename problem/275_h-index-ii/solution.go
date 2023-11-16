package main

import "sort"

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hIndex(citations []int) int {
	var n = len(citations)
	return n - sort.Search(n, func(i int) bool {
		return n-i <= citations[i]
	})
}
