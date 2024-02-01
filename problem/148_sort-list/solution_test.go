package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_sort_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		sortList,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,2,1,3]
[1,2,3,4]

[-1,5,3,4,0]
[-1,0,3,4,5]

[]
[]

`
