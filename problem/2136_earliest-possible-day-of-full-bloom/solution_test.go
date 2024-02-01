package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_earliest_possible_day_of_full_bloom(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		earliestFullBloom,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,4,3]
[2,3,1]
9

[1,2,3,2]
[2,1,2,1]
9

[1]
[1]
2

`
