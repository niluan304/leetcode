package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count_the_digits_that_divide_a_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		countDigits,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
7
1

121
2

1248
4

`
