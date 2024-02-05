package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxResult(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		bruteForce,
		bruteForce2,
		maxResult,
		maxResult2,
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
[1,-1,-2,4,-7,3]
2
7

[10,-5,-2,4,0,3]
3
17

[1,-5,-20,4,-1,3,-6,-3]
2
0


`
