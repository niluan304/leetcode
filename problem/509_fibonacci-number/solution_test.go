package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_fibonacci_number(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		fib,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `

2
1

3
2

4
3

`
