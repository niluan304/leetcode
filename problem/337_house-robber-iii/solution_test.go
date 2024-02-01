package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_house_robber_iii(t *testing.T) {
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
[3,2,3,null,3,null,1]
7

[3,4,5,1,3,null,1]
9

`
