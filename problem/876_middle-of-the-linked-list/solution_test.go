package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_middle_of_the_linked_list(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		middleNode,
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
[3,4,5]

[1,2,3,4,5,6]
[4,5,6]

`
