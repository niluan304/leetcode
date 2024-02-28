package main

const mx = 1e5

// NotPrimes 质数 = false，非质数 = true
var NotPrimes = notPrimes(mx) // 预处理质数

func notPrimes(n int) (notPrimes []bool) {
	notPrimes = make([]bool, max(n+1, 2))
	for i := 2; i <= n; i++ {
		if !notPrimes[i] {
			for j := i * i; j <= n; j += i {
				notPrimes[j] = true
			}
		}
	}
	notPrimes[0], notPrimes[1] = true, true // 0 和 1 既不是质数，也不是合数
	return notPrimes
}

// 树形 dp + 分类讨论
// - 时间复杂度：O(n)。
// - 空间复杂度：O(n)。
func countPaths(n int, edges [][]int) int64 {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type Pair struct{ prime, notPrime int }

	ans := 0

	// dfs(i, fa) 表示当前遍历到第 i 个节点，父节点为 fa
	var dfs func(i int, fa int) (prime int, notPrime int)
	dfs = func(i int, fa int) (prime int, notPrime int) {
		pairs := make([]Pair, 0, len(graph[i]))
		for _, j := range graph[i] { // 每个节点只会访问一次
			if j == fa {
				continue
			}
			p, np := dfs(j, i)
			pairs = append(pairs, Pair{prime: p, notPrime: np})
			prime += p
			notPrime += np
		}

		// 分类讨论
		if NotPrimes[i] {
			// 以非质数的根节点为拐点，新增的合法路径的数量
			for _, pair := range pairs {
				ans += pair.prime * (notPrime - pair.notPrime + 1)
			}
			return prime, notPrime + 1
		} else {
			// 以质数的根节点为拐点，新增的合法路径的数量
			np := notPrime
			for _, pair := range pairs {
				np -= pair.notPrime // 避免重复统计 a->b, b->a
				ans += pair.notPrime * (np + 1)
			}
			return notPrime + 1, 0 // 根节点为质数时的「1 质数路径数」 = 自身 + 子树非质数路径数之和
		}
	}
	dfs(1, -1)
	return int64(ans)
}

// todo 并查集
func countPaths2(n int, edges [][]int) int64 {
	panic("unfinished")
}
