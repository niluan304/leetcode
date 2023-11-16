package main

import (
	"math"
)

// 暴力解法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
//
// 题目要求找到 -1 和 1 最多能夹多少个0
func captureForts(forts []int) int {
	var ans = 0
	for i, fort := range forts {
		if fort == 0 {
			continue
		}

		var path = 0
		for _, f := range forts[i+1:] {
			if f != 0 {
				if f+fort == 0 {
					ans = _max(ans, path)
				}
				break
			}
			path++
		}
	}

	return ans
}

func _max(list ...int) int {
	var ans = math.MinInt
	for _, n := range list {
		if ans < n {
			ans = n
		}
	}

	return ans
}

// 双指针, 滑动窗口
// 时间复杂度：O(n)
// 空间复杂度：O(1)
//
// 题目要求找到 -1 和 1 最多能夹多少个0
func captureForts2(forts []int) int {
	var ans = 0
	var left = 0
	for right := range forts {
		if forts[right] == 0 {
			continue
		}

		// 更新城堡数目
		// 求出 -1和1 之间 0 的个数
		if forts[right]+forts[left] == 0 {
			ans = _max(ans, right-left-1)
		}
		// 将 left 移动至非 0 的位置
		left = right
	}

	return ans
}
