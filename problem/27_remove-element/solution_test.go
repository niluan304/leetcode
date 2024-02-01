package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_remove_element(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		removeElement,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `

[3,2,2,3]
3
2

[0,1,2,2,3,0,4,2]
2
5

`
