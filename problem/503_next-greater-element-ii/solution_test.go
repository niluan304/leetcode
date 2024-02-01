package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		nextGreaterElements,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,1]
[2,-1,2]

[1,2,3,4,3]
[2,3,4,-1,4]

[4,2,3,4,3]
[-1,3,4,-1,4]


`
