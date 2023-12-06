package main

import (
	"math"
)

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时 (2 <= points.length <= 10^5)
func findMaxValueOfEquation(points [][]int, k int) int {
	var ans = math.MinInt
	for i, point := range points {
		x1, y1 := point[0], point[1]
		for _, p := range points[i+1:] {
			x2, y2 := p[0], p[1]
			if x2-x1 > k {
				break
			}
			ans = max(ans, y1+y2+x2-x1)
		}
	}

	return ans
}

// 单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func findMaxValueOfEquation2(points [][]int, k int) int {
	var ans = math.MinInt
	type Value struct {
		Index int
		Value int
	}
	var queue []Value
	for _, point := range points {
		x, y := point[0], point[1]
		for len(queue) > 0 && x-queue[0].Index > k {
			queue = queue[1:]
		}
		if len(queue) > 0 {
			ans = max(ans, queue[0].Value+x+y)
		}

		v := y - x
		for m := len(queue); m > 0 && queue[m-1].Value <= v; m-- {
			queue = queue[:m-1]
		}
		queue = append(queue, Value{Index: x, Value: v})
	}

	return ans
}
