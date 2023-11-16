package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	return bySortFunc(intervals, func(intervals [][]int) {
		quick(intervals, 0, len(intervals)-1)
	})
}

func eraseOverlapIntervals2(intervals [][]int) int {
	return bySortFunc(intervals, func(intervals [][]int) {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	})
}

func bySortFunc(intervals [][]int, sortFunc func(intervals [][]int)) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sortFunc(intervals)

	var cnt, k int
	for i := 1; i < len(intervals); i++ {
		vi := intervals[i]
		vk := intervals[k]

		if vk[1] > vi[0] {
			cnt++
			if vk[1] >= vi[1] {
				k = i
			}
		} else {
			k = i
		}
	}

	return cnt
}

func quick(intervals [][]int, l, r int) {
	if l >= r {
		return
	}
	pivot := intervals[l]
	i := l
	j := r
	for i < j {
		for i < j {
			if intervals[j][0] < pivot[0] || (intervals[j][0] == pivot[0] && intervals[j][1] < pivot[1]) {
				intervals[i] = intervals[j]
				i++
				break
			} else {
				j--
			}
		}
		for i < j {
			if intervals[i][0] > pivot[0] || (intervals[i][0] == pivot[0] && intervals[i][1] > pivot[1]) {
				intervals[j] = intervals[i]
				j--
				break
			} else {
				i++
			}
		}
	}
	intervals[i] = pivot
	quick(intervals, l, i-1)
	quick(intervals, i+1, r)
}

func leetcode1(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return n - max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func leetcode2(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}

func leetcode3(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < right {
			res++
		} else {
			right = intervals[i][1]
		}
	}
	return res
}
