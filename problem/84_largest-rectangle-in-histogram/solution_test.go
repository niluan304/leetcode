package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		largestRectangleArea,
		largestRectangleArea2,
		largestRectangleArea3,
		largestRectangleArea4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,1,5,6,2,3]
10

[2,4]
4

[2,1,5,6,3,3]
12

[2,1,5,6,4,3]
12

[2,1,5,6,4,3,15]
15

[1,2,3,4,3,2,1,1,2,3,4,3,2,1]
14
`
