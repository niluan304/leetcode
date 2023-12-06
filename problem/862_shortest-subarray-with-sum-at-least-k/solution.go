package main

import (
	"math"
	"sort"
)

// 单调队列 + 二分查找
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
// 求子数组的和，那么就可以转化为前缀和之差，从而等价于为寻找：
// i > j
// prefix[i] - prefix[j] >= k，且 i - j  的最小值
//
// 遍历前缀和数组，假设当前 i 满足条件，那么寻找最大的 j 以最小化 i - j
// 也就是寻找 最大的 j 使得 prefix[j] <= prefix[i] - k
func shortestSubarray(nums []int, k int) int {
	var n = len(nums)
	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	var queue = []int{0}
	var ans = math.MaxInt

	for i := 1; i <= n; i++ {
		m := len(queue)
		for m > 0 && preSum[i] <= preSum[queue[m-1]] {
			m--
		}
		queue = append(queue[:m], i)

		diff := preSum[i] - k
		// <= x 等价于 >= x + 1 左边的那个数
		j := sort.Search(len(queue), func(i int) bool {
			return preSum[queue[i]] > diff
		}) - 1
		if j >= 0 && preSum[queue[j]] <= diff {
			ans = min(ans, i-queue[j])
			queue = queue[j+1:]
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// 单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func shortestSubarray2(nums []int, k int) int {
	var preSum = make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	var queue []int
	var ans = math.MaxInt
	for i, s := range preSum {
		for len(queue) > 0 && s-preSum[queue[0]] >= k {
			ans = min(ans, i-queue[0])
			queue = queue[1:]
		}

		for m := len(queue); m > 0 && s <= preSum[queue[m-1]]; m-- {
			queue = queue[:m-1]
		}
		queue = append(queue, i)
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
