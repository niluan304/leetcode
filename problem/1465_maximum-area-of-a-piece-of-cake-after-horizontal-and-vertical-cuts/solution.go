package main

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	const mod = 1e9 + 7
	var (
		x = maxSize(h, horizontalCuts)
		y = maxSize(w, verticalCuts)
	)
	return (x * y) % mod // 1e9 * 1e9 < math.MaxInt64, 不会溢出
}

func maxSize(size int, cuts []int) int {
	sort.Ints(cuts)
	anx := max(cuts[0], size-cuts[len(cuts)-1])
	for i := 1; i < len(cuts); i++ {
		anx = max(anx, cuts[i]-cuts[i-1])
	}
	return anx
}
