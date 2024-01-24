package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_checkArray(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) bool{
		//bruteForce,
		checkArray,
		//checkArray2,
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
[2,2,3,1,1,0]
3
true

[1,3,1,1]
2
false

[0,45,82,98,99]
4
false

`
