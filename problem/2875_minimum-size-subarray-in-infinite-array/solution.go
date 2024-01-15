package main

import (
	"cmp"
	"math"
	"sort"
)

// 前缀和 + 二分搜索
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func minSizeSubarray(nums []int, target int) int {
	var n = len(nums)
	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num // num >= 1，因此 preSum 是递增的
	}

	// target = preSum[j] - preSum[i] + y * preSum[n]
	var ans = math.MaxInt32
	for i := 0; i <= n; i++ {
		sum := target + preSum[i]
		y := sum / preSum[n]
		x := sum - preSum[n]*y

		j := sort.SearchInts(preSum, x)
		if j <= n && preSum[j] == x {
			ans = min(ans, (n-i)+j+(y-1)*n)
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 滑动窗口
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func minSizeSubarray2(nums []int, target int) int {
	var n, sum = len(nums), 0
	var ans = math.MaxInt32
	var total = Sum(nums)

	y := target / total
	target = target - y*total
	// target = nums[i] +...+ nums[j] + y*Sum(nums), i ∈ [0, n-1], j ∈ [i+1, 2n-2]

	var left = 0
	for right := 0; right < n*2; right++ {
		sum += nums[right%n]
		for sum > target {
			sum -= nums[left%n]
			left++
		}

		if sum == target {
			ans = min(ans, right-left+1)
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans + y*n
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}
