package graph

import (
	"reflect"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func TestFloyd(t *testing.T) {
	type Edge struct {
		from, to int
		weight   int
	}

	buildInit := func(edges []Edge) func(graph [][]int) {
		return func(graph [][]int) {
			for _, edge := range edges {
				from, to, w := edge.from, edge.to, edge.weight
				graph[from][to] = w
			}
		}
	}

	type args struct {
		n    int
		init func(graph [][]int)
	}
	ts := []struct {
		name     string
		args     args
		wantPath [][]int
	}{
		{
			name: "",
			args: args{
				n: 5,
				init: buildInit([]Edge{
					{from: 0, to: 1, weight: 2},
					{from: 0, to: 4, weight: 8},
					{from: 1, to: 2, weight: 3},
					{from: 1, to: 4, weight: 2},
					{from: 2, to: 3, weight: 1},
					{from: 3, to: 4, weight: 1},

					// 无向图，得双向
					{from: 1, to: 0, weight: 2},
					{from: 4, to: 0, weight: 8},
					{from: 2, to: 1, weight: 3},
					{from: 4, to: 1, weight: 2},
					{from: 3, to: 2, weight: 1},
					{from: 4, to: 3, weight: 1},
				}),
			},
			wantPath: [][]int{
				{0, 2, 5, 5, 4},
				{2, 0, 3, 3, 2},
				{5, 3, 0, 1, 2},
				{5, 3, 1, 0, 1},
				{4, 2, 2, 1, 0},
			},
		},
		// todo 增加有向图的用例
	}

	fs := []func(n int, init func(path [][]int)) (path [][]int){
		Floyd,
		FloydDFS,
	}

	for _, f := range fs {
		name := tests.FuncName(f)
		t.Run(name, func(t *testing.T) {
			for _, tt := range ts {
				t.Run(tt.name, func(t *testing.T) {
					gotPath := f(tt.args.n, tt.args.init)
					if !reflect.DeepEqual(gotPath, tt.wantPath) {
						t.Errorf("%v() gotDistance = %v, want %v", name, gotPath, tt.wantPath)
					}
				})
			}
		})
	}
}

func TestDijkstra(t *testing.T) {
	type Edge struct {
		from, to int
		weight   int
	}

	buildInit := func(edges []Edge) func(graph [][]DijkstraEdge) {
		return func(graph [][]DijkstraEdge) {
			for _, edge := range edges {
				from, to, w := edge.from, edge.to, edge.weight
				graph[from] = append(graph[from], DijkstraEdge{To: to, Weight: w})
			}
		}
	}

	type args struct {
		n     int
		start int
		init  func(graph [][]DijkstraEdge)
	}
	ts := []struct {
		name         string
		args         args
		wantDistance []int
		wantFrom     []int
	}{
		{
			name: "",
			args: args{
				n:     9,
				start: 0,
				init: buildInit([]Edge{
					{from: 0, to: 1, weight: 4},
					{from: 0, to: 7, weight: 8},
					{from: 1, to: 2, weight: 8},
					{from: 1, to: 7, weight: 11},
					{from: 2, to: 8, weight: 2},
					{from: 2, to: 3, weight: 7},
					{from: 2, to: 5, weight: 4},
					{from: 3, to: 4, weight: 9},
					{from: 3, to: 5, weight: 14},
					{from: 4, to: 5, weight: 10},
					{from: 5, to: 6, weight: 2},
					{from: 6, to: 8, weight: 6},
					{from: 6, to: 7, weight: 1},
					{from: 7, to: 8, weight: 7},

					// 无向图，得双向
					{from: 1, to: 0, weight: 4},
					{from: 7, to: 0, weight: 8},
					{from: 2, to: 1, weight: 8},
					{from: 7, to: 1, weight: 11},
					{from: 8, to: 2, weight: 2},
					{from: 3, to: 2, weight: 7},
					{from: 5, to: 2, weight: 4},
					{from: 4, to: 3, weight: 9},
					{from: 5, to: 3, weight: 14},
					{from: 5, to: 4, weight: 10},
					{from: 6, to: 5, weight: 2},
					{from: 8, to: 6, weight: 6},
					{from: 7, to: 6, weight: 1},
					{from: 8, to: 7, weight: 7},
				}),
			},
			wantDistance: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
			wantFrom:     []int{-1, 0, 1, +2, +5, +6, 7, 0, +2}, // 使用 + 正号对齐
		},
		// todo 增加有向图的用例
	}

	fs := []func(n int, start int, init func(graph [][]DijkstraEdge)) (distance []int, from []int){
		Dijkstra,
		DijkstraHeap,
	}

	for _, f := range fs {
		name := tests.FuncName(f)
		t.Run(name, func(t *testing.T) {
			for _, tt := range ts {
				t.Run(tt.name, func(t *testing.T) {
					gotDistance, gotFrom := f(tt.args.n, tt.args.start, tt.args.init)
					if !reflect.DeepEqual(gotDistance, tt.wantDistance) {
						t.Errorf("%v() gotDistance = %v, want %v", name, gotDistance, tt.wantDistance)
					}
					if !reflect.DeepEqual(gotFrom, tt.wantFrom) {
						t.Errorf("%v() gotFrom = %v, want %v", name, gotFrom, tt.wantFrom)
					}
				})
			}
		})
	}
}
