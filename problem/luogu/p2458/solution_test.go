package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_string_transformation(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCost2,
	}

	for _, f := range fs {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		i := strings.LastIndex(name, ".")

		t.Run(name[i+1:], func(t *testing.T) {
			err := testutil.RunLeetCodeFuncWithFile(t, f, "sample.txt", targetCaseNum)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func minCost2(n int, workers [][]int) int {
	var graph = make([][]int, n+1)
	var costs = make([]int, n+1)
	for _, worker := range workers {
		i := worker[0]
		v := worker[1]
		graph[i] = worker[3:]
		costs[i] = v
	}
	return minCost(1, graph, costs)
}
