package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_car_pooling(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		carPooling,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[2,1,5],[3,3,7]]
4
false

[[2,1,5],[3,3,7]]
5
true

`
