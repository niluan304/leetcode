package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_eliminate_maximum_number_of_monsters(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		eliminateMaximum,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,3,4]
[1,1,1]
3

[1,1,2,3]
[1,1,1,1]
1

[3,2,4]
[5,3,2]
1

`
