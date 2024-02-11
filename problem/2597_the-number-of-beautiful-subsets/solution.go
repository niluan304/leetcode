package main

import (
	"slices"
)

// # 方法一：回溯
//
// ### 前置知识：子集型回溯
//
// 见[【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。
//
// ### 思路
//
// 在枚举 [78. 子集](https://leetcode.cn/problems/subsets/) 的基础上加个判断。
//
// 在选择 $x=\textit{nums}[i]$ 的时候，如果之前选过 $x-k$ 或 $x+k$，则不能选，否则可以选。
//
// 代码实现时，可以用哈希表或者数组来记录选过的数，从而 $O(1)$ 判断 $x-k$ 和 $x+k$ 是否选过。
//
// ### 复杂度分析
//
// - 时间复杂度：$O(2^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(n)$。用哈希表实现是 $O(n)$。（数组需要一些额外空间，但是更快。）
func beautifulSubsets(nums []int, k int) int {
	var n = len(nums)
	slices.Sort(nums[:]) // 排序后只需要判断 nums[i] - k，不需要判断 nums[i] + k
	count := map[int]int{}
	ans := 0
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans++ // 合法子集
			return
		}
		dfs(i + 1) // 不选

		num := nums[i]
		if count[num-k] == 0 {
			count[num]++ // 选
			dfs(i + 1)
			count[num]-- // 恢复现场
		}
	}
	dfs(0)
	return ans - 1 // 去除空集
}

// ### 复杂度分析
//
// - 时间复杂度：$O(2^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(n)$。用哈希表实现是 $O(n)$。（数组需要一些额外空间，但是更快。）
func beautifulSubsets2(nums []int, k int) int {
	ans := 0
	n := len(nums)

	slices.Sort(nums[:]) // 排序后只需要判断 nums[i] - k，不需要判断 nums[i] + k
	cnt := make(map[int]int)
	var dfs func(i int)
	dfs = func(i int) { // 从 i 开始选
		ans++ // 合法子集
		if i == n {
			return
		}
		for j := i; j < n; j++ { // 枚举选哪个
			x := nums[j]
			if cnt[x-k] != 0 {
				continue
			}
			cnt[x]++ // 选
			dfs(j + 1)
			cnt[x]-- // 恢复现场
		}
	}

	dfs(0)
	return ans - 1 // -1，去掉空集
}
