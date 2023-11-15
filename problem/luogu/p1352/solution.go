package p1352

func withoutLeader(n int, rs []int, leaders [][]int) int {
	var graph = make([][]int, n)
	var sum = n * (n - 1) / 2
	for _, leader := range leaders {
		l, k := leader[0]-1, leader[1]-1
		graph[k] = append(graph[k], l)
		sum -= l // 缺失的数字即为 root
	}

	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		var withI, withoutI = rs[i], 0
		for _, j := range graph[i] {
			withJ, withoutJ := dfs(j)
			withI += withoutJ
			withoutI += max(withJ, withoutJ)
		}
		return withI, withoutI
	}

	return max(dfs(sum))
}
