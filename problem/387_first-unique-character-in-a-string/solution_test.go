package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_first_unique_character_in_a_string(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		firstUniqChar,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"leetcode"
0

"loveleetcode"
2

"aabb"
-1

`
