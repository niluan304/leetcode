package main

import (
	"math"
	"slices"
	"sort"
)

// 暴力穷举
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 3
//
// - 时间复杂度：$\mathcal{O}(n^3)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func minimumOperations(nums []int) int {
	var n = len(nums)
	var ans = math.MaxInt32 // math.MinInt32
	for i := -1; i < n; i++ {
		for j := i; j < n; j++ {
			count := 0
			for k, num := range nums {
				if k <= i {
					if num != 1 {
						count++
					}
				} else if i < k && k <= j {
					if num != 2 {
						count++
					}
				} else if j < k {
					if num != 3 {
						count++
					}
				}
			}

			ans = min(ans, count)
		}
	}
	return ans
}

// 前缀和
// 使用前缀和之差得到操作区间 [i:j] 得到 1,2,3 的代价
//
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumOperations2(nums []int) int {
	var n = len(nums)
	var ans = math.MaxInt32

	var costs = make([][3]int, n+1)
	for i, num := range nums {
		costs[i+1] = costs[i]
		costs[i+1][num%3]++
		costs[i+1][(num+1)%3]++
	}

	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			cost := costs[i][0] + // 区间[0 : i] 变 1 的代价
				(costs[j][1] - costs[i][1]) + // 区间 [i+1: j] 变 2 的代价
				(costs[n][2] - costs[j][2]) // 区间 [j+1: n] 变 3 的代价
			ans = min(ans, cost)
		}
	}

	return ans
}

// LIS
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumOperations3(nums []int) int {
	var n = len(nums)

	var dp = make([]int, n)
	for i, num := range nums {
		length := 0
		for j := 0; j < i; j++ {
			if nums[j] > num {
				continue
			}
			length = max(length, dp[j])
		}
		dp[i] = length + 1
	}
	return n - slices.Max(dp)
}

// LIS
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumOperations4(nums []int) int {
	var greedy []int
	for _, num := range nums {
		j := sort.SearchInts(greedy, num+1)
		if j < len(greedy) {
			greedy[j] = num
		} else {
			greedy = append(greedy, num)
		}
	}

	return len(nums) - len(greedy)
}

// 状态机 DP
// - 时间复杂度：$\mathcal{O}(n \cdot 6)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minimumOperations5(nums []int) int {
	var n = len(nums)
	var dp = make([][4]int, n+1)

	for i, num := range nums {
		for j := 1; j <= 3; j++ {
			dp[i+1][j] = dp[i][j]
			for k := 1; k <= j; k++ {
				dp[i+1][j] = min(dp[i+1][j], dp[i][k])
			}

			if j != num {
				dp[i+1][j]++
			}
		}
	}

	return slices.Min(dp[n][1:])
}
