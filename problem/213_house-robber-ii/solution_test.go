package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_house_robber_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		rob,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,2]
3

[1,2,3,1]
4

[1,2,3]
3

`
