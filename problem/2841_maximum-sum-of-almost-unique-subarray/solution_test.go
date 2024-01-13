package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, m int, k int) int64{
		maxSum,
		// maxSum2,
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
[2,6,7,3,1,7]
3
4
18

[5,9,9,2,4,5,4]
1
3
23

[1,2,1,2,1,2,1]
3
3
0

[2,6,7,3,1,99]
3
4
110

`
