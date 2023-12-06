package main

import "math"

// 求出 $\textit{nums}$ 的前缀和 $s$ 后，我们可以写一个暴力算法，
// 枚举所有满足 $i>j$ 且 $s[i]-s[j]\ge k$ 的子数组 $[j,i)$，取其中最小的 $i-j$ 作为答案。
//
// 但这个暴力算法是 $O(n^2)$ 的，如何优化呢？
//
// 我们可以遍历 $s$，同时用某个合适的数据结构来维护遍历过的 $s[i]$，并**及时移除无用的 $s[i]$**。
//
// 优化一：
//
// ![862-1-2.png](https://pic.leetcode.cn/1666668814-ikkWBN-862-1-2.png)
//
// 优化二：
//
// ![862-2-3.png](https://pic.leetcode.cn/1666669250-KypIVI-862-2-3.png)
//
// 做完这两个优化后，再把 $s[i]$ 加到这个数据结构中。
//
// 由于优化二保证了数据结构中的 $s[i]$ 会形成一个递增的序列，因此优化一移除的是序列最左侧的若干元素，
// 优化二移除的是序列最右侧的若干元素。我们需要一个数据结构，它支持移除最左端的元素和最右端的元素，以及在最右端添加元素，故选用**双端队列**。
//
// > 注：由于双端队列的元素始终保持单调递增，因此这种数据结构也叫做**单调队列**。
func endlesscheng(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, x := range nums {
		prefixSum[i+1] = prefixSum[i] + x // 计算前缀和
	}
	ans := math.MaxInt
	var queue []int
	for i, s := range prefixSum {
		for len(queue) > 0 && s-prefixSum[queue[0]] >= k {
			ans = min(ans, i-queue[0])
			queue = queue[1:] // 优化一
		}
		for len(queue) > 0 && prefixSum[queue[len(queue)-1]] >= s {
			queue = queue[:len(queue)-1] // 优化二
		}
		queue = append(queue, i)
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
