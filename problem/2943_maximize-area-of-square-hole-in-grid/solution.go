package main

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) (ans int) {
	x := min(maxN(hBars), maxN2(vBars))
	return x * x
}

func maxN(nums []int) int {
	var ans = 2
	sort.Ints(nums)

	var j = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			ans = max(ans, i-j+2)
		} else {
			j = i
		}
	}
	return ans
}

// 分组循环
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxN2(nums []int) (ans int) {
	sort.Ints(nums)
	var n = len(nums)
	for i := 0; i < n; {
		j := i
		for i++; i < n; i++ {
			if nums[i] != nums[i-1]+1 {
				break
			}
		}

		ans = max(ans, i-j+1)
	}
	return ans
}
