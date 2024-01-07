package main

import (
	"sort"
)

// 暴力穷举
// 时间复杂度: O(n)
// 空间复杂度: O(n)
// Deprecated: 超时
func maximizeWin(prizePositions []int, k int) int {
	var n = len(prizePositions)
	var prizes = make([]int, n)

	for i := 0; i < n; i++ {
		j := sort.SearchInts(prizePositions, prizePositions[i]+k+1) - 1
		prizes[i] = j - i + 1
	}

	// 通过暴力解法，发现了答案是：prize[i] + 后缀最大值
	var ans = 0
	for i, prize := range prizes {
		j := sort.SearchInts(prizePositions, prizePositions[i]+k+1)
		for ; j < n; j++ {
			ans = max(ans, prize+prizes[j])
		}
	}
	return min(ans, n)
}

// 前后缀分解
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
func maximizeWin2(prizePositions []int, k int) int {
	var n = len(prizePositions)
	var prizes, suffix = make([]int, n), make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		// 设 x = prizePositions[i]+k, 第一个大于等于 x+1 的元素的前一个位置, 就是 <= x 的位置

		j := sort.SearchInts(prizePositions, prizePositions[i]+k+1) - 1
		prizes[i] = j - i + 1
		suffix[i] = max(suffix[i+1], prizes[i])
	}

	var ans = 0
	for i, prize := range prizes {
		j := i + prizes[i]
		ans = max(ans, prize+suffix[j])
	}
	return ans
}
