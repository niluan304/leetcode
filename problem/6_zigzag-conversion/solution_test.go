package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_zigzag_conversion(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		convert,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `


"PAYPALISHIRING"
3
"PAHNAPLSIIGYIR"

"PAYPALISHIRING"
4
"PINALSIGYAHRPI"

"A"
1
"A"

`
