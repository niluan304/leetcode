package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) []int{
		maxSlidingWindow,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,-1,-3,5,3,6,7]
3
[3,3,5,5,6,7]

[1]
1
[1]

[7,2,4]
2
[7,4]

`
