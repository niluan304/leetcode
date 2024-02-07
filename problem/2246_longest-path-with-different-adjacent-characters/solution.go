package main

func longestPath(parent []int, s string) int {
	var children = make([][]int, len(s))
	for i := 1; i < len(s); i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	var ans = 0
	var dfs func(p int) int
	dfs = func(p int) int {
		var x1, x2 = 0, 0
		for _, child := range children[p] {
			y := dfs(child)
			if s[p] == s[child] {
				continue
			}

			if y > x1 {
				x1, x2 = y, x1
			} else if y > x2 {
				x2 = y
			}
		}
		ans = max(ans, x1+x2+1)
		return x1 + 1
	}

	dfs(0)
	return ans
}
