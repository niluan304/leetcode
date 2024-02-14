package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combine(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, k int) [][]int{
		//bruteForce,
		combine,
		//combine2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
4
2
[[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]

1
1
[[1]]


`
