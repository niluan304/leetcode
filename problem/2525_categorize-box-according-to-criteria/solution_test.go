package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_categorize_box_according_to_criteria(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		categorizeBox,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
1000
35
700
300
"Heavy"

200
50
800
50
"Neither"

`
