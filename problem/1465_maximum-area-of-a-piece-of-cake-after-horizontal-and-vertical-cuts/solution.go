package main

import "sort"

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
	anx := _max(cuts[0], size-cuts[len(cuts)-1])
	for i := 1; i < len(cuts); i++ {
		anx = _max(anx, cuts[i]-cuts[i-1])
	}
	return anx
}
