package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countSubarrays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int64{
		countSubarrays,
		//countSubarrays2,
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
[1,3,2,3,3]
2
6

[1,4,2,1]
3
0

[4,4,4,4,3,4,4,4,4,4]
9
1

[4,4,4,4,3,4,4,4,4,4]
4
24

[9,9,9,9,8,7,7,6,8,9,9,8,8,5,5]
2
75

[28,5,58,91,24,91,53,9,48,85,16,70,91,91,47,91,61,4,54,61,49]
1
187


`
