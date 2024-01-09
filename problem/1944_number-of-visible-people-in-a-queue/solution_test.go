package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(heights []int) []int{
		canSeePersonsCount,
		canSeePersonsCount2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,6,8,5,11,9]
[3,1,2,1,1,0]

[5,1,2,3,10]
[4,1,1,1,0]


`
