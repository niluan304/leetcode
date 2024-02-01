package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_remove_linked_list_elements(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		removeElements,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,6,3,4,5,6]
6
[1,2,3,4,5]

[]
1
[]

[7,7,7,7]
7
[]

`
