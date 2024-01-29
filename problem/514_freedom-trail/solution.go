package main

import (
	"math"
	"sort"
)

// dfs + 记忆化搜搜
// - 时间复杂度：$\mathcal{O}(n^2 \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func findRotateSteps(ring string, key string) int {

	var n, m = len(ring), len(key)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var cache = make([][]*int, n)
	for i, _ := range cache {
		cache[i] = make([]*int, m)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if j == m {
			return 0
		}

		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		res = math.MaxInt32
		for k := 0; k <= (n+1)/2; k++ { // 可以用哈希表预处理出索引
			x := (i + k) % n
			if ring[x] == key[j] {
				res = min(res, dfs(x, j+1)+k+1)
			}
		}
		for k := 1; k <= (n+1)/2; k++ {
			x := (i - k + n) % n
			if ring[x] == key[j] {
				res = min(res, dfs(x, j+1)+k+1)
			}
		}

		return res
	}

	ans = dfs(0, 0)
	return ans
}

// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n\log n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func findRotateSteps2(ring string, key string) int {

	var n, m = len(ring), len(key)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var cache = make([][]*int, n)
	for i, _ := range cache {
		cache[i] = make([]*int, m)
	}

	var index = make(map[rune][]int, n)
	for i, b := range ring {
		index[b] = append(index[b], i)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if j == m {
			return 0
		}

		ptr := &cache[i][j]
		if *ptr != nil {
			return **ptr
		}
		defer func() { *ptr = &res }()

		res = math.MaxInt32

		// 选左 or 选右
		idx := index[rune(key[j])]
		idxLen := len(idx)
		k := sort.SearchInts(idx, i)

		// 可以进一步预处理 left, right
		left := idx[(k-1+idxLen)%idxLen]
		right := idx[k%idxLen]

		return min(
			dfs(left, j+1)+(i+n-left)%n+1,
			dfs(right, j+1)+(right+n-i)%n+1,
		)
	}

	ans = dfs(0, 0)
	return ans
}
