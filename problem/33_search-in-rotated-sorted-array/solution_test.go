package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_search_in_rotated_sorted_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		search,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,5,6,7,0,1,2]
0
4

[4,5,6,7,0,1,2]
3
-1

[1]
0
-1

`
