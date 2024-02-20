package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func withoutLeader(n int, rs []int, leaders [][]int) int {
	graph := make([][]int, n)
	sum := n * (n - 1) / 2
	for _, leader := range leaders {
		l, k := leader[0]-1, leader[1]-1
		graph[k] = append(graph[k], l)
		sum -= l // 缺失的数字即为 root
	}

	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		withI, withoutI := rs[i], 0
		for _, j := range graph[i] {
			withJ, withoutJ := dfs(j)
			withI += withoutJ
			withoutI += max(withJ, withoutJ)
		}
		return withI, withoutI
	}

	return max(dfs(sum))
}

func bufferIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscanln(in, &n)

	rs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &rs[i])
	}
	leaders := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		leaders[i] = make([]int, 2)
		fmt.Fscanln(in, &leaders[i][0], &leaders[i][1])
	}
	fmt.Fprint(out, withoutLeader(n, rs, leaders))
}

func main() {
	bufferIO(os.Stdin, os.Stdout)
}

func max(x, y int, z ...int) int {
	if x < y {
		x = y
	}
	for _, v := range z {
		if x < v {
			x = v
		}
	}
	return x
}

func min(x, y int, z ...int) int {
	if x > y {
		x = y
	}
	for _, v := range z {
		if x > v {
			x = v
		}
	}
	return x
}
