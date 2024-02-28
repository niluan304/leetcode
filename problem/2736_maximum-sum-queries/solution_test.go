package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumSumQueries(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums1 []int, nums2 []int, queries [][]int) []int{
		bruteForce,
		maximumSumQueries,
		maximumSumQueries2,
		leetcodeMinTime,
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
[4,3,1,2]
[2,4,9,5]
[[4,1],[1,3],[2,5]]
[6,10,7]

[3,2,5]
[2,3,4]
[[4,4],[3,2],[1,1]]
[9,9,9]

[2,1]
[2,3]
[[3,3]]
[-1]

[66,69]
[39,19]
[[4,63]]
[-1]

[76,35,72]
[29,94,20]
[[69,34],[88,71]]
[-1,-1]

[72,88,53,63,95,46]
[78,56,35,72,56,63]
[[86,86],[24,8]]
[-1,151]

[13,67,90,92,47]
[52,60,69,49,73]
[[32,70]]
[120]

`
