package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_search_insert_position(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		searchInsert,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,5,6]
5
2

[1,3,5,6]
2
1

[1,3,5,6]
7
4

`
