package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, edges [][]int) []int{
		// bruteForce,
		sumOfDistancesInTree,
		// sumOfDistancesInTree2,
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
6
[[0,1],[0,2],[2,3],[2,4],[2,5]]
[8,12,6,10,10,10]

1
[]
[0]

2
[[1,0]]
[1,1]


`
