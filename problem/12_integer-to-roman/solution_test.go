package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_integer_to_roman(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		intToRoman,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `

3
"III"

58
"LVIII"

1994
"MCMXCIV"

`
