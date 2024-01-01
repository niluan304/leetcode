package main

import (
	"sort"
)

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时
func getSubarrayBeauty(nums []int, k int, x int) []int {
	var n = len(nums)
	var ans = make([]int, n-k+1)
	var tmp = make([]int, k)
	for i := 0; i <= n-k; i++ {
		copy(tmp, nums[i:i+k])
		sort.Ints(tmp)
		ans[i] = min(0, tmp[x-1])
	}

	return ans
}

// ### 思路
//
// 滑动窗口。由于值域很小，所以借鉴**计数排序**，用一个 $\textit{cnt}$ 数组维护窗口内每个数的出现次数。然后遍历 $\textit{cnt}$ 去求第 $x$ 小的数。
//
// 什么是第 $x$ 小的数？
//
// 设它是 $\textit{num}$，那么 $<\textit{num}$ 的数有 $<x$ 个，$\le\textit{num}$ 的数有 $\ge x$ 个，就说明 $\textit{num}$ 是第 $x$ 小的数。
// ## 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=50$。
// - 空间复杂度：$\mathcal{O}(U)$。
func getSubarrayBeauty2(nums []int, k int, x int) []int {
	var n = len(nums)
	var ans = make([]int, n-k+1)

	const Base = 50
	var set [Base*2 + 1]int

	for i, num := range nums {
		set[num+Base]++

		j := i + 1 - k
		if j < 0 {
			continue
		}

		count := 0
		for v, cnt := range set[:Base] {
			count += cnt
			if count >= x {
				ans[j] = v - Base
				break
			}
		}
		set[nums[j]+Base]--
	}
	return ans
}
