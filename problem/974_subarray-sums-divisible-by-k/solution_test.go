package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_subarraysDivByK(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		//bruteForce,
		subarraysDivByK,
		subarraysDivByK2,
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
[4,5,0,-2,-3,1]
5
7

[5]
9
0

[-1,2,9]
2
2


`
