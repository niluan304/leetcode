package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumBeauty(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		maximumBeauty,
		// maximumBeauty2,
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
[4,6,1,2]
2
3

[1,1,1,1]
10
4


[1,2,4,6]
1
2

[1,2,4,5,6]
1
3

[75,15,9]
28
2

`
