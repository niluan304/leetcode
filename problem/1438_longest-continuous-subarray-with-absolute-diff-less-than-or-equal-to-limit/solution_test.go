package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, limit int) int{
		longestSubarray,
		longestSubarray2,
		longestSubarray3,
		longestSubarray4,
		leetcode2,
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
[8,2,4,7]
4
2

[10,1,2,4,7,2]
5
4

[4,2,2,2,4,4,2,2]
0
3

[8,6,2,4,7]
4
3

`
