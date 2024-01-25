package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_findIndices(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, indexDifference int, valueDifference int) []int{
		bruteForce,
		findIndices,
		// findIndices2,
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
[5,1,4,1]
2
4
[0,3]

[2,1]
0
0
[0,0]

[1,2,3]
2
4
[-1,-1]

[3,12,40]
0
9
[0,1]

[22,28,50,25]
0
28
[0,2]

[0,8,2,5,10]
0
10
[0,4]

`
