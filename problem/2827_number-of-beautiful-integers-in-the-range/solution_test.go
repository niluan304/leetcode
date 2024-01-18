package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_numberOfBeautifulIntegers(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(low int, high int, k int) int{
		numberOfBeautifulIntegers,
		// numberOfBeautifulIntegers2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
10
20
3
2

1
10
1
1

5
5
2
0


`
