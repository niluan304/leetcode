package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_stoneGameVI(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(aliceValues []int, bobValues []int) int{
		//bruteForce,
		stoneGameVI,
		stoneGameVI2,
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
[1,3]
[2,1]
1

[1,2]
[3,1]
0

[2,4,3]
[1,6,7]
-1


`
