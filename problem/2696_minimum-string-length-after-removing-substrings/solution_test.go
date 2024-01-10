package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minLength(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) int{
		minLength,
		minLength2,
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
"ABFCACDB"
2

"ACBBD"
5


`
