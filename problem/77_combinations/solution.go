package main

func combine(n int, k int) [][]int {
	var ans [][]int
	var nodes []int

	//var dfs func()
	var dfs func(node int)
	dfs = func(node int) {
		if len(nodes) == k {
			v := append([]int(nil), nodes...)
			ans = append(ans, v)
			return
		}

		for i := node; i <= n; i++ {
			nodes = append(nodes, i)
			dfs(i + 1)
			nodes = nodes[:len(nodes)-1]
		}
	}

	dfs(1)
	return ans
}
