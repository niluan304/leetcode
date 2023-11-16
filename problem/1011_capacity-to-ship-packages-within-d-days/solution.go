package main

import (
	"sort"
)

// 二分查找
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func shipWithinDays(weights []int, days int) int {
	const N = 500 * 5 * 1e4
	return sort.Search(N, func(w int) bool {
		var d, sum = 1, 0
		for _, weight := range weights {
			if weight > w {
				return false
			}

			sum += weight
			if sum > w {
				d++
				sum = weight
			}
		}

		return d <= days
	})
}
