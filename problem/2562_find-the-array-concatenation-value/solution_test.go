package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_the_array_concatenation_value(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findTheArrayConcVal,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[7,52,2,4]
596

[5,14,13,8,12]
673

`
