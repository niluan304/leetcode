package main

import (
	"sort"
)

const mod int = 1e9 + 7

// 排序 + 动态规划
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)
func sumOfPower(nums []int) int {
	sort.Ints(nums)

	var dp, power = nums[0], ((nums[0] * nums[0] % mod) * nums[0]) % mod

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		dp = (num - nums[i-1] + 2*dp%mod) % mod
		power = (power + (num*num%mod)*dp) % mod
	}

	return power
}

// 排序 + 动态规划
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)
func sumOfPower2(nums []int) int {
	nums = append(nums, 0) // 哨兵
	sort.Ints(nums)
	var n = len(nums)
	var dp, power = make([]int, n+1), 0

	for i := 1; i < n; i++ {
		num := nums[i]
		dp[i] = (num - nums[i-1] + 2*dp[i-1]%mod) % mod
		power = (power + (num*num%mod)*dp[i]) % mod
	}

	return power
}
