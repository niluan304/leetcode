package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lexicographicallySmallestArray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, limit int) []int{
		//bruteForce,
		lexicographicallySmallestArray,
		//lexicographicallySmallestArray2,
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
[1,5,3,9,8]
2
[1,3,5,8,9]

[1,7,6,18,2,1]
3
[1,6,7,18,1,2]

[1,7,28,19,10]
3
[1,7,28,19,10]


`
