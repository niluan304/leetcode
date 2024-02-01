package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countPairs,
		countPairs2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-1,1,2,3,1]
2
3

[-6,2,5,-2,-7,-1,3]
-2
10

[1,2,3,1]
7
6
`
