package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_ways_to_build_good_strings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countGoodStrings,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
3
3
1
1
8

2
3
1
2
5

`
