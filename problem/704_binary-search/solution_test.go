package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_binary_search(t *testing.T) {
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
[-1,0,3,5,9,12]
9
4

[-1,0,3,5,9,12]
2
-1

`
