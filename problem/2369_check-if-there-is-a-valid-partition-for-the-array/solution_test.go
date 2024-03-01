package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_validPartition(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) bool{
		// bruteForce,
		validPartition,
		validPartition2,
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
[4,4,4,5,6]
true

[1,1,1,2]
false


`
