package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countPaths(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, edges [][]int) int64{
		// bruteForce,
		countPaths,
		// countPaths2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
5
[[1,2],[1,3],[2,4],[2,5]]
4

6
[[1,2],[1,3],[2,4],[3,5],[3,6]]
6

4
[[1,2],[4,1],[3,4]]
4


9
[[7,4],[3,4],[5,4],[1,5],[6,4],[9,5],[8,7],[2,8]]
17

12
[[1,3],[1,4],[1,6],[3,5],[3,8],[3,9],[4,7],[4,10],[4,11],[6,12],[6,2]]
33

12
[[2,3],[2,4],[2,6],[3,5],[3,8],[3,9],[4,7],[4,10],[4,11],[6,12],[6,1]]
18

`
