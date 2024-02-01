package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_increasing_subsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		lengthOfLIS,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,9,2,5,3,7,101,18]
4

[0,1,0,3,2,3]
4

[7,7,7,7,7,7,7]
1

`
