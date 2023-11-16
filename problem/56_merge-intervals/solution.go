package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	return mergeBySort(intervals, func(intervals [][]int) {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	})
}

func merge2(intervals [][]int) [][]int {
	return mergeBySort(intervals, func(intervals [][]int) {
		quick(intervals, 0, len(intervals)-1)
	})
}

func mergeBySort(intervals [][]int, sortFunc func(intervals [][]int)) [][]int {
	n := len(intervals) - 1
	if n < 1 {
		return intervals
	}

	sortFunc(intervals)

	var ans = make([][]int, 0, n)
	ans = append(ans, intervals[0])

	for _, curr := range intervals[1:] {
		r2 := ans[len(ans)-1][1]
		if curr[0] > r2 {
			ans = append(ans, curr)
			continue
		}

		r1 := curr[1]
		if r1 > r2 {
			ans[len(ans)-1][1] = r1
		}
	}

	return ans
}

func leetcode2(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	quick(intervals, 0, len(intervals)-1)

	ret := make([][]int, 0)
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		base := ret[len(ret)-1]
		if base[0] <= intervals[i][0] && base[1] < intervals[i][0] {
			ret = append(ret, intervals[i])
			continue
		}
		if base[0] <= intervals[i][0] && base[1] >= intervals[i][1] {
			continue
		}
		if base[0] <= intervals[i][0] && base[1] >= intervals[i][0] {
			base[1] = intervals[i][1]
		}
	}
	return ret
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

func leetcode3(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res = make([][]int, 0, len(intervals))
	res = append(res, intervals[0])
	for _, interval := range intervals[1:] {
		var n = len(res) - 1
		if res[n][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[n][1] = max(res[n][1], interval[1])
		}
	}
	return res
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
