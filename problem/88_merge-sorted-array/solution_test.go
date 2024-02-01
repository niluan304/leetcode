package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_merge_sorted_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		merge,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,0,0,0]
3
[2,5,6]
3
[1,2,2,3,5,6]

[1]
1
[]
0
[1]

[0]
0
[1]
1
[1]

`
