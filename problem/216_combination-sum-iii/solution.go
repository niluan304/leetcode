package main

func combinationSum3(k int, n int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if len(path) == k {
			if sum == n {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}

		for j := i; j <= 9; j++ {
			// 减枝
			if sum+j > n {
				return
			}

			path = append(path, j)
			dfs(j+1, sum+j)
			path = path[:len(path)-1]
		}
	}

	dfs(1, 0)

	return ans
}
