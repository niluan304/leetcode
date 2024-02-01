package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_russian_doll_envelopes(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxEnvelopes,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[5,4],[6,4],[6,7],[2,3]]
3

[[1,1],[1,1],[1,1]]
1

`
