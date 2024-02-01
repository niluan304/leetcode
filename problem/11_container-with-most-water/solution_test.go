package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_container_with_most_water(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxArea,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,8,6,2,5,4,8,3,7]
49

[1,1]
1

`
