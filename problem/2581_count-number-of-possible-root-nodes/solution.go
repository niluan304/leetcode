package main

import (
	. "github.com/niluan304/leetcode/copypasta/dp"
)

// 换根 dp
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type pair struct{ fa, i int }
	s := make(map[pair]int, len(guesses))
	for _, guess := range guesses { // guesses 转成哈希表
		s[pair{fa: guess[0], i: guess[1]}] = 1
	}

	size := make([]int, n)
	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		if s[pair{fa: fa, i: x}] == 1 { // 以 0 为根时，猜对了
			size[0]++
		}
		for _, y := range graph[x] {
			if y == fa {
				continue
			}
			dfs(y, x)
		}
	}
	dfs(0, -1)

	ans := 0
	var reroot func(x, fa int)
	reroot = func(x, fa int) {
		for _, y := range graph[x] {
			if y == fa {
				continue
			}
			origin := pair{fa: x, i: y}
			cur := pair{fa: y, i: x}
			size[y] = size[x] - s[origin] + s[cur]

			reroot(y, x)
		}
	}
	reroot(0, -1)

	for _, v := range size {
		if v >= k {
			ans++
		}
	}
	return ans
}

// 记忆化搜索
//
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(n)。
// Deprecated: 超时
func rootCount2(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type Pair struct{ fa, i int }
	pairs := make(map[Pair]int, len(guesses))
	for _, guess := range guesses {
		fa, i := guess[0], guess[1]
		pairs[Pair{fa: fa, i: i}] = 1
	}

	// 由于只会访问相邻的边，那 dfs 只有三种情况：dfs(i, fa), dfs(fa, i) 以及 dfs(i, -1)
	var dfs func(i, fa int) (count int)
	dfs = func(i, fa int) (count int) {
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			count += dfs(j, i)
		}
		return count + pairs[Pair{fa: fa, i: i}]
	}

	MemorySearch2(&dfs) // 类似 Python 的 @cache 装饰器

	ans := 0
	for i := 0; i < n; i++ {
		if dfs(i, -1) >= k {
			ans++
		}
	}
	return ans
}
