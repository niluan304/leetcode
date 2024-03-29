package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_combinationSum3(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(k int, n int) [][]int{
		//bruteForce,
		combinationSum3,
		combinationSum3ii,
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
3
7
[[1,2,4]]

3
9
[[1,2,6],[1,3,5],[2,3,4]]

4
1
[]

9
45
[[1,2,3,4,5,6,7,8,9]]


`
