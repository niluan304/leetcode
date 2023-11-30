package main

import (
	"math"
)

const mod = 1e9 + 7

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时
func sumSubarrayMins(arr []int) int {
	var ans = 0
	for i, _ := range arr {
		v := math.MaxInt
		for j := i; j >= 0; j-- {
			v = min(v, arr[j])
			ans = (ans + v) % mod
		}
	}
	return ans
}

// 单调栈 + 动态规划
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func sumSubarrayMins2(arr []int) int {
	arr = append([]int{0}, arr...) // 1 <= arr[i] <= 3 * 10^4

	var n = len(arr)
	var dp = make([]int, n)
	var stack []int
	var ans = 0
	for i, num := range arr {
		for m := len(stack); m > 0; m-- {
			j := stack[m-1]
			if num > arr[j] {
				// 处理后 arr[0] 一定为最小值
				dp[i] = dp[j] + num*(i-j)%mod
				ans = (ans + dp[i]) % mod
				break
			}
			stack = stack[:m-1]
		}
		stack = append(stack, i)
	}
	return ans
}
