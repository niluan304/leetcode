package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reverse_linked_list_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		reverseBetween,
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
2
4
[1,4,3,2,5]

[5]
1
1
[5]

`
