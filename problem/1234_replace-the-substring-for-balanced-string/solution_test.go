package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_replace_the_substring_for_balanced_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		balancedString,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"QWER"
0

"QQWE"
1

"QQQW"
2

`
