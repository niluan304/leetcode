package main

import (
	"math"
	"sort"
)

// 方法一：暴力法
// 暴力法是最直观的方法。初始化子数组的最小长度为无穷大，
// 枚举数组 nums 中的每个下标作为子数组的开始下标，
// 对于每个开始下标 i，需要找到大于或等于 iii 的最小下标 jjj，
// 使得从 nums[i] 到 nums[j] 的元素和大于或等于 sss，
// 并更新子数组的最小长度（此时子数组的长度是 j−i+1）。
// 注意：使用 Python 语言实现方法一会超出时间限制。
func leetcode1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 方法二：前缀和 + 二分查找
// 方法一的时间复杂度是 O(n2)，因为在确定每个子数组的开始下标后，
// 找到长度最小的子数组需要 O(n)的时间。如果使用二分查找，则可以将时间优化到 O(logn)。
//
// 为了使用二分查找，需要额外创建一个数组 sums 用于存储数组 nums 的前缀和，
// 其中 sums[i] 表示从 nums[0] 到 nums[i−1] 的元素和。得到前缀和之后，
// 对于每个开始下标 i，可通过二分查找得到大于或等于 i 的最小下标 bound，
// 使得 sums[bound]−sums[i−1]≥s，并更新子数组的最小长度（此时子数组的长度是 bound−(i−1)。
//
// 因为这道题保证了数组中每个元素都为正，所以前缀和一定是递增的，这一点保证了二分的正确性。
// 如果题目没有说明数组中每个元素都为正，这里就不能使用二分来查找这个位置了。
func leetcode2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
