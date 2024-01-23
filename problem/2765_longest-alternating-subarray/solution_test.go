package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_alternatingSubarray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		alternatingSubarray,
		alternatingSubarray2,
		alternatingSubarray3,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,4,3,4]
4

[4,5,6]
2

[21,9,5]
-1

[64,65,64,65,64,65,66,65,66,65]
6

`
