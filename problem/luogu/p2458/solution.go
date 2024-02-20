package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func minCost(root int, graph [][]int, costs []int) int {
	type Value struct{ Self, Parent, Children int }

	var dfs func(i int) Value
	dfs = func(i int) Value {
		self, parent, minExtra := 0, 0, math.MaxInt32
		for _, j := range graph[i] {
			v := dfs(j)

			self += min(v.Self, v.Parent)
			parent += min(v.Self, v.Children)
			minExtra = min(minExtra, v.Self-v.Children)
		}
		return Value{
			Self:     self + costs[i],
			Parent:   parent,
			Children: parent + max(0, minExtra),
		}
	}
	v := dfs(root)
	return min(v.Self, v.Children)
}

func bufferIO(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, m, w int
	fmt.Fscan(in, &n)
	cost := make([]int, n)
	g := make([][]int, n)
	root := -1
	for ; n > 0; n-- {
		fmt.Fscan(in, &v)
		v--
		if root == -1 {
			root = v
		}

		fmt.Fscan(in, &cost[v], &m)
		for ; m > 0; m-- {
			fmt.Fscan(in, &w)
			w--
			g[v] = append(g[v], w)
		}
	}
	fmt.Fprint(out, minCost(root, g, cost))
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
