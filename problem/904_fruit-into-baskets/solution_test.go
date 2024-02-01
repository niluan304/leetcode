package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_fruit_into_baskets(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		totalFruit,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,1]
3

[0,1,2,2]
3

[1,2,3,2,2]
4

[3,3,3,1,2,1,1,2,3,3,4]
5

`
