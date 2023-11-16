package main

import (
	"sort"
)

// 排序 + 二分搜索
// 时间复杂度：O(nlogn + mlogn), n=len(flowers), m=len(people)
// 空间复杂度：O(n)
func fullBloomFlowers(flowers [][]int, people []int) []int {
	var start, end = make([]int, 0, len(flowers)), make([]int, 0, len(flowers))
	for _, flower := range flowers {
		start = append(start, flower[0])
		end = append(end, flower[1])
	}
	sort.Ints(start)
	sort.Ints(end)

	var ans = make([]int, len(people))
	for i, person := range people {
		x := sort.SearchInts(start, person+1) // 绽放过的数量
		y := sort.SearchInts(end, person)     // 已凋零的数量
		ans[i] = x - y
	}

	return ans
}

// 差分数组
// 时间复杂度：O(mlogm+nlogn), n=len(flowers), m=len(people)
// 空间复杂度：O(m+n)
func fullBloomFlowers2(flowers [][]int, people []int) []int {
	var diff = make(map[int]int)
	for _, flower := range flowers {
		left, right := flower[0], flower[1]
		diff[left]++
		diff[right+1]--
	}

	var n = len(diff)
	times := make([]int, 0, n)
	for t := range diff {
		times = append(times, t)
	}
	sort.Ints(times)

	idx := make([]int, len(people))
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool { return people[idx[i]] < people[idx[j]] })
	var ans = make([]int, len(people))

	j, sum := 0, 0
	for _, i := range idx {
		for ; j < n && times[j] <= people[i]; j++ {
			sum += diff[times[j]] // 累加不超过 people[i] 的差分值
		}
		ans[i] = sum // 从而得到这个时刻花的数量
	}
	return ans
}
