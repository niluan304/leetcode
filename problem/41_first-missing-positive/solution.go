package main

import "math"

// 哈希表
//
// Think about how you would solve the problem in non-constant space.
// Can you apply that logic to the existing space?
//
// 类似解法：
// [LC128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/)
// [LC3020. 子集中元素的最大数量](https://leetcode.cn/problems/find-the-maximum-number-of-elements-in-subset/)
//
// - 时间复杂度：O(nlogn)。
// - 空间复杂度：O(n)。
// Deprecated: Memory out limit
func tip1(nums []int) int {
	exist := map[int]bool{}
	for _, num := range nums {
		if num > 0 {
			exist[num] = true
		}
	}
	for i := 1; ; i++ {
		if !exist[i] {
			return i
		}
	}
}

// 原地哈希表
//
// 借助原数组空间
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(1)。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for x := nums[i]; 1 <= x && x <= n && nums[x-1] != math.MaxInt; x = nums[i] {
			nums[i], nums[x-1] = nums[x-1], math.MaxInt
		}
	}

	for i, num := range nums {
		if num != math.MaxInt {
			return i + 1
		}
	}
	return n + 1
}
