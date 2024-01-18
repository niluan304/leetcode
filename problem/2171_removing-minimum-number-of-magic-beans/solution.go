package main

import (
	"math"
	"sort"
)

// 排序 + 贪心
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var ans = math.MaxInt

	var n = len(beans)
	var preSum = make([]int, n+1)

	for i, bean := range beans {
		preSum[i+1] = preSum[i] + bean
	}

	for i, bean := range beans {
		cost := preSum[n] - preSum[i]
		ans = min(ans, cost-bean*(n-i)+preSum[i])
	}

	return int64(ans)
}

// 排序 + 贪心
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
//
// minimumRemoval 的空间优化版本
func minimumRemoval2(beans []int) int64 {
	sort.Ints(beans)

	n := len(beans)
	sum, most := 0, 0
	for i, bean := range beans {
		sum += bean
		most = max(most, bean*(n-i))
	}

	return int64(sum - most)
}
