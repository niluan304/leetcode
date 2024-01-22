package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumJumps(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, target int) int{
		maximumJumps,
		// maximumJumps2,
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
[1,3,6,4,1,2]
2
3

[1,3,6,4,1,2]
3
5

[1,3,6,4,1,2]
0
-1

[99,1,1,1,1,1,1,1]
3
-1


[1,1,1,1,1,1,1,99]
3
-1

[0,2,1]
1
1

`
