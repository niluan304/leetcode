package main

import (
	"math"
)

// 模拟
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		give, get := operation[0], operation[1]
		half := gem[give] / 2
		gem[give] -= half
		gem[get] += half
	}

	var minGem, maxGem = math.MaxInt, math.MinInt
	for _, g := range gem {
		minGem = _min(minGem, g)
		maxGem = _max(maxGem, g)
	}
	return maxGem - minGem
}

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
