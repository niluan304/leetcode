package main

import (
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	zeroIdx := sort.SearchInts(satisfaction, 0)

	var ans, sn = 0, 0
	for i, num := range satisfaction[zeroIdx:] {
		sn += num
		ans += num * (i + 1) // n <= 500, num <= 1000, 不会溢出
	}

	for i := zeroIdx - 1; i >= 0; i-- {
		sn += satisfaction[i]
		if sn <= 0 {
			break
		}
		ans += sn
	}

	return ans
}

// 排序 + 贪心
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func maxSatisfaction2(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	var sn = 0  // satisfaction 的前缀和
	var ans = 0 // f(0) = 0
	for _, num := range satisfaction {
		sn += num
		if sn <= 0 {
			break
		}
		ans += sn // f(k) = f(k-1) + s
	}
	return ans
}
