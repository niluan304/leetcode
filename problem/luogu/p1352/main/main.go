package main

import (
	"fmt"
)

func withoutLeader(n int, rs []int, leaders [][2]int) int {
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
			withoutI += _max(withJ, withoutJ)
		}
		return withI, withoutI
	}

	return _max(dfs(sum))
}

func main() {
	var n int
	fmt.Scanln(&n)

	var rs = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&rs[i])
	}
	leaders := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scanln(&leaders[i][0], &leaders[i][1])
	}
	fmt.Println(withoutLeader(n, rs, leaders))
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}