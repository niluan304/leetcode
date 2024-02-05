package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minIncrementOperations(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int64{
		//bruteForce,
		minIncrementOperations,
		minIncrementOperations2,
		minIncrementOperations3,
		minIncrementOperations4,
		//leetcode,
		endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,0,0,2]
4
3

[0,1,3,3]
5
2

[1,1,2]
1
0

[20,38,29,34,6]
95
66

[43,31,14,4]
73
42

`
