package main

import "sort"

func hIndex(citations []int) int {
	return sort.Search(len(citations)+1, func(i int) bool {
		var cnt = 0
		for _, citation := range citations {
			if citation >= i {
				cnt++
			}
		}
		return cnt < i
	}) - 1
}

func hIndex2(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	var count = 0
	for _, citation := range citations {
		if citation <= count {
			return count
		}
		count++
	}
	return count
}
