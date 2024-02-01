package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_smallest_missing_genetic_value_in_each_subtree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		smallestMissingValueSubtree,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-1,0,0,2]
[1,2,3,4]
[5,1,1,1]

[-1,0,1,0,3,3]
[5,4,6,2,1,3]
[7,1,1,4,2,1]

[-1,2,3,0,2,4,1]
[2,3,4,5,6,7,8]
[1,1,1,1,1,1,1]

`
