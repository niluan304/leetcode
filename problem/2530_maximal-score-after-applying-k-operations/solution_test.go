package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximal_score_after_applying_k_operations(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxKelements,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[10,10,10,10,10]
5
50

[1,10,3,3,3]
3
17

`
