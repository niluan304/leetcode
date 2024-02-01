package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_roman_to_integer(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		romanToInt,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `

"III"
3

"LVIII"
58

"MCMXCIV"
1994

`
