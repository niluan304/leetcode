package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reverse_linked_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		reverseList,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5]
[5,4,3,2,1]

[1,2]
[2,1]

[]
[]

`
