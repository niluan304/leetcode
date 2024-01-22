package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_splitArray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		splitArray,
		splitArray2,
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
[7,2,5,10,8]
2
18

[1,2,3,4,5]
2
9

[1,4,4]
3
4

[1,2,3,4,5]
1
15

[2,3,1,2,4,3]
5
4

`
