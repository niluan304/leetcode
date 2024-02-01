package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_powx_n(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		myPow,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
2.00000
10
1024.00000

2.10000
3
9.26100

2.00000
-2
0.25000

`
