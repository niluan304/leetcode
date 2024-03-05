package graph

import (
	"reflect"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func TestDijkstra(t *testing.T) {
	type Edge struct {
		i, j   int
		weight int
	}

	init := func(edges []Edge) func(graph [][]DijkstraEdge) {
		return func(graph [][]DijkstraEdge) {
			for _, edge := range edges {
				x, y, w := edge.i, edge.j, edge.weight
				graph[x] = append(graph[x], DijkstraEdge{To: y, Weight: w})
				graph[y] = append(graph[y], DijkstraEdge{To: x, Weight: w})
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
				init: init([]Edge{
					{i: 0, j: 1, weight: 4},
					{i: 0, j: 7, weight: 8},
					{i: 1, j: 2, weight: 8},
					{i: 1, j: 7, weight: 11},
					{i: 2, j: 8, weight: 2},
					{i: 2, j: 3, weight: 7},
					{i: 2, j: 5, weight: 4},
					{i: 3, j: 4, weight: 9},
					{i: 3, j: 5, weight: 14},
					{i: 4, j: 5, weight: 10},
					{i: 5, j: 6, weight: 2},
					{i: 6, j: 8, weight: 6},
					{i: 6, j: 7, weight: 1},
					{i: 7, j: 8, weight: 7},
				}),
			},
			wantDistance: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
			wantFrom:     []int{-1, 0, 1, 0o2, 0o5, 0o6, 7, 0, 0o2},
		},
	}

	fs := []func(n int, start int, init func(graph [][]DijkstraEdge)) (distance []int, from []int){
		Dijkstra,
		DijkstraPriorityQueue,
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
