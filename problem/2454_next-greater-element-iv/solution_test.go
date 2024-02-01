package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		secondGreaterElement,
		secondGreaterElement2,
		leetcodeEndless,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,4,0,9,6]
[9,6,6,-1,-1]

[3,3]
[-1,-1]

[1,2,4,3]
[4,3,-1,-1]

[2,4,3,9,6]
[3,6,6,-1,-1]

`
