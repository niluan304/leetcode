package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_single_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		singleNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,2,1]
1

[4,1,2,1,2]
4

[1]
1

`
