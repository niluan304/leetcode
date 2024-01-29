package main

import (
	"math"
	"slices"
)

// 状态机 dp
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxScore(nums []int, x int) int64 {
	var n = len(nums)
	type Pair struct {
		Odd  int // 奇数
		Even int // 偶数
	}

	var dp = make([]Pair, n+1)
	dp[0] = Pair{
		Odd:  math.MinInt32,
		Even: math.MinInt32,
	}
	if num := nums[0]; num%2 == 0 {
		dp[0].Even = 0
	} else {
		dp[0].Odd = 0
	}

	for i, num := range nums {
		if num%2 == 0 { // 偶数
			dp[i+1] = Pair{
				Odd:  dp[i].Odd,
				Even: max(dp[i].Even, dp[i].Odd-x) + num,
			}
		} else {
			dp[i+1] = Pair{
				Odd:  max(dp[i].Odd, dp[i].Even-x) + num,
				Even: dp[i].Even,
			}
		}
	}

	ans := dp[n]
	return int64(max(ans.Odd, ans.Even))
}

// 状态机 dp
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxScore2(nums []int, x int) int64 {
	odd, even := math.MinInt32, math.MinInt32
	if nums[0]%2 == 0 {
		even = 0
	} else {
		odd = 0
	}
	// 也可以直接初始化，然后数组从 1 开始遍历

	for _, num := range nums {
		if num%2 == 0 { // 偶数
			even = max(even, odd-x) + num
		} else {
			odd = max(odd, even-x) + num
		}
	}
	return int64(max(odd, even))
}

// 状态机 dp
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxScore3(nums []int, x int) int64 {
	odd, even := nums[0], nums[0]
	if nums[0]%2 == 0 {
		odd -= x
	} else {
		even -= x
	}

	// 初始化完毕，数组需要从 nums[1:] 开始遍历
	for _, num := range nums[1:] {
		if num%2 == 0 { // 偶数
			even = max(even, odd-x) + num
		} else {
			odd = max(odd, even-x) + num
		}
	}
	return int64(max(odd, even))
}

// 状态机 dp
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxScore4(nums []int, x int) int64 {
	var dp = [2]int{-x, -x}
	dp[nums[0]%2] = nums[0]

	// 初始化完毕，数组需要从 nums[1:] 开始遍历
	for _, num := range nums[1:] {
		dp[num%2] = max(
			dp[num%2],       // 奇偶性相同
			dp[(num+1)%2]-x, // 奇偶性不同
		) + num
	}
	return int64(slices.Max(dp[:]))
}
