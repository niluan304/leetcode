package main

import (
	"bufio"
	"fmt"
	"os"
)

func mostSavedTree(n int, edges [][]int) int {
	var graph = make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type Value struct {
		Saved   int
		Unsaved int
	}

	var dfs func(i int, fa int) Value
	dfs = func(i int, fa int) (v Value) {
		var res []Value
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			res = append(res, dfs(j, i))
		}

		switch len(res) {
		case 0:
			return Value{
				Saved:   1,
				Unsaved: 0,
			}
		case 1:
			x := res[0]
			return Value{
				Saved:   x.Saved + 1,
				Unsaved: x.Saved - 1,
			}
		case 2:
			x, y := res[0], res[1]
			return Value{
				Saved:   x.Saved + y.Saved + 1,
				Unsaved: max(x.Saved+y.Unsaved, x.Unsaved+y.Saved) - 1,
			}
		}
		return
	}

	v := dfs(0, -1)
	return v.Unsaved
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func print(a ...interface{})            { fmt.Println(a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f+"\n", a...) }

func main() {
	defer writer.Flush()

	var testcases int
	scanf("%d", &testcases)
	for i := 0; i < testcases; i++ {
		scanf("%s")
		var n int
		scanf("%d", &n)
		edges := make([][]int, n-1)
		for j := 0; j < n-1; j++ {
			edges[j] = make([]int, 2)
			scanf("%d %d", &edges[j][0], &edges[j][1])
		}
		print(mostSavedTree(n, edges))
	}
}
