package main

import (
	"math"
)

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

func _sum(x []int) int {
	sum := 0
	for i, _ := range x {
		sum += x[i]
	}
	return sum
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^3)
func mergeStones(stones []int, k int) int {
	var n = len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	var preSum = make([]int, n+1)
	for i, stone := range stones {
		preSum[i+1] = preSum[i] + stone // 前缀和
	}

	var cache = make([][][]*int, n)
	for i, _ := range cache {
		cache[i] = make([][]*int, n)
		for j, _ := range cache[i] {
			cache[i][j] = make([]*int, k+1)
		}
	}

	var dfs func(i, j, p int) int
	dfs = func(i, j, p int) int {
		ptr := &cache[i][j][p]
		if *ptr != nil {
			return **ptr
		}

		if p == 1 {
			if i == j { // 只有一堆石头，无需合并
				return 0
			}
			return dfs(i, j, k) + preSum[j+1] - preSum[i]
		}

		cost := math.MaxInt32
		for m := i; m < j; m += k - 1 { // 枚举哪些石头堆合并成第一堆
			cost = _min(cost, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}

		*ptr = &cost
		return cost
	}
	return dfs(0, n-1, 1)
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^3)
// 空间复杂度：O(n^3)
func mergeStones2(stones []int, k int) int {
	var n = len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	var preSum = make([]int, n+1)
	for i, stone := range stones {
		preSum[i+1] = preSum[i] + stone // 前缀和
	}

	var cache = make([][][]*int, n)
	for i, _ := range cache {
		cache[i] = make([][]*int, n)
		for j, _ := range cache[i] {
			cache[i][j] = make([]*int, k+1)
		}
		cache[i][i][1] = new(int)
	}

	var dfs func(i, j, p int) int
	dfs = func(i, j, p int) int {
		if p == 1 {
			if i == j {
				return 0
			} else {
				return dfs(i, j, k) + preSum[j+1] - preSum[i]
			}
		}

		ptr := &cache[i][j][p]
		if *ptr != nil {
			return **ptr
		}
		cost := math.MaxInt32
		for x := i; x < j; x += k - 1 {
			cost = _min(cost, dfs(i, x, 1)+dfs(x+1, j, p-1))
		}

		*ptr = &cost
		return cost
	}
	return dfs(0, n-1, 1)
}

func mergeStones3(stones []int, k int) int {
	var n = len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	var dfs func(stones []int) int
	dfs = func(stones []int) int {
		if n == k {
			return _sum(stones)
		}

		cost := math.MaxInt32
		for x := 0; x <= n-k; x++ {
			var nums = make([]int, x)
			copy(nums, stones[:x])

			v := _sum(stones[x : x+k])
			nums = append(nums, v)
			nums = append(nums, stones[x+k:]...)

			cost = _min(
				cost,
				dfs(nums)+v,
			)
		}

		return cost
	}

	return dfs(stones)
}
