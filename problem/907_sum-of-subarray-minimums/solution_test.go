package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(arr []int) int{
		bruteForce,
		sumSubarrayMins,
		sumSubarrayMins2,
		sumSubarrayMins3,
		sumSubarrayMins4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,1,2,4]
17

[11,81,94,43,3]
444

[3,1,3,2,4]
26

[3,1,3,2,4,10]
46

[1,2,3,4,5,6]
56

`
