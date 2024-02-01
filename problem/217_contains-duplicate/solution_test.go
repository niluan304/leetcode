package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_contains_duplicate(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		containsDuplicate,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,1]
true

[1,2,3,4]
false

[1,1,1,3,3,4,3,2,4,2]
true

`
