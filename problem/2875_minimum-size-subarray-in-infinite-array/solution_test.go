package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minSizeSubarray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, target int) int{
		minSizeSubarray,
		minSizeSubarray2,
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
[1,2,3]
5
2

[1,1,1,2,3]
4
2

[2,4,6,8]
3
-1

[1,2,3,4]
29
11

[1,2,3,4]
10
4

[1,2,3,4]
8
3

`
