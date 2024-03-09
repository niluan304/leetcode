package main

import (
	"sort"

	. "github.com/niluan304/leetcode/container"
)

func pickGifts(gifts []int, k int) int64 {
	hp := NewMaxIntHeap(gifts...)

	for i := 0; i < k; i++ {
		hp.Update(0, mySqrt(hp.Peek()))
	}

	total := 0
	for _, gift := range hp.Values() {
		total += gift
	}
	return int64(total)
}

// [69. x 的平方根](https://leetcode.cn/problems/sqrtx/description/)
func mySqrt(x int) int {
	return sort.Search(x+1, func(i int) bool { return i*i > x }) - 1
}
