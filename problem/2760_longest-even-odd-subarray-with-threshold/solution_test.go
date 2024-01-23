package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longestAlternatingSubarray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, threshold int) int{
		longestAlternatingSubarray,
		//longestAlternatingSubarray2,
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
[3,2,5,4]
5
3

[1,2]
2
1

[2,3,4,5]
4
3

[4,10,3]
10
2

[4]
1
0

[2,2]
18
1

[2]
2
1

`
