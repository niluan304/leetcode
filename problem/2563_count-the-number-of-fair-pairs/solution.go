package main

import "sort"

// ### 前置知识：二分查找
//
// 有关二分查找的写法，可以看我的 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频。
//
// **视频中详细介绍了如何把「小于」和「小于等于」转换成「大于等于」。**
//
// ### 思路
//
// 由于排序不会影响数对的个数，为了能够二分，可以先排序。
//
// 然后枚举 $\textit{nums}[j]$，二分查找符合要求的 $\textit{nums}[i]$ 的个数。
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(1)$。忽略排序的栈开销，仅用到若干额外变量。
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	var ans = 0
	for i, num := range nums {
		x := sort.SearchInts(nums[i+1:], lower-num)
		y := sort.SearchInts(nums[i+1:], upper-num+1) - 1
		ans += max(0, y-x+1)
	}
	return int64(ans)
}
