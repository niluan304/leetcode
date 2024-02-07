package main

import "sort"

func hIndex(citations []int) int {
	var n = len(citations)
	return n - sort.Search(n, func(i int) bool {
		return n-i <= citations[i]
	})
}
