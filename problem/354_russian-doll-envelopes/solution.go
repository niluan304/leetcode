package main

import (
	"sort"
)

// 贪心 + 二分搜索
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func maxEnvelopes(envelopes [][]int) int {
	// 按宽度升序排列，如果宽度一样，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		wi, hi := envelopes[i][0], envelopes[i][1]
		wj, hj := envelopes[j][0], envelopes[j][1]
		if wi == wj {
			return hi > hj
		}
		return wi < wj
	})

	// 对高度数组寻找 LIS
	var heights = make([]int, len(envelopes))
	for i, envelope := range envelopes {
		heights[i] = envelope[1]
	}

	return lengthOfLIS(heights)
}

func lengthOfLIS(nums []int) int {
	var greed []int
	for i := 0; i < len(nums); i++ {
		j := sort.SearchInts(greed, nums[i])
		if j == len(greed) {
			greed = append(greed, nums[i])
		} else {
			greed[j] = nums[i]
		}
	}
	return len(greed)
}

// dp 动态规划，LIS
// 时间复杂度：O(n^2)，n <= 1e5 超时
// 空间复杂度：O(n)
func maxEnvelopes2(envelopes [][]int) int {
	// 按宽度升序排列，如果宽度一样，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		wi, hi := envelopes[i][0], envelopes[i][1]
		wj, hj := envelopes[j][0], envelopes[j][1]
		if wi == wj {
			return hi > hj
		}
		return wi < wj
	})

	var ans = 0
	var dp = make([]int, len(envelopes))
	for i, envelope := range envelopes {
		w, h := envelope[0], envelope[1]
		length := 0
		for j := 0; j < i; j++ {
			_w, _h := envelopes[j][0], envelopes[j][1]
			if w > _w && h > _h {
				length = _max(length, dp[j])
			}
		}

		dp[i] = length + 1
		ans = _max(ans, dp[i])
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
