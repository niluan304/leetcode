package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countRangeSum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, lower int, upper int) int{
		bruteForce,
		countRangeSum,
		countRangeSum2,
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
[-2,5,-1]
-2
2
3

[0]
0
0
1

[-2147483647,0,-2147483647,2147483647]
-564
3864
3

[1,4,-2,3,-4,3,0,-4,4]
3
6
16

[2147483647,2147483647,0,-2147483647]
4232
7660
0

`
