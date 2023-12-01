package main

import "math"

// 单调栈 + 前缀最小值
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func find132pattern(nums []int) bool {
	var stack []int // 递减单调栈
	var minPrefix = make([]int, len(nums)+1)
	minPrefix[0] = math.MaxInt
	for i, num := range nums {
		minPrefix[i+1] = min(minPrefix[i], num)

		m := len(stack)
		for m > 0 && num >= nums[stack[m-1]] {
			m--
		}
		if m > 0 && num > minPrefix[stack[m-1]] {
			return true
		}
		stack = append(stack[:m], i)
	}
	return false
}
