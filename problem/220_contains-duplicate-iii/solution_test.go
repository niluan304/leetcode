package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_contains_duplicate_iii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		containsNearbyAlmostDuplicate,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,1]
3
0
true

[1,5,9,1,5,9]
2
3
false

`
